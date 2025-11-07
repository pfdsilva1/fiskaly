package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pfdsilva1/fiskaly/signing-service-challenge-go/domain"
	servicemocks "github.com/pfdsilva1/fiskaly/signing-service-challenge-go/service/mocks"
	"go.uber.org/mock/gomock"
)

func TestListSignatureDevices_HappyPath(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := servicemocks.NewMockSignatureDeviceService(ctrl)
	s := &Server{service: mockSvc}

	expected := []*domain.SignatureDevice{{Algorithm: "RSA", PrivateKey: "secret", PublicKey: "pub"}}
	mockSvc.EXPECT().ListSignatureDevices().Return(expected, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/v0/devices", nil)
	rr := httptest.NewRecorder()

	s.ListSignatureDevices(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}

	// Ensure we don't leak private_key in the API response
	var resp struct {
		Data []map[string]any `json:"data"`
	}
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(resp.Data) != 1 {
		t.Fatalf("expected 1 device, got %d", len(resp.Data))
	}
	if _, ok := resp.Data[0]["private_key"]; ok {
		t.Fatalf("response leaked private_key")
	}
}

func TestCreateSignatureDevice_HappyPath(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := servicemocks.NewMockSignatureDeviceService(ctrl)
	s := &Server{service: mockSvc}

	mockSvc.EXPECT().NewSignatureDevice("RSA", "label").Return("id-123", nil)

	req := httptest.NewRequest(http.MethodPost, "/api/v0/devices?algorithm=RSA&label=label", nil)
	rr := httptest.NewRecorder()

	s.CreateSignatureDevice(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
}

func TestCreateSignatureDevice_BadRequest_MissingAlgorithm(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := servicemocks.NewMockSignatureDeviceService(ctrl)
	s := &Server{service: mockSvc}

	req := httptest.NewRequest(http.MethodPost, "/api/v0/devices", nil)
	rr := httptest.NewRecorder()

	s.CreateSignatureDevice(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rr.Code)
	}
}

func TestSignTransaction_HappyPath(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := servicemocks.NewMockSignatureDeviceService(ctrl)
	s := &Server{service: mockSvc}

	rec := domain.SignatureRecord{Counter: 0}
	mockSvc.EXPECT().SignTransaction("dev123", []byte("payload")).Return(rec, nil)

	req := httptest.NewRequest(http.MethodPost, "/api/v0/devices/dev123/sign", bytes.NewBufferString("payload"))
	req.SetPathValue("device_id", "dev123")
	rr := httptest.NewRecorder()

	s.SignTransaction(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
}

func TestSignTransaction_BadRequest_MissingDeviceID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := servicemocks.NewMockSignatureDeviceService(ctrl)
	s := &Server{service: mockSvc}

	req := httptest.NewRequest(http.MethodPost, "/api/v0/devices//sign", bytes.NewBufferString("payload"))
	rr := httptest.NewRecorder()

	s.SignTransaction(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rr.Code)
	}
}

func TestDeviceSignatures_HappyPath(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := servicemocks.NewMockSignatureDeviceService(ctrl)
	s := &Server{service: mockSvc}

	expected := []domain.SignatureRecord{{Counter: 1}}
	mockSvc.EXPECT().ListSignatures("dev123").Return(expected, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/v0/devices/dev123/signatures", nil)
	req.SetPathValue("device_id", "dev123")
	rr := httptest.NewRecorder()

	s.DeviceSignatures(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
}

func TestDeviceSignatures_BadRequest_MissingDeviceID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := servicemocks.NewMockSignatureDeviceService(ctrl)
	s := &Server{service: mockSvc}

	req := httptest.NewRequest(http.MethodGet, "/api/v0/devices//signatures", nil)
	rr := httptest.NewRecorder()

	s.DeviceSignatures(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rr.Code)
	}
}

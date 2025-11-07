package service

import (
	"testing"

	"github.com/pfdsilva1/fiskaly/signing-service-challenge-go/crypto"
	"github.com/pfdsilva1/fiskaly/signing-service-challenge-go/domain"
	reposmocks "github.com/pfdsilva1/fiskaly/signing-service-challenge-go/persistence/mocks"
	"go.uber.org/mock/gomock"
)

func TestSignatureService_NewSignatureDevice_HappyPath(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := reposmocks.NewMockSignatureDeviceRepository(ctrl)
	svc := NewSignatureService(repo)

	// We don't control actual values from crypto, only ensure repo is called.
	repo.EXPECT().CreateSignatureDevice(gomock.Any()).Return(nil)

	id, err := svc.NewSignatureDevice("ECDSA", "label")
	if err != nil || id == "" {
		t.Fatalf("unexpected error or empty id: %v", err)
	}
}

func TestSignatureService_SignTransaction_HappyPath(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := reposmocks.NewMockSignatureDeviceRepository(ctrl)
	svc := NewSignatureService(repo)

	pub, priv, err := crypto.GenerateKeyPair("ECDSA")
	if err != nil {
		t.Fatalf("GenerateKeyPair: %v", err)
	}
	dev := domain.NewSignatureDevice("ECDSA", "label", pub, priv)

	repo.EXPECT().GetSignatureDevice(dev.ID.String()).Return(dev, nil)
	repo.EXPECT().SaveSignature(dev.ID.String(), gomock.Any()).Return(nil)

	rec, err := svc.SignTransaction(dev.ID.String(), []byte("data"))
	if err != nil || rec.Signature == "" {
		t.Fatalf("unexpected error or empty signature: %v", err)
	}
}

func TestSignatureService_ListSignatureDevices_HappyPath(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := reposmocks.NewMockSignatureDeviceRepository(ctrl)
	svc := NewSignatureService(repo)

	repo.EXPECT().ListSignatureDevices().Return([]*domain.SignatureDevice{}, nil)

	_, err := svc.ListSignatureDevices()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSignatureService_ListSignatures_HappyPath(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := reposmocks.NewMockSignatureDeviceRepository(ctrl)
	svc := NewSignatureService(repo)

	repo.EXPECT().ListSignatures("dev").Return([]domain.SignatureRecord{}, nil)

	_, err := svc.ListSignatures("dev")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

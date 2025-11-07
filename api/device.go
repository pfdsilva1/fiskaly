package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/pfdsilva1/fiskaly/signing-service-challenge-go/domain"
	"github.com/pfdsilva1/fiskaly/signing-service-challenge-go/persistence"
)

func (s *Server) SignatureDevices(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		s.ListSignatureDevices(response, request)
	case http.MethodPost:
		s.CreateSignatureDevice(response, request)
	default:
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
	}
}

// SignatureDeviceReadModel represents the safe API view of a signature device.
// It intentionally excludes the PrivateKey to avoid leaking sensitive material.
type SignatureDeviceReadModel struct {
	ID               string `json:"id"`
	Algorithm        string `json:"algorithm"`
	Label            string `json:"label"`
	PublicKey        string `json:"public_key"`
	SignatureCounter uint64 `json:"signature_counter"`
	LastSignature    string `json:"last_signature"`
}

func toSignatureDeviceReadModel(d *domain.SignatureDevice) SignatureDeviceReadModel {
	return SignatureDeviceReadModel{
		ID:               d.ID.String(),
		Algorithm:        d.Algorithm,
		Label:            d.Label,
		PublicKey:        d.PublicKey,
		SignatureCounter: d.SignatureCounter,
		LastSignature:    d.LastSignature,
	}
}

func (s *Server) ListSignatureDevices(response http.ResponseWriter, request *http.Request) {
	signatureDevices, err := s.service.ListSignatureDevices()
	if err != nil {
		WriteErrorResponse(response, http.StatusInternalServerError, []string{
			fmt.Sprintf("Failed to list signature devices: %v", err),
		})
		return
	}

	safe := make([]SignatureDeviceReadModel, 0, len(signatureDevices))
	for _, d := range signatureDevices {
		safe = append(safe, toSignatureDeviceReadModel(d))
	}

	WriteAPIResponse(response, http.StatusOK, safe)
}

func (s *Server) CreateSignatureDevice(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	algorithm := request.URL.Query().Get("algorithm")
	if algorithm == "" {
		WriteErrorResponse(response, http.StatusBadRequest, []string{
			"Algorithm is required as query parameter",
		})
		return
	}
	if !isAlgorithmAllowed(algorithm) {
		WriteErrorResponse(response, http.StatusBadRequest, []string{
			"Invalid algorithm. Only RSA and ECDSA are supported.",
		})
		return
	}

	deviceID, err := s.service.NewSignatureDevice(algorithm, request.URL.Query().Get("label"))
	if err != nil {
		WriteErrorResponse(response, http.StatusInternalServerError, []string{
			fmt.Sprintf("Failed to create signature device: %v", err),
		})
		return
	}

	WriteAPIResponse(response, http.StatusOK, deviceID)
}

// isAlgorithmAllowed checks if the provided algorithm is allowed.
func isAlgorithmAllowed(algorithm string) bool {
	switch algorithm {
	case "RSA", "ECDSA":
		return true
	default:
		return false
	}
}

func (s *Server) SignTransaction(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	deviceID := request.PathValue("device_id")
	if deviceID == "" {
		WriteErrorResponse(response, http.StatusBadRequest, []string{
			"Device ID is required",
		})
		return
	}

	data, err := io.ReadAll(request.Body)
	if err != nil {
		WriteErrorResponse(response, http.StatusInternalServerError, []string{
			fmt.Sprintf("Failed to read data: %v", err),
		})
		return
	}

	signature, err := s.service.SignTransaction(deviceID, data)
	if err != nil {
		var nf *persistence.SignatureDeviceNotFoundError
		if errors.As(err, &nf) {
			WriteErrorResponse(response, http.StatusNotFound, []string{nf.Error()})
			return
		}
		WriteErrorResponse(response, http.StatusInternalServerError, []string{
			fmt.Sprintf("Failed to sign transaction: %v", err),
		})
		return
	}

	WriteAPIResponse(response, http.StatusOK, signature)
}

func (s *Server) DeviceSignatures(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	deviceID := request.PathValue("device_id")
	if deviceID == "" {
		WriteErrorResponse(response, http.StatusBadRequest, []string{
			"Device ID is required",
		})
		return
	}

	signatures, err := s.service.ListSignatures(deviceID)
	if err != nil {
		var nf *persistence.SignatureDeviceNotFoundError
		if errors.As(err, &nf) {
			WriteErrorResponse(response, http.StatusNotFound, []string{nf.Error()})
			return
		}
		WriteErrorResponse(response, http.StatusInternalServerError, []string{
			fmt.Sprintf("Failed to list signatures: %v", err),
		})
		return
	}

	WriteAPIResponse(response, http.StatusOK, signatures)
}

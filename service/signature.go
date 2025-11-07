package service

import (
	"fmt"

	"github.com/pfdsilva1/fiskaly/signing-service-challenge-go/crypto"
	"github.com/pfdsilva1/fiskaly/signing-service-challenge-go/domain"
	"github.com/pfdsilva1/fiskaly/signing-service-challenge-go/persistence"
)

type SignatureDeviceService interface {
	NewSignatureDevice(algorithm string, label string) (string, error)
	SignTransaction(deviceID string, data []byte) (domain.SignatureRecord, error)
	ListSignatureDevices() ([]*domain.SignatureDevice, error)
	ListSignatures(deviceID string) ([]domain.SignatureRecord, error)
}

type SignatureService struct {
	signatureDeviceRepository persistence.SignatureDeviceRepository
}

func NewSignatureService(signatureDeviceRepository persistence.SignatureDeviceRepository) *SignatureService {
	return &SignatureService{signatureDeviceRepository: signatureDeviceRepository}
}

func (s *SignatureService) NewSignatureDevice(algorithm string, label string) (string, error) {
	publicKey, privateKey, err := crypto.GenerateKeyPair(algorithm)
	if err != nil {
		return "", fmt.Errorf("failed to generate key pair: %w", err)
	}

	signatureDevice := domain.NewSignatureDevice(algorithm, label, publicKey, privateKey)

	if err := s.signatureDeviceRepository.CreateSignatureDevice(signatureDevice); err != nil {
		return "", fmt.Errorf("failed to create signature device: %w", err)
	}

	return signatureDevice.ID.String(), nil
}

func (s *SignatureService) SignTransaction(deviceID string, data []byte) (domain.SignatureRecord, error) {
	signatureDevice, err := s.signatureDeviceRepository.GetSignatureDevice(deviceID)
	if err != nil {
		return domain.SignatureRecord{}, fmt.Errorf("failed to get signature device: %w", err)
	}

	signature, err := signatureDevice.Sign(data)
	if err != nil {
		return domain.SignatureRecord{}, fmt.Errorf("failed to sign transaction: %w", err)
	}

	if err := s.signatureDeviceRepository.SaveSignature(deviceID, signature); err != nil {
		return domain.SignatureRecord{}, fmt.Errorf("failed to store signature: %w", err)
	}

	return signature, nil
}

func (s *SignatureService) ListSignatureDevices() ([]*domain.SignatureDevice, error) {
	return s.signatureDeviceRepository.ListSignatureDevices()
}

func (s *SignatureService) ListSignatures(deviceID string) ([]domain.SignatureRecord, error) {
	signatures, err := s.signatureDeviceRepository.ListSignatures(deviceID)
	if err != nil {
		return nil, fmt.Errorf("failed to list signatures: %w", err)
	}

	return signatures, nil
}

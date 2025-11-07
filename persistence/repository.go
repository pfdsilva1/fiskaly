package persistence

import "github.com/pfdsilva1/fiskaly/signing-service-challenge-go/domain"

type SignatureDeviceRepository interface {
	CreateSignatureDevice(signatureDevice *domain.SignatureDevice) error
	GetSignatureDevice(id string) (*domain.SignatureDevice, error)
	ListSignatureDevices() ([]*domain.SignatureDevice, error)
	SaveSignature(deviceID string, signature domain.SignatureRecord) error
	ListSignatures(deviceID string) ([]domain.SignatureRecord, error)
}

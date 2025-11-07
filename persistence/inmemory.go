package persistence

import (
	"sync"

	"github.com/pfdsilva1/fiskaly/signing-service-challenge-go/domain"
)

type InMemorySignatureDeviceRepository struct {
	signatureDevices map[string]*domain.SignatureDevice
	signatures       map[string][]domain.SignatureRecord
	mu               sync.Mutex
}

func NewInMemorySignatureDeviceRepository() *InMemorySignatureDeviceRepository {
	return &InMemorySignatureDeviceRepository{
		signatureDevices: make(map[string]*domain.SignatureDevice),
		signatures:       make(map[string][]domain.SignatureRecord),
	}
}

func (r *InMemorySignatureDeviceRepository) CreateSignatureDevice(signatureDevice *domain.SignatureDevice) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.signatureDevices[signatureDevice.ID.String()] = signatureDevice
	r.signatures[signatureDevice.ID.String()] = make([]domain.SignatureRecord, 0)
	return nil
}

func (r *InMemorySignatureDeviceRepository) GetSignatureDevice(id string) (*domain.SignatureDevice, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	signatureDevice, ok := r.signatureDevices[id]
	if !ok {
		return nil, &SignatureDeviceNotFoundError{DeviceID: id}
	}

	return signatureDevice, nil
}

func (r *InMemorySignatureDeviceRepository) ListSignatureDevices() ([]*domain.SignatureDevice, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Return sanitized copies that do not include PrivateKey
	signatureDevices := make([]*domain.SignatureDevice, 0, len(r.signatureDevices))
	for _, d := range r.signatureDevices {
		copy := &domain.SignatureDevice{
			ID:               d.ID,
			Algorithm:        d.Algorithm,
			Label:            d.Label,
			PublicKey:        d.PublicKey,
			PrivateKey:       "",
			SignatureCounter: d.SignatureCounter,
			LastSignature:    d.LastSignature,
		}
		signatureDevices = append(signatureDevices, copy)
	}

	return signatureDevices, nil
}

func (r *InMemorySignatureDeviceRepository) SaveSignature(deviceID string, signature domain.SignatureRecord) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.signatureDevices[deviceID]; !ok {
		return &SignatureDeviceNotFoundError{DeviceID: deviceID}
	}

	r.signatures[deviceID] = append(r.signatures[deviceID], signature)
	return nil
}

func (r *InMemorySignatureDeviceRepository) ListSignatures(deviceID string) ([]domain.SignatureRecord, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	signatures, ok := r.signatures[deviceID]
	if !ok {
		return nil, &SignatureDeviceNotFoundError{DeviceID: deviceID}
	}

	result := make([]domain.SignatureRecord, len(signatures))
	copy(result, signatures)
	return result, nil
}

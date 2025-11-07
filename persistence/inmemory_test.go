package persistence

import (
	"testing"

	"github.com/pfdsilva1/fiskaly/signing-service-challenge-go/domain"
)

func TestInMemoryRepository_HappyPath(t *testing.T) {
	t.Parallel()
	repo := NewInMemorySignatureDeviceRepository()

	dev := domain.NewSignatureDevice("RSA", "label", "pub", "priv")
	if err := repo.CreateSignatureDevice(dev); err != nil {
		t.Fatalf("CreateSignatureDevice: %v", err)
	}

	got, err := repo.GetSignatureDevice(dev.ID.String())
	if err != nil || got == nil {
		t.Fatalf("GetSignatureDevice: %v", err)
	}

    list, err := repo.ListSignatureDevices()
    if err != nil || len(list) != 1 {
        t.Fatalf("ListSignatureDevices: %v len=%d", err, len(list))
    }
    if list[0].PrivateKey != "" {
        t.Fatalf("ListSignatureDevices leaked PrivateKey")
    }

	rec := domain.SignatureRecord{Counter: 0, Signature: "sig", SignedData: "data"}
	if err := repo.SaveSignature(dev.ID.String(), rec); err != nil {
		t.Fatalf("SaveSignature: %v", err)
	}

	sigs, err := repo.ListSignatures(dev.ID.String())
	if err != nil || len(sigs) != 1 {
		t.Fatalf("ListSignatures: %v len=%d", err, len(sigs))
	}
}

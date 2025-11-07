package domain

import (
	"encoding/base64"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/pfdsilva1/fiskaly/signing-service-challenge-go/crypto"
)

// TODO: signature device domain model ...
type SignatureDevice struct {
	ID               uuid.UUID
	Algorithm        string
	Label            string
	PublicKey        string
	PrivateKey       string
	SignatureCounter uint64
	LastSignature    string
	mu               sync.Mutex
}

type SignatureRecord struct {
	Counter    uint64 `json:"counter"`
	Signature  string `json:"signature"`
	SignedData string `json:"signed_data"`
}

func NewSignatureDevice(algorithm string, label string, publicKey string, privateKey string) *SignatureDevice {
	id := uuid.New()
	lastSignature := base64.StdEncoding.EncodeToString(id[:])
	return &SignatureDevice{
		ID:               id,
		Algorithm:        algorithm,
		Label:            label,
		PublicKey:        publicKey,
		PrivateKey:       privateKey,
		SignatureCounter: 0,
		LastSignature:    lastSignature,
	}
}

func (d *SignatureDevice) Sign(data []byte) (SignatureRecord, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	counter := d.SignatureCounter
	securedData := fmt.Sprintf("%d_%s_%s", d.SignatureCounter, string(data), d.LastSignature)
	signature, err := crypto.Sign(d.Algorithm, d.PrivateKey, []byte(securedData))
	if err != nil {
		return SignatureRecord{}, fmt.Errorf("failed to sign transaction: %w", err)
	}

	d.SignatureCounter++
	encodedSignature := base64.StdEncoding.EncodeToString(signature)
	d.LastSignature = encodedSignature

	return SignatureRecord{
		Counter:    counter,
		Signature:  encodedSignature,
		SignedData: securedData,
	}, nil
}

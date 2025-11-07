package crypto

import (
	stdcrypto "crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

// Signer defines a contract for different types of signing implementations.
type Signer interface {
	Sign(privatePEM string, dataToBeSigned []byte) ([]byte, error)
}

type RSASigner struct{}

func NewRSASigner() *RSASigner {
	return &RSASigner{}
}

func (s *RSASigner) Sign(privatePEM string, data []byte) ([]byte, error) {
	pair, err := NewRSAMarshaler().Unmarshal([]byte(privatePEM))
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal private key: %w", err)
	}
	hashed := sha256.Sum256(data)
	return rsa.SignPKCS1v15(rand.Reader, pair.Private, stdcrypto.SHA256, hashed[:])
}

type ECCSigner struct {
}

func NewECCSigner() *ECCSigner {
	return &ECCSigner{}
}

func (s *ECCSigner) Sign(privatePEM string, data []byte) ([]byte, error) {
	pair, err := NewECCMarshaler().Decode([]byte(privatePEM))
	if err != nil {
		return nil, fmt.Errorf("failed to decode private key: %w", err)
	}
	hashed := sha256.Sum256(data)
	return ecdsa.SignASN1(rand.Reader, pair.Private, hashed[:])
}

package domain

import (
	"testing"

	"github.com/pfdsilva1/fiskaly/signing-service-challenge-go/crypto"
)

func TestSignatureDevice_Sign_HappyPath(t *testing.T) {
	t.Parallel()
	pub, priv, err := crypto.GenerateKeyPair("ECDSA")
	if err != nil {
		t.Fatalf("GenerateKeyPair: %v", err)
	}
	d := NewSignatureDevice("ECDSA", "label", pub, priv)
	rec1, err := d.Sign([]byte("data1"))
	if err != nil {
		t.Fatalf("Sign: %v", err)
	}
	if rec1.Counter != 0 || rec1.Signature == "" || rec1.SignedData == "" {
		t.Fatalf("invalid first signature record")
	}
	rec2, err := d.Sign([]byte("data2"))
	if err != nil {
		t.Fatalf("Sign2: %v", err)
	}
	if rec2.Counter != 1 || rec2.Signature == "" || rec2.SignedData == "" {
		t.Fatalf("invalid second signature record")
	}
}

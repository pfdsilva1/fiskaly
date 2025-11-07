package crypto

import (
    "testing"
)

func TestAlgorithmsInitAndUsage_HappyPath(t *testing.T) {
    t.Parallel()

    alg := "ECDSA"
    pub, priv, err := GenerateKeyPair(alg)
    if err != nil {
        t.Fatalf("GenerateKeyPair(%s) error: %v", alg, err)
    }
    if pub == "" || priv == "" {
        t.Fatalf("GenerateKeyPair(%s) returned empty keys", alg)
    }

    sig, err := Sign(alg, priv, []byte("data"))
    if err != nil {
        t.Fatalf("Sign(%s) error: %v", alg, err)
    }
    if len(sig) == 0 {
        t.Fatalf("Sign(%s) returned empty signature", alg)
    }
}



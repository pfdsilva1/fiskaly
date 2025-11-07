package crypto

import "testing"

func TestECCSigner_HappyPath(t *testing.T) {
    t.Parallel()
    kp, err := (&ECCGenerator{}).Generate()
    if err != nil {
        t.Fatalf("generate: %v", err)
    }
    pub, priv, err := NewECCMarshaler().Encode(kp)
    if err != nil || len(pub) == 0 || len(priv) == 0 {
        t.Fatalf("encode: %v", err)
    }
    sig, err := NewECCSigner().Sign(string(priv), []byte("hello"))
    if err != nil || len(sig) == 0 {
        t.Fatalf("sign: %v", err)
    }
}



package crypto

import "testing"

func TestECCGenerator_Generate_HappyPath(t *testing.T) {
    t.Parallel()
    g := &ECCGenerator{}
    kp, err := g.Generate()
    if err != nil {
        t.Fatalf("Generate error: %v", err)
    }
    if kp == nil || kp.Private == nil || kp.Public == nil {
        t.Fatalf("invalid key pair")
    }
}



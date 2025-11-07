package crypto

import (
	"fmt"
	"strings"
)

func init() {
	RegisterAlgorithmStrategy("RSA", AlgorithmStrategy{
		keyPairGenerator: func() (string, string, error) {
			keyPair, err := (&RSAGenerator{}).Generate()
			if err != nil {
				return "", "", err
			}

			pub, priv, err := NewRSAMarshaler().Marshal(keyPair)
			return string(pub), string(priv), err
		},
		signer: func() (Signer, error) {
			return NewRSASigner(), nil
		},
	})

	RegisterAlgorithmStrategy("ECDSA", AlgorithmStrategy{
		keyPairGenerator: func() (string, string, error) {
			keyPair, err := (&ECCGenerator{}).Generate()
			if err != nil {
				return "", "", err
			}
			pub, priv, err := NewECCMarshaler().Encode(keyPair)
			return string(pub), string(priv), err
		},
		signer: func() (Signer, error) {
			return NewECCSigner(), nil
		},
	})
}

type KeyPairGenerator func() (publicPEM string, privatePEM string, err error)
type SignerFactory func() (Signer, error)

type AlgorithmStrategy struct {
	keyPairGenerator KeyPairGenerator
	signer           SignerFactory
}

var strategies = map[string]AlgorithmStrategy{}

func RegisterAlgorithmStrategy(algorithm string, strategy AlgorithmStrategy) {
	strategies[strings.ToUpper(algorithm)] = strategy
}

func GenerateKeyPair(algorithm string) (string, string, error) {
	strategy, ok := strategies[strings.ToUpper(algorithm)]
	if !ok {
		return "", "", fmt.Errorf("unsupported algorithm %q", algorithm)
	}

	return strategy.keyPairGenerator()
}

func Sign(algorithm string, privateKey string, data []byte) ([]byte, error) {
	strategy, ok := strategies[strings.ToUpper(algorithm)]
	if !ok {
		return nil, fmt.Errorf("unsupported algorithm %q", algorithm)
	}

	signer, err := strategy.signer()
	if err != nil {
		return nil, fmt.Errorf("failed to get signer: %w", err)
	}

	return signer.Sign(privateKey, data)
}

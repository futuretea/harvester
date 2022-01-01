package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"golang.org/x/crypto/ssh"
)

func GeneratePrivateKey(bitSize int) (*rsa.PrivateKey, error) {
	__traceStack()

	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}
	if err = privateKey.Validate(); err != nil {
		return nil, err
	}
	return privateKey, nil
}

func EncodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	__traceStack()

	privateDER := x509.MarshalPKCS1PrivateKey(privateKey)
	privateBlock := pem.Block{
		Type:		"RSA PRIVATE KEY",
		Headers:	nil,
		Bytes:		privateDER,
	}
	return pem.EncodeToMemory(&privateBlock)
}

func GeneratePublicKey(publicKey *rsa.PublicKey) ([]byte, error) {
	__traceStack()

	publicRsaKey, err := ssh.NewPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	return ssh.MarshalAuthorizedKey(publicRsaKey), nil
}

package tls

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
)

func ValidateServingBundle(data []byte) error {
	__traceStack()

	var exists bool

	i := 0

	for containsPEMHeader(data) {
		var block *pem.Block
		block, data = pem.Decode(data)
		if block == nil {
			return errors.New("failed to parse PEM block")
		}
		if block.Type != "CERTIFICATE" {
			return fmt.Errorf("unexpected block type '%s'", block.Type)
		}
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return err
		}

		if i == 0 && !hasCommonName(cert) && !hasSubjectAltNames(cert) {
			return errors.New("certificate has no common name or subject alt name")
		}

		exists = true
		i++
	}

	if !exists {
		return errors.New("failed to locate certificate")
	}

	return nil
}

func ValidateCABundle(data []byte) error {
	__traceStack()

	var exists bool

	for containsPEMHeader(data) {
		var block *pem.Block
		block, data = pem.Decode(data)
		if block == nil {
			return errors.New("failed to parse PEM block")
		}
		if block.Type != "CERTIFICATE" {
			return fmt.Errorf("unexpected block type '%s'", block.Type)
		}

		exists = true
	}

	if !exists {
		return errors.New("failed to locate certificate")
	}
	return nil
}

func ValidatePrivateKey(data []byte) error {
	__traceStack()

	var keys int

	for containsPEMHeader(data) {
		var block *pem.Block
		block, data = pem.Decode(data)
		if block == nil {
			return errors.New("failed to parse PEM block")
		}
		switch block.Type {
		case "PRIVATE KEY":
			if _, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
				return err
			}
			keys++
		case "RSA PRIVATE KEY":
			if _, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
				return err
			}
			keys++
		case "EC PRIVATE KEY":
			if _, err := x509.ParseECPrivateKey(block.Bytes); err != nil {
				return err
			}
			keys++
		case "EC PARAMETERS":

		default:
			return fmt.Errorf("unexpected block type '%s'", block.Type)
		}
	}

	switch keys {
	case 0:
		return errors.New("failed to locate private key")
	case 1:
		return nil
	default:
		return errors.New("multiple private keys")
	}
}

func hasCommonName(c *x509.Certificate) bool {
	__traceStack()

	return strings.TrimSpace(c.Subject.CommonName) != ""
}

func hasSubjectAltNames(c *x509.Certificate) bool {
	__traceStack()

	return len(c.DNSNames) > 0 || len(c.IPAddresses) > 0
}

func containsPEMHeader(data []byte) bool {
	__traceStack()

	start := bytes.Index(data, []byte("-----BEGIN"))
	if start == -1 {
		return false
	}

	end := bytes.Index(data[start+10:], []byte("-----"))
	if end == -1 {
		return false
	}

	if bytes.Contains(data[start:start+end], []byte("\n")) {
		return false
	}

	return true
}

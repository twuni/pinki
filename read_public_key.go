package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
)

const (
	PublicKey = "PUBLIC KEY"
)

func readPublicKey(in io.Reader) (*ecdsa.PublicKey, error) {
	buffer, err := io.ReadAll(in)

	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(buffer)

	if block == nil || block.Type != PublicKey {
		return nil, errors.New("Expected a PEM-encoded public key")
	}

	key, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	return key.(*ecdsa.PublicKey), nil
}

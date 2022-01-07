package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
)

const (
	PrivateKey = "PRIVATE KEY"
)

func readPrivateKey(in io.Reader) (*ecdsa.PrivateKey, error) {
	buffer, err := io.ReadAll(in)

	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(buffer)

	if block == nil || block.Type != PrivateKey {
		return nil, errors.New("Expected a PEM-encoded private key")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	return key.(*ecdsa.PrivateKey), nil
}

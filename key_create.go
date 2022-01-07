package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"io"
)

func createPrivateKey(out io.Writer) error {
	key, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)

	if err != nil {
		return err
	}

	der, err := x509.MarshalPKCS8PrivateKey(key)

	if err != nil {
		return err
	}

	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: der,
	}

	return pem.Encode(out, block)
}

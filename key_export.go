package main

import (
	"crypto/x509"
	"encoding/pem"
	"io"
)

func exportPublicKey(in io.Reader, out io.Writer) error {
	key, err := readPrivateKey(in)

	if err != nil {
		return err
	}

	der, err := x509.MarshalPKIXPublicKey(key.Public())

	if err != nil {
		return err
	}

	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: der,
	}

	return pem.Encode(out, block)
}

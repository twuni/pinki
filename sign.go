package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/pem"
	"io"
)

func sign(privateKeyAsString string, in io.Reader, out io.Writer) error {
	privateKey, err := readPrivateKey(bytes.NewBufferString(privateKeyAsString))

	if err != nil {
		return err
	}

	digest := sha256.New()

	_, err = io.Copy(digest, in)

	if err != nil {
		return err
	}

	signature, err := ecdsa.SignASN1(rand.Reader, privateKey, digest.Sum(nil))

	if err != nil {
		return err
	}

	block := &pem.Block{
		Type:  "SIGNATURE",
		Bytes: signature,
	}

	return pem.Encode(out, block)
}

package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"errors"
	"io"
)

func verify(publicKeyAsString string, signatureAsString string, in io.Reader, out io.Writer) error {
	publicKey, err := readPublicKey(bytes.NewBufferString(publicKeyAsString))

	if err != nil {
		return err
	}

	signature, err := readSignature(bytes.NewBufferString(signatureAsString))

	if err != nil {
		return err
	}

	digest := sha256.New()

	_, err = io.Copy(digest, in)

	if err != nil {
		return err
	}

	isVerified := ecdsa.VerifyASN1(publicKey, digest.Sum(nil), signature)

	if !isVerified {
		return errors.New("Bad Signature")
	}

	_, err = out.Write([]byte("OK\n"))

	return err
}

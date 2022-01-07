package main

import (
	"encoding/pem"
	"errors"
	"io"
)

const (
	Signature = "SIGNATURE"
)

func readSignature(in io.Reader) ([]byte, error) {
	buffer, err := io.ReadAll(in)

	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(buffer)

	if block == nil || block.Type != Signature {
		return nil, errors.New("Expected a PEM-encoded signature")
	}

	return block.Bytes, nil
}

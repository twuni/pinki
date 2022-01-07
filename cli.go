package main

import (
	"io"
)

func cli(args []string, in io.Reader, out io.Writer) error {
	if len(args) < 1 {
		return help(out)
	}

	switch args[0] {
	case "help":
		return help(out)
	case "key":
		if len(args) < 2 {
			return help(out)
		}

		switch args[1] {
		case "create":
			return createPrivateKey(out)
		case "export":
			return exportPublicKey(in, out)
		default:
			return help(out)
		}
	case "sign":
		if len(args) < 2 {
			return help(out)
		}

		return sign(args[1], in, out)
	case "verify":
		if len(args) < 3 {
			return help(out)
		}

		return verify(args[1], args[2], in, out)
	}

	return help(out)
}

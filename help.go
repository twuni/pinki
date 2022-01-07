package main

import (
	"errors"
	"io"
)

func help(out io.Writer) error {
	return errors.New(`USAGE: pinki <command>

EXAMPLES

  $ pinki help
  Display this help message.

  $ pinki key create > private.pem
  Generate a new private key and write the result to "private.pem".

  $ pinki key export < private.pem > public.pem
  Extract the public key from "private.pem" and write the result to "public.pem".

  $ pinki sign "$(cat private.pem)" < package-1.0.0.tgz > package-1.0.0.tgz.asc
  Sign "package-1.0.0.tgz" using the private key from "private.pem" and write the result to "package-1.0.0.tgz.asc".

  $ pinki verify "$(cat public.pem)" "$(cat package-1.0.0.tgz.asc)" < package-1.0.0.tgz
  Verify the signature in "package-1.0.0.tgz.asc" of "package-1.0.0.tgz" using the public key from "public.pem".
`)
}

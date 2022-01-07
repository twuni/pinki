package main

import (
	"io"
)

func keyHelp(out io.Writer) error {
	out.Write([]byte("USAGE: pinki key create|export|help\n"))
	return nil
}

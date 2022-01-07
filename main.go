package main

import (
	"fmt"
	"os"
)

func main() {
	err := cli(os.Args[1:], os.Stdin, os.Stdout)

	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

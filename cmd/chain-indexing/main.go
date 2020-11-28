package main

import (
	"fmt"
	"os"
)

func main() {
	if err := CliApp(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

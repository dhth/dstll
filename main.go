package main

import (
	"os"

	"github.com/dhth/dstll/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

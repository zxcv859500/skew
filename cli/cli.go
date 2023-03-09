package main

import (
	"github.com/zxcv859500/skew/cmd"
	"os"
)

func main() {
	command := cmd.NewDefaultSkewCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

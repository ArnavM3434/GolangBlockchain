package main

import (
	"os"

	"github.com/ArnavM3434/GolangBlockchain/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
}

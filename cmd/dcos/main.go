package main

import (
	"os"

	"github.com/dcos/dcos-cli/pkg/cli"
	"github.com/dcos/dcos-cli/pkg/cmd"
)

func main() {
	if err := cmd.NewDCOSCommand(cli.DefaultContext()).Execute(); err != nil {
		os.Exit(1)
	}
}

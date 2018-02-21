package cmd

import (
	"fmt"
	"strings"

	"github.com/dcos/dcos-cli/pkg/config"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MaximumNArgs(1),
	RunE: runConfigShowCmd,
}

func runConfigShowCmd(cmd *cobra.Command, args []string) error {
	// this is a dummy path for testing purpose
	cfg, err := config.FromPath("/home/bamarni/.dcos/clusters/456934a6-fc1f-4d52-9419-a4c722abe2da/dcos.toml")
	if err != nil {
		return err
	}

	if len(args) == 1 {
		key := strings.Split(args[0], ".")
		out, err := cfg.StringAt(key)
		if err != nil {
			return err
		}
		fmt.Print(out)
	} else {
		fmt.Print(cfg.String())
	}

	return nil
}

func init() {
	configCmd.AddCommand(configShowCmd)
}

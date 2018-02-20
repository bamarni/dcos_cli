package cmd

import (
//	"bufio"
	"fmt"
//	"strings"

	"github.com/dcos/dcos-cli/pkg/config"
//	"github.com/dcos/dcos-cli/pkg/config/tomlutil"
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
	cfg, err := config.FromPath(config.TomlPath)
	if err != nil {
		return err
	}

	fmt.Print(cfg)

	return nil

/*
	// load the TOML config as a tree
	cfg, err = config.FromPath(config.TomlPath)
	if err != nil {
		return err
	}

	// when an argument is passed, generate from it a slice of the different key sections :
	//   "core"          -> []string{"core"}
	//   "core.dcos_url" -> []string{"core", "dcos_url"}
	keys := []string{}
	if len(args) == 1 {
		keys = strings.Split(args[0], ".")
	}

	// write the tree map into a buffer on top of stdout
	w := bufio.NewWriter(os.Stdout)
	if err := tomlutil.WriteTreeTo(w, cfg.Tree, keys); err != nil {
		return err
	}

	// no error occured, print the configs
	return w.Flush()
	*/
}

func init() {
	configCmd.AddCommand(configShowCmd)
}

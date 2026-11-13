package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"razor/pkg/config"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init project",
	Long:  "Init project structure, create Razor.toml and src/ folder",
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Create("Razor.toml")
		if err != nil {
			fmt.Printf("Failed to create Razor.toml: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()

		pwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Failed to get current directory: %v\n", err)
			os.Exit(1)
		}

		cfg := config.Config{
			Package: config.TomlPackage{
				Name:    filepath.Base(pwd),
				Version: "0.1.0",
			},
			Conf: config.TomlConfig{
				Compiler: "clang",
				Standard: 11,
			},
		}

		encoder := toml.NewEncoder(f)
		encoder.Indent = ""

		if err := encoder.Encode(cfg); err != nil {
			fmt.Printf("Failed to create default Razor.toml: %v\n", err)
			os.Exit(1)
		}
	},
}

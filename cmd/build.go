package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"razor/pkg/builder"
	"razor/pkg/config"
	"razor/pkg/workspace"

	"github.com/spf13/cobra"
)

func filterSources(files []string) []string {
	var sources []string
	for _, v := range files {
		if filepath.Ext(v) == ".c" {
			sources = append(sources, v)
		}
	}

	return sources
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build project",
	Long:  "Build project as executable into build/ folder",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Parse()
		if err != nil {
			fmt.Printf("Failed to parse Razor.toml: %v\n", err)
			os.Exit(1)
		}

		files, err := workspace.FindSources()
		if err != nil {
			fmt.Printf("Failed to find source files: %v\n", err)
			os.Exit(1)
		}

		b := builder.New(cfg, filterSources(files))

		if err := b.Build(); err != nil {
			fmt.Printf("Failed to build project: %v\n", err)
			os.Exit(1)
		}
	},
}

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"razor/pkg/builder"
	"razor/pkg/config"
	"razor/pkg/workspace"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run executable",
	Long:  "Build project and run executable",
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

		out, err := exec.Command(filepath.Join(".", "build", cfg.Package.Name)).CombinedOutput()
		if len(out) > 0 {
			fmt.Print(string(out))
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

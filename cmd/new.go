package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [workdir]",
	Short: "Create directory and init project",
	Long:  "Create directory with name from args, and init project",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Failed to get current directory name: %v\n", err)
			os.Exit(1)
		}

		absolute, err := filepath.Abs(pwd)
		if err != nil {
			fmt.Printf("Failed to get absolute path of %s: %v\n", pwd, err)
			os.Exit(1)
		}

		workpath := args[0]

		if err := os.Mkdir(workpath, 0755); err != nil {
			fmt.Printf("Failed to create %s: %v\n", args[0], err)
			os.Exit(1)
		}

		if err := os.Chdir(workpath); err != nil {
			fmt.Printf("Failed to create %s directory: %v\n", args[0], err)
			os.Exit(1)
		}

		createRazorToml()
		createSrcDirectory()

		os.Chdir(absolute)
	},
}

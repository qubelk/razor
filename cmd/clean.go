package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove build folder",
	Long:  "Clean project build by remove build folder",
	Run: func(cmd *cobra.Command, args []string) {
		if err := os.RemoveAll("build"); err != nil {
			fmt.Printf("Failed to remove build folder: %v\n", err)
			os.Exit(1)
		}
	},
}

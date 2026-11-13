package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "razor",
	Short: "Simple build system for C",
	Long:  "Simple Cargo style build system for C",
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

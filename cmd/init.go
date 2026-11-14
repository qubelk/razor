package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"razor/pkg/config"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
)

func createRazorToml() {
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
}

func createSrcDirectory() {
	if err := os.MkdirAll("src", 0755); err != nil {
		fmt.Printf("Failed to create src/ directory: %v\n", err)
		os.Exit(1)
	}

	if err := os.Chdir("src"); err != nil {
		fmt.Printf("Failed to create %s file: %v\n", filepath.Join("src", "main.c"), err)
		os.Exit(1)
	}

	f, err := os.Create("main.c")
	if err != nil {
		fmt.Printf("Failed to create %s file: %v\n", filepath.Join("src", "main.c"), err)
		os.Exit(1)
	}
	defer f.Close()

	data := `#include <stdio.h>
	
	int main(void) {
		printf("Hello, World!\n");
		return 0;
	}
	`

	if _, err := f.WriteString(data); err != nil {
		fmt.Printf("Failed to write into %s file: %v\n", filepath.Join("src", "main.c"), err)
		os.Exit(1)
	}
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init project",
	Long:  "Init project structure, create Razor.toml and src/ folder",
	Run: func(cmd *cobra.Command, args []string) {
		createRazorToml()
		createSrcDirectory()
	},
}

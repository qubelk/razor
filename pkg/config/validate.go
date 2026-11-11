package config

import (
	"fmt"
	"slices"
)

func (c *Config) validateCompiler() error {
	supportedCompilers := []string{
		"gcc",
		"clang",
	}

	if !slices.Contains(supportedCompilers, c.Conf.Compiler) {
		return fmt.Errorf("%s is not supported compiler. Now supported only gcc and clang",
			c.Conf.Compiler)
	}

	return nil
}

func (c *Config) validateStandard() error {
	standards := []int{
		89,
		99,
		11,
		17,
		23,
	}

	if !slices.Contains(standards, c.Conf.Standard) {
		return fmt.Errorf("unknown standart %d", c.Conf.Standard)
	}

	return nil
}

func (c *Config) validate() error {
	if err := c.validateCompiler(); err != nil {
		return err
	}

	if err := c.validateStandard(); err != nil {
		return err
	}

	return nil
}

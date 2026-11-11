package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Package tomlPackage `toml:"package"`
	Conf    tomlConfig  `toml:"config"`
}

func Parse() (*Config, error) {
	data, err := os.ReadFile("Razor.toml")
	if err != nil {
		return nil, err
	}

	var conf Config
	err = toml.Unmarshal(data, &conf)

	return &conf, err
}

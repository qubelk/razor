package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Package TomlPackage `toml:"package"`
	Conf    TomlConfig  `toml:"config"`
}

func Parse() (*Config, error) {
	var conf Config
	_, err := toml.DecodeFile("Razor.toml", &conf)
	if err != nil {
		return nil, err
	}

	if err := conf.validate(); err != nil {
		return nil, err
	}

	return &conf, nil
}

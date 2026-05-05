package models

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Unit    Unit    `toml:"Unit"`
	Users   []User  `toml:"Users"`
	Classes []Class `toml:"Classes"`
}

func GetConfig() Config {
	var cfg Config
	data, err := os.ReadFile("config.toml")
	if err != nil {
		panic("Failed to get config file")
	}
	_, err = toml.Decode(string(data), &cfg)
	if err != nil {
		panic("Failed to decode toml file")
	}
	return cfg
}

package config

import (
	"fmt"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Database DB `koanf:"database"`
}

type DB struct {
	Host     string `koanf:"host"`
	Port     uint16 `koanf:"port"`
	User     string `koanf:"user"`
	Password string `koanf:"password"`
	Database string `koanf:"db"`
}

func LoadConfig(path string) (*Config, error) {
	parser := yaml.Parser()
	reader := koanf.New(".")

	err := reader.Load(file.Provider(path), parser)
	if err != nil {
		return nil, fmt.Errorf("config: failed to load config: %w", err)
	}

	var config Config
	err = reader.Unmarshal("", &config)
	if err != nil {
		return nil, fmt.Errorf("config: failed to unmarshal config: %w", err)
	}

	return &config, nil
}

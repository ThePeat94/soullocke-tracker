package config

import (
	"fmt"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Database Database `koanf:"database"`
}

type Database struct {
	Host     string `koanf:"host"`
	Port     uint16 `koanf:"port"`
	User     string `koanf:"user"`
	Password string `koanf:"password"`
	Database string `koanf:"db"`
}

func (dbConfig *Database) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
	)
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

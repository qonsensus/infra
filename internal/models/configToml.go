package models

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type DatabaseConfig struct {
	Username        string `toml:"username"`
	Password        string `toml:"password"`
	DefaultDatabase string `toml:"default_database"`
}

type ConfigToml struct {
	Version  string         `toml:"version"`
	Database DatabaseConfig `toml:"database"`
	basePath string         `toml:"-"`
}

func NewConfigToml(version, basePath string) *ConfigToml {
	return &ConfigToml{
		Version:  version,
		basePath: basePath,
	}
}

// LoadConfigTomlFromFile loads the config.toml file from the specified basePath and decodes it into a ConfigToml struct. It returns an error if the file cannot be read or decoded.
func LoadConfigTomlFromFile(basePath string) (*ConfigToml, error) {
	var config ConfigToml
	fullPath := filepath.Join(basePath, "config.toml")
	// decode existing config file
	_, err := toml.DecodeFile(fullPath, &config)
	if err != nil {
		return nil, err
	}
	config.basePath = basePath
	return &config, nil
}

// ToTomlString encodes the ConfigToml struct into a TOML string. If encoding fails, it returns an empty string.
func (c *ConfigToml) ToTomlString() string {
	var buf bytes.Buffer
	encoder := toml.NewEncoder(&buf)
	err := encoder.Encode(c)
	if err != nil {
		return ""
	}
	return buf.String()
}

// SaveToFile saves the ConfigToml struct to a config.toml file in the basePath directory. It returns an error if writing to the file fails.
func (c *ConfigToml) SaveToFile() error {
	// Ensure the basePath directory exists and the current user has read and write permissions
	if err := os.MkdirAll(c.basePath, 0755); err != nil {
		return err
	}
	buf := bytes.NewBufferString(c.ToTomlString())
	return os.WriteFile(filepath.Join(c.basePath, "config.toml"), buf.Bytes(), 0644)
}

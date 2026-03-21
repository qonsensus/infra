package models

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/qonsensus/infra/internal"
)

type ConfigToml struct {
	ProjectName string `toml:"project_name"`
	Version     string `toml:"version"`
	basePath    string `toml:"-"`
}

// NewConfigToml creates a new ConfigToml instance. If a config.toml file exists at the given basePath, it will be decoded and returned. Otherwise, a new ConfigToml with default values will be returned.
func NewConfigToml(basePath string) *ConfigToml {
	var config ConfigToml
	fullPath := filepath.Join(basePath, "config.toml")
	// return new config with default values if file does not exist
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		config.basePath = basePath
		config.Version = internal.CurrentVersion
		return &config
	}
	// decode existing config file
	_, err := toml.DecodeFile(fullPath, &config)
	if err != nil {
		return nil
	}
	config.basePath = basePath
	return &config
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
	buf := bytes.NewBufferString(c.ToTomlString())
	return os.WriteFile(filepath.Join(c.basePath, "config.toml"), buf.Bytes(), 0644)
}

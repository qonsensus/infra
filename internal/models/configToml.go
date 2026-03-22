package models

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type DatabaseConfig struct {
	Username        string `toml:"username"`
	Password        string `toml:"password"`
	DefaultDatabase string `toml:"default_database"`
	RestartStrategy string `toml:"restart"`
	PersistData     bool   `toml:"persist_data"`
}

type ExposeConfig struct {
	Expose bool `toml:"expose"`
	Port   int  `toml:"port"`
}

type ContainerNames struct {
	Postgres string `toml:"postgres"`
	Frontend string `toml:"frontend"`
	Backend  string `toml:"backend"`
}

type ConfigToml struct {
	Version           string         `toml:"version"`
	BackendPublicURL  string         `toml:"backend_public_url"`
	FrontendPublicURL string         `toml:"frontend_public_url"`
	ExposeDatabase    ExposeConfig   `toml:"expose_database"`
	ExposeBackend     ExposeConfig   `toml:"expose_backend"`
	ExposeFrontend    ExposeConfig   `toml:"expose_frontend"`
	CORSOrigins       []string       `toml:"allowed_cors_origins"`
	ContainerNames    ContainerNames `toml:"container_names"`
	Database          DatabaseConfig `toml:"database"`
	basePath          string         `toml:"-"`
}

func NewConfigToml(version, basePath string) *ConfigToml {
	return &ConfigToml{
		Version:           version,
		basePath:          basePath,
		BackendPublicURL:  fmt.Sprintf("http://%s:%d", "localhost", 7000),
		FrontendPublicURL: fmt.Sprintf("http://%s:%d", "localhost", 7001),

		ExposeBackend: ExposeConfig{
			Expose: true,
			Port:   7000,
		},
		ExposeFrontend: ExposeConfig{
			Expose: true,
			Port:   7001,
		},
		ExposeDatabase: ExposeConfig{
			Expose: false,
			Port:   5432,
		},
		CORSOrigins: []string{
			"http://localhost:7001",
		},
		Database: DatabaseConfig{
			Username:        "admin",
			Password:        "secret",
			DefaultDatabase: "quonsensus",
			PersistData:     true,
			RestartStrategy: "always",
		},
		ContainerNames: ContainerNames{
			Postgres: "quonsensus-postgres",
			Frontend: "quonsensus-frontend",
			Backend:  "quonsensus-backend",
		},
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

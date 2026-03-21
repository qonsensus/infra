package models

import (
	"fmt"

	"github.com/compose-spec/compose-go/v2/types"
)

type BackendService struct {
	types.ServiceConfig
}

func NewBackendService(config *ConfigToml) *BackendService {
	dbHost := "postgres"
	dbPort := "5432"
	return &BackendService{
		ServiceConfig: types.ServiceConfig{
			Image:         fmt.Sprintf("ghcr.io/qonsensus/backend:%s", config.Version),
			ContainerName: config.ContainerNames.Backend,
			Ports: []types.ServicePortConfig{
				{
					Target:    3000,
					Published: fmt.Sprint(config.HostPorts.Backend),
				},
			},
			Environment: types.MappingWithEquals{
				"DB_HOST":     &dbHost,
				"DB_PORT":     &dbPort,
				"DB_USER":     &config.Database.Username,
				"DB_PASSWORD": &config.Database.Password,
				"DB_NAME":     &config.Database.DefaultDatabase,
			},
			Restart: "always",
		},
	}
}

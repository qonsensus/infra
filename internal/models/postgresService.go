package models

import (
	"fmt"

	"github.com/compose-spec/compose-go/v2/types"
)

type PostgresService struct {
	types.ServiceConfig
}

func NewPostgresService(config *ConfigToml) *PostgresService {
	service := &PostgresService{
		ServiceConfig: types.ServiceConfig{
			Image:         "postgres:latest",
			ContainerName: config.ContainerNames.Postgres,
			Environment: types.MappingWithEquals{
				"POSTGRES_USER":     &config.Database.Username,
				"POSTGRES_PASSWORD": &config.Database.Password,
				"POSTGRES_DB":       &config.Database.DefaultDatabase,
			},
			Restart: config.Database.RestartStrategy,
		},
	}
	if config.Database.PersistData {
		service.Volumes = []types.ServiceVolumeConfig{
			{
				Type:   types.VolumeTypeBind,
				Source: "./data/postgres",
				Target: "/var/lib/postgresql",
			},
		}

	}
	if config.ExposeDatabase.Expose == true && config.ExposeDatabase.Port != 0 {
		service.Ports = []types.ServicePortConfig{
			{
				Target:    5432,
				Published: fmt.Sprint(config.ExposeDatabase.Port),
			},
		}
	}
	return service
}

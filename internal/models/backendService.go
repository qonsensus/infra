package models

import (
	"fmt"
	"strings"

	"github.com/compose-spec/compose-go/v2/types"
)

type BackendService struct {
	types.ServiceConfig
}

func NewBackendService(config *ConfigToml) *BackendService {
	dbHost := "postgres"
	dbPort := "5432"
	service := &BackendService{
		ServiceConfig: types.ServiceConfig{
			Image:         fmt.Sprintf("ghcr.io/qonsensus/backend:%s", config.Version),
			ContainerName: config.ContainerNames.Backend,
			Environment: types.MappingWithEquals{
				"DB_HOST":     &dbHost,
				"DB_PORT":     &dbPort,
				"DB_USER":     &config.Database.Username,
				"DB_PASSWORD": &config.Database.Password,
				"DB_NAME":     &config.Database.DefaultDatabase,
				"CORS_ALLOWED_ORIGINS": func() *string {
					if len(config.CORSOrigins) == 0 {
						return nil
					}
					origins := strings.Join(config.CORSOrigins, ",")
					return &origins
				}(),
			},
			Restart: "unless-stopped",
		},
	}
	if config.ExposeBackend.Expose {
		service.Ports = append(service.Ports, types.ServicePortConfig{
			Target:    3000,
			Published: fmt.Sprint(config.ExposeBackend.Port),
		})
	}
	return service
}

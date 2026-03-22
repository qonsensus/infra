package models

import (
	"fmt"

	"github.com/compose-spec/compose-go/v2/types"
)

type FrontendService struct {
	types.ServiceConfig
}

func NewFrontendService(config *ConfigToml) *FrontendService {
	service := &FrontendService{
		ServiceConfig: types.ServiceConfig{
			Image:         fmt.Sprintf("ghcr.io/qonsensus/frontend:%s", config.Version),
			ContainerName: config.ContainerNames.Frontend,
			Environment: types.MappingWithEquals{
				"API_URL": &config.BackendPublicURL,
			},
			Restart: "always",
		},
	}
	if config.ExposeFrontend.Expose {
		service.Ports = append(service.Ports, types.ServicePortConfig{
			Target:    80,
			Published: fmt.Sprint(config.ExposeFrontend.Port),
		})
	}
	return service
}

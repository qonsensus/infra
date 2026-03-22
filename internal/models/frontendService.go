package models

import (
	"fmt"

	"github.com/compose-spec/compose-go/v2/types"
)

type FrontendService struct {
	types.ServiceConfig
}

func NewFrontendService(config *ConfigToml) *FrontendService {
	backendUrl := fmt.Sprintf("http://localhost:%d", config.HostPorts.Backend)
	return &FrontendService{
		ServiceConfig: types.ServiceConfig{
			Image:         fmt.Sprintf("ghcr.io/qonsensus/frontend:%s", config.Version),
			ContainerName: config.ContainerNames.Frontend,
			Ports: []types.ServicePortConfig{
				{
					Target:    80,
					Published: fmt.Sprint(config.HostPorts.Frontend),
				},
			},
			Environment: types.MappingWithEquals{
				"API_URL": &backendUrl,
			},
			Restart: "always",
		},
	}
}

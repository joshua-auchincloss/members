package registry

import (
	"members/common"
	"members/grpc/api/v1/registry"
	"members/service"
	"members/service/core"
)

type (
	RegistryClient = *service.ClientFactory[registry.RegistryClient]
)

var (
	ClientFactory = core.NewClient(
		"registry",
		common.ServiceRegistry,
		registry.NewRegistryClient,
	)
)

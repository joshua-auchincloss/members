package registry

import (
	"members/common"
	"members/grpc/api/v1/registry"
	"members/server"
	"members/service"

	"go.uber.org/fx"
)

type (
	RegistryClient = *service.ClientFactory[registry.RegistryClient]
)

var (
	client_factory RegistryClient = service.NewClientFactory(common.ServiceRegistry, registry.NewRegistryClient)
	ClientFactory                 = fx.Module(
		"registry-client-factory",
		fx.Supply(client_factory),
		server.LoadBalancerFor(common.ServiceRegistry),
	)
)

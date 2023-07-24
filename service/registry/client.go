package registry

import (
	"members/common"
	"members/grpc/api/v1/registry"
	server "members/http"
	"members/service"

	"go.uber.org/fx"
)

var (
	client_factory = service.NewClientFactory(common.ServiceRegistry, registry.NewRegistryClient)
	ClientFactory  = fx.Module(
		"registry-client-factory",
		fx.Supply(client_factory),
		server.LoadBalancerFor(common.ServiceRegistry),
	)
)

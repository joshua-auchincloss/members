package registry

import (
	"members/common"
	"members/grpc/api/v1/registry"
	"members/service"

	"go.uber.org/fx"
)

var (
	client_factory = service.NewClientFactory(common.ServiceRegistry, registry.NewRegistryClient)
	ClientFactory  = fx.Module("registry-client-factory", fx.Provide(client_factory))
)

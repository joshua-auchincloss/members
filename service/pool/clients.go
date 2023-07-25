package pool

import (
	"members/service/admin"
	"members/service/health"
	"members/service/registry"

	"go.uber.org/fx"
)

type (
	ClientPool interface {
		GetHealth() health.HealthClient
		GetAdmin() admin.AdminClient
		GetRegistry() registry.RegistryClient
	}

	client_pool struct {
		health_factory   health.HealthClient
		admin_factory    admin.AdminClient
		registry_factory registry.RegistryClient
	}
)

func create(
	h health.HealthClient,
	a admin.AdminClient,
	r registry.RegistryClient) *client_pool {
	return &client_pool{
		h,
		a,
		r,
	}
}

var (
	Module = fx.Module(
		"client-pool",
		fx.Provide(create),
	)
)

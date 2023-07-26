package core

import (
	"members/common"
	"members/service"

	"go.uber.org/fx"
)

func NewModule[
	ServiceFactory service.ServiceFactory[ServiceKind],
	ServiceKind service.Service,
](
	name string,
	svc common.Service,
	create func() ServiceFactory,
) fx.Option {
	return fx.Module(
		name,
		fx.Provide(
			fx.Annotate(
				create,
				fx.As(new(ServiceFactory)),
			),
		),
		fx.Invoke(
			service.Create[ServiceKind](svc),
		),
	)
}

package core

import (
	"fmt"
	"members/common"
	"members/server"
	"members/service"

	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func NewClient[
	T any,
	F func(ci grpc.ClientConnInterface) T,
](
	name string,
	svc common.Service,
	fu F,
) fx.Option {
	client_factory := service.NewClientFactory(svc, fu)
	return fx.Module(
		fmt.Sprintf("%s-client-factory", name),
		fx.Supply(client_factory),
		server.LoadBalancerFor(svc),
	)
}

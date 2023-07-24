package server

import (
	"members/common"
	"members/config"
	"members/storage"

	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
	"google.golang.org/grpc/resolver"
)

func LoadBalancerFor(svc common.Service) fx.Option {
	return fx.Invoke(func(ctx *cli.Context, prov config.ConfigProvider, store storage.Store) error {
		resolver.Register(&resolverBuilder{svc, store, prov})
		return nil
	})
}

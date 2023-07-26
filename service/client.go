package service

import (
	"context"
	"crypto/tls"
	"members/common"
	"members/config"
	"members/server"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	ClientFactory[T any] struct {
		impl func(prov config.ConfigProvider, args *common.DialArgs) (*T, error)
	}
)

func (f ClientFactory[T]) New(prov config.ConfigProvider, args *common.DialArgs) (*T, error) {
	return f.impl(prov, args)
}

func NewClientFactory[T any](key common.Service, impl func(ci grpc.ClientConnInterface) T) *ClientFactory[T] {
	return &ClientFactory[T]{
		func(prov config.ConfigProvider, args *common.DialArgs) (*T, error) {
			call := []grpc.DialOption{}
			cfg := prov.GetConfig()
			var err error
			var certs *tls.Config
			var cfg_tls *config.ClientTls
			if args.TLS && cfg.Tls.Enabled {
				global := cfg.Members.Global
				cli := cfg.Members.GetClient(key)
				if ct, ok := cli.Trusted[args.DNS]; !ok {
					cfg_tls = &global
				} else {
					cfg_tls = &ct
				}
				certs, err = cfg_tls.Build()
				if err != nil {
					return nil, err
				}
				call = append(call, grpc.WithTransportCredentials(
					credentials.NewTLS(certs),
				))
			} else {
				call = append(call, grpc.WithTransportCredentials(insecure.NewCredentials()))
			}
			dial, err := server.NewClient(key, prov, certs, args)
			if err != nil {
				return nil, err
			}
			call = append(call, grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
				log.Info().Str("address", addr).Msg("dialling...")
				return dial.Dial(ctx, addr)
			}))
			conn, err := server.DialGrpc(context.TODO(), key, args.DNS, call...)
			if err != nil {
				return nil, err
			}
			svc := impl(conn)
			return &svc, nil
		},
	}
}

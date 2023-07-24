package service

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"members/common"
	"members/config"
	server "members/http"
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
			var root_pool *x509.CertPool
			var root_tls *tls.Config
			var cfg_tls *config.Tls
			if args.TLS && cfg.Tls.Enabled {
				cfg_tls = cfg.Tls.GetService(key)
				ca, err := cfg_tls.LoadCA()
				if err != nil {
					return nil, err
				}
				root_pool = ca
				root_tls = &tls.Config{
					RootCAs: root_pool,
				}
				call = append(call, grpc.WithTransportCredentials(
					credentials.NewClientTLSFromCert(root_pool, cfg_tls.ServerName),
				))
			} else {
				call = append(call, grpc.WithTransportCredentials(insecure.NewCredentials()))
			}
			dial, err := server.NewClient(key, prov, root_tls, args)
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

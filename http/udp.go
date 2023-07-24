package server

import (
	"crypto/tls"
	"members/config"
	errs "members/errors"
	"net/http"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
)

type (
	udpServer struct {
		base
	}
)

var (
	_ Server = ((*udpServer)(nil))
)

func NewUDP(
	watcher errs.Watcher,
	cfg *config.Tls,
	root *config.TlsConfig,
	t *tls.Config,
	conf *quic.Config,
	addr string,
	handler http.Handler,
) (Server, error) {
	return &udpServer{
		base: base{
			err: watcher.Subscription(),
			srv: &http3.Server{
				Handler:    handler,
				Addr:       addr,
				QuicConfig: conf,
				TLSConfig:  http3.ConfigureTLSConfig(t),
			},
			cfg:  cfg,
			root: root,
		},
	}, nil
}

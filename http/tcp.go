package server

import (
	"crypto/tls"
	"members/config"
	errs "members/errors"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type (
	tcpServer struct {
		base
	}
)

var (
	_ Server = ((*tcpServer)(nil))
)

func NewTCP(
	watcher errs.Watcher,
	cfg *config.Tls,
	root *config.TlsConfig,
	t *tls.Config,
	addr string,
	handler http.Handler,
) (Server, error) {
	if !root.Enabled {
		handler = h2c.NewHandler(handler, &http2.Server{})
	}
	return &tcpServer{
		base: base{
			err: watcher.Subscription(),
			srv: &http.Server{
				Handler:   handler,
				Addr:      addr,
				TLSConfig: t,
			},
			cfg:  cfg,
			tls:  t,
			root: root,
		},
	}, nil
}

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

func udp_starter(
	u Server,
) error {
	srv := u.GetServer().(*http3.Server)
	ln, err := quic.ListenAddrEarly(
		srv.Addr,
		srv.TLSConfig,
		srv.QuicConfig,
	)
	if err != nil {
		return err
	}
	u.WithCloser(func() error {
		return ln.Close()
	})
	return srv.ServeListener(ln)
}

func NewUDP(
	watcher errs.Watcher,
	cfg *config.ServerTls,
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
			cfg:     cfg,
			root:    root,
			starter: starter_for(udp_starter),
		},
	}, nil
}

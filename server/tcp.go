package server

import (
	"crypto/tls"
	"errors"
	"members/config"
	errs "members/errors"
	"net"
	"net/http"
	"strings"

	"github.com/valyala/fasthttp/reuseport"
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

func tcp_starter(t Server) error {
	srv, ok := t.GetServer().(*http.Server)
	if !ok {
		return errors.Join(errs.ErrCastInvalid, errs.ErrServerStarter)
	}
	ip := net.ParseIP(strings.Split(srv.Addr, ":")[0])
	nw := "tcp4"
	if ip != nil {
		ipv4 := ip.DefaultMask() != nil || ip.IsUnspecified()
		if !ipv4 {
			nw = "tcp6"
		}
	}
	ln, err := reuseport.Listen(nw, srv.Addr)
	if err != nil {
		return err
	}
	if t.TlsEnabled() {
		tl := t.GetTLS()
		ln = tls.NewListener(ln, tl)
	}
	t.WithCloser(func() error {
		return ln.Close()
	})
	return srv.Serve(ln)
}

func NewTCP(
	watcher errs.Watcher,
	cfg *config.ServerTls,
	root *config.TlsConfig,
	t *tls.Config,
	addr string,
	handler http.Handler,
) (Server, error) {
	// if !root.Enabled {
	handler = h2c.NewHandler(handler, &http2.Server{})
	// }
	return &tcpServer{
		NewBase(
			watcher.Subscription(),
			&http.Server{
				Handler:   handler,
				Addr:      addr,
				TLSConfig: t,
			},
			t,
			cfg,
			root,
			starter_for(tcp_starter),
		),
	}, nil
}

package server

import (
	"context"
	"crypto/tls"
	"errors"
	"members/config"
	errs "members/errors"
	"net"
	"net/http"
	"time"

	"github.com/quic-go/quic-go"
)

type (
	Server interface {
		Start() chan error
		ReportErr(err error)
		GetServer() invariant
		TlsEnabled() bool
		GetTLS() *tls.Config
		GetConfig() *config.Tls
		Stop(context.Context, time.Duration) error
	}

	invariant interface {
		Close() error
		ListenAndServe() error
		ListenAndServeTLS(string, string) error
	}
)

var (
	timeoutChecks time.Duration = 4
)

func New(
	prov config.ConfigProvider,
	t *config.Tls,
	watcher errs.Watcher,
	addr string,
	handler http.Handler,
) (Server, error) {
	cfg := prov.GetConfig()
	var tlscfg *tls.Config
	var err error
	if t != nil && cfg.Tls.Enabled {
		tlscfg, err = t.Build()
		if err != nil {
			return nil, err
		}
	}

	switch cfg.Members.Protocol {
	case "tcp":
		return NewTCP(
			watcher,
			t,
			cfg.Tls,
			tlscfg,
			addr,
			handler,
		)
	case "udp":
		if tlscfg == nil {
			tlscfg = &tls.Config{}
		}
		udpcfg := &quic.Config{
			RequireAddressValidation: func(a net.Addr) bool { return cfg.Tls.Validation },
		}
		return NewUDP(
			watcher,
			t,
			cfg.Tls,
			tlscfg,
			udpcfg,
			addr,
			handler,
		)
	}

	return nil, nil
}

func starter(srv Server) chan error {
	ch := make(chan error, 1)
	go func(ch chan error) {
		if srv.TlsEnabled() {
			cf := srv.GetConfig()
			ch <- srv.GetServer().ListenAndServeTLS(cf.CertFile, cf.KeyFile)
		} else {
			ch <- srv.GetServer().ListenAndServe()
		}
	}(ch)
	return ch
}

func stopper(srv Server, ctx context.Context, ttc time.Duration) error {
	// check n times
	poll := time.NewTicker(ttc / timeoutChecks)
	ch := make(chan error, 1)
	go func() {
		err := srv.GetServer().Close()
		if err != nil {
			ch <- err
		}
	}()
	ctx, clean := context.WithTimeout(ctx, ttc)
	defer clean()
	for {
		select {
		case err := <-ch:
			return err
		case <-ctx.Done():
			return errors.New("timed out for shutdown")
		case <-poll.C:
		}
	}
}

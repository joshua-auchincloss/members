package service

import (
	"context"
	"members/server"
	"net/http"
	"time"
)

func (h *BaseService) GrpcStarter(
	addr, path string, handler http.Handler) (func(ctx context.Context) error, error) {
	cfg := h.prov.GetConfig()
	tc := cfg.Tls.GetService(h.key)
	srv, err := server.New(
		h.prov,
		tc,
		h.watcher,
		addr,
		handler,
	)
	if err != nil {
		return nil, err
	}
	ch := srv.Start()
	h.WithLoop(
		func(ctx context.Context) error {
			sub, cancel := context.WithTimeout(ctx, time.Millisecond*50)
			defer cancel()
			select {
			case <-sub.Done():
			case err := <-ch:
				h.logger.Err(err).Msg("here")
				return err
			}
			return nil
		},
	)
	return func(ctx context.Context) error {
		return srv.Stop(ctx, time.Second*5)
	}, nil
}

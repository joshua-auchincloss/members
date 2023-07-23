package service

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func GrpcStarter(addr, path string, handler http.Handler) func(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.Handle(path, handler)
	if err := http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Print("err")
		return nil
	}
	return func(ctx context.Context) error {

		return nil
	}
}

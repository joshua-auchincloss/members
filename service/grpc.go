package service

import (
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func GrpcStarter(addr, path string, handler http.Handler) error {
	mux := http.NewServeMux()
	mux.Handle(path, handler)
	if err := http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Print("err")
		panic(err)
	}
	return nil
}

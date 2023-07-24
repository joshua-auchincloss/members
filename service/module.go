package service

// import (
// 	"context"
// 	"fmt"
// 	"members/common"
// 	"members/config"
// 	"members/storage"
// 	"net/http"

// 	"github.com/bufbuild/connect-go"
// 	"github.com/rs/zerolog/log"
// 	"go.uber.org/fx"
// )

// type (
// 	grpcService[Embed any] struct {
// 		BaseService
// 		init  func(svc Embed, opts ...connect.HandlerOption) (string, http.Handler)
// 		store storage.Store
// 	}
// 	grpcFactory[Handler any, Embed any] struct {
// 		init func(svc Handler, opts ...connect.HandlerOption) (string, http.Handler)
// 	}
// )

// func GrpcModule[Handler any, Embed any](key common.Service,
// 	init func(svc Handler, opts ...connect.HandlerOption) (string, http.Handler),
// ) fx.Option {
// 	nm := service_keys[key]
// 	var (
// 		_ ServiceFactory[*grpcService[Handler, Embed]] = ((*grpcFactory[Handler, Embed])(nil))
// 	)
// 	var ve func(svc Embed, opts ...connect.HandlerOption) (string, http.Handler)
// 	return fx.Module(
// 		fmt.Sprintf("%s-service", nm),
// 		fx.Supply(fx.Annotate(
// 			init,
// 			fx.As(ve),
// 		)),
// 		fx.Provide(
// 			fx.Annotate(
// 				CreateGrpcFactory[Handler, Embed],
// 				fx.As(new(ServiceFactory[*grpcService[Handler, Embed]])),
// 			),
// 		),
// 		fx.Invoke(
// 			Create[*grpcService[Handler, Embed]](key),
// 		),
// 	)
// }

// func CreateGrpcFactory[Handler, Embed any](
// 	svc *grpcFactory[Handler, Embed],
// 	init func(svc Handler, opts ...connect.HandlerOption) (string, http.Handler),
// ) ServiceFactory[*grpcService[Handler, Embed]] {
// 	return &grpcFactory[Handler, Embed]{init}
// }

// func (h *grpcFactory[Handler, Embed]) CreateService(
// 	cfg config.ConfigProvider,
// 	store storage.Store) *grpcService[Handler, Embed] {
// 	return &grpcService[Handler, Embed]{
// 		store: store,
// 		init:  h.init,
// 	}
// }

// func (h *grpcService[Handler, Embed]) WithBase(base BaseService) {
// 	h.BaseService = base
// }

// func (h *grpcService[Handler, Embed]) Start(ctx context.Context) error {
// 	pth, handle := h.init(h.h)
// 	clean, err := h.GrpcStarter(h.GetService(), pth, handle)
// 	if err != nil {
// 		return err
// 	}
// 	go h.LoopedStarter(ctx, clean)
// 	return nil
// }

// func (h *grpcService[Handler, Embed]) Stop(ctx context.Context) error {
// 	log.Print("registry stopping")
// 	return nil
// }

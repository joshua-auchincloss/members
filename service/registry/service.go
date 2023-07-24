package registry

import (
	"context"
	"members/grpc/api/v1/registry/registryconnect"
	registry "members/grpc/api/v1/registry/registryconnect"
	"members/service"
	"members/storage"

	"github.com/rs/zerolog/log"
)

type (
	registryService struct {
		registry.UnimplementedRegistryHandler
		service.BaseService
		store storage.Store
	}
)

var (
	_ service.Service = ((*registryService)(nil))
)

func (h *registryService) WithBase(base service.BaseService) {
	h.BaseService = base
}

func (h *registryService) Start(ctx context.Context) error {
	pth, handle := registryconnect.NewRegistryHandler(h)
	clean, err := h.GrpcStarter(h.GetService(), pth, handle)
	if err != nil {
		return err
	}
	go h.LoopedStarter(ctx, clean)
	return nil
}

func (h *registryService) Stop(ctx context.Context) error {
	log.Print("registry stopping")
	return nil
}

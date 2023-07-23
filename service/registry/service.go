package registry

import (
	"context"
	"log"
	"members/grpc/api/v1/registry/registryconnect"
	registry "members/grpc/api/v1/registry/registryconnect"
	"members/service"
	"members/storage"
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

func (h *registryService) Start(ctx context.Context) {
	pth, handle := registryconnect.NewRegistryHandler(h)
	go service.GrpcStarter(h.GetService(), pth, handle)
	h.LoopedStarter(ctx)
}

func (h *registryService) Stop() error {
	log.Print("registry stopping")
	return nil
}

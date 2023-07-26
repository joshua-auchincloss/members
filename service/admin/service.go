package admin

import (
	"context"
	"members/grpc/api/v1/admin/adminconnect"
	admin "members/grpc/api/v1/admin/adminconnect"
	"members/grpc/api/v1/common"
	"members/service"
	"members/storage/base"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/rs/zerolog/log"
)

type (
	adminService struct {
		admin.UnimplementedAdminHandler
		service.BaseService
		store base.BaseStore
	}
)

var (
	_ service.Service = ((*adminService)(nil))
)

func (h *adminService) WithBase(base service.BaseService) {
	h.BaseService = base
}

func (h *adminService) Start(ctx context.Context) error {
	pth, handle := adminconnect.NewAdminHandler(h)
	clean, err := h.GrpcStarter(h.GetService(), pth, handle)
	if err != nil {
		return err
	}
	go h.LoopedStarter(ctx, clean)
	return nil
}

func (h *adminService) Stop(ctx context.Context) error {
	log.Print("admin stopping")
	return nil
}

func (h *adminService) DescribeCluster(ctx context.Context, req *connect.Request[common.Empty], strm *connect.ServerStream[common.Member]) error {
	memb, err := h.store.GetMembers(ctx)
	if err != nil {
		return err
	}
	log.Info().Interface("scanned", memb).Send()
	for _, clust := range memb {
		mb := &common.Member{
			Dns:        clust.Dns,
			Service:    clust.Service,
			Address:    clust.PublicAddress,
			JoinTime:   clust.JoinTime.Format(time.RFC3339Nano),
			LastHealth: clust.LastHealthTime.Format(time.RFC3339Nano),
		}
		log.Info().Interface("send", mb).Send()
		if err := strm.Send(
			mb,
		); err != nil {
			return err
		}

	}
	return nil
}
func (h *adminService) Empty(context.Context, *connect.Request[common.Empty]) (*connect.Response[common.Empty], error) {
	return nil, nil
}

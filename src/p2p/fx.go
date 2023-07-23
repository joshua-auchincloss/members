package p2p

import (
	"context"
	"fmt"
	"log"
	"members/common"
	"members/config"
	"members/service"
	"time"

	"github.com/hashicorp/memberlist"
	"go.uber.org/fx"
)

var (
	Module = fx.Module(
		"p2p",
		fx.Provide(
			newRegistry,
		),
		fx.Invoke(
			newList,
			ensure_registry,
			// startNetwork,
		),
	)
)

func newList(prov config.ConfigProvider) error {
	mbrcfg := prov.ToMembership()
	mbrcfg.Events = &eventDelegate{}
	list, err := memberlist.Create(mbrcfg)
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}
	cfg := prov.GetConfig()

	known := []string{}
	if len(cfg.Members.Join) != 0 {
		known = append(known, cfg.Members.Join...)
		_, err = list.Join(known)
	}
	if err != nil {
		panic("Failed to join cluster: " + err.Error())
	}
	prov.EnsureList(list)
	return nil
}

func ensure_registry(reg *P2PRegistry) {
	log.Print(reg.store.Registered("abc"))
	project := common.ProtoProject{
		Name:  "abc-repo",
		Owner: "ja",
	}
	proto := common.ProtoMeta{
		Key:     "abc",
		Version: "v0.0.1",
	}
	reg.store.CreateProject(
		context.TODO(),
		&project,
		&proto,
	)
	b := `syntax = "proto3"`
	log.Print(
		reg.store.RegisterProto(
			context.TODO(),
			&proto,
			&common.RegisteredProto{
				FileName: "test.proto",
				Data:     []byte(b),
			},
		))
}

func startNetwork(prov config.ConfigProvider, ar *service.SvcFramework) {
	list := prov.GetList()
	for {
		log.Printf("num members: %d", list.NumMembers())
		for _, member := range list.Members() {
			fmt.Printf("Member: %s %s\n", member.Name, member.Addr)
			if member.State == memberlist.StateAlive {
				list.SendReliable(member, []byte("hello"))
			}
		}
		time.Sleep(time.Millisecond * 1000)
	}
}

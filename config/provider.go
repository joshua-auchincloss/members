package config

import (
	"members/common"
	"net"
	"strconv"
	"sync"

	"github.com/google/uuid"
	"github.com/hashicorp/memberlist"
	"github.com/urfave/cli/v2"
)

type (
	ConfigProvider interface {
		GetConfig() *Config
		ToMembership() *memberlist.Config
		EnsureList(li *memberlist.Memberlist) error
		GetList() *memberlist.Memberlist
		HostPort(port uint32) string
		EnsureFileCache(fn string)
		GetDynamic() *DynamicConfig
	}

	configProvider struct {
		lock *sync.Mutex

		ctx *cli.Context
		cfg *Config
		mls *memberlist.Config
		dyn *DynamicConfig
	}
)

var (
	_ ConfigProvider = ((*configProvider)(nil))
)

func New(ctx *cli.Context) (ConfigProvider, error) {
	cfg, err := getConfig(ctx)
	if err != nil {
		return nil, err
	}
	dyn, err := getDynamic()
	if err != nil {
		return nil, err
	}
	dyn.AddCluster(common.ServiceAdmin, "127.0.0.1:9010", "127.0.0.1:9010")
	dyn.AddCluster(common.ServiceAdmin, "127.0.0.1:8009", "127.0.0.1:8009")
	return &configProvider{
		lock: new(sync.Mutex),
		ctx:  ctx,
		cfg:  cfg,
		mls:  nil,
		dyn:  dyn,
	}, nil
}

func ensureFileCache(name string) func(prov ConfigProvider) {
	return func(prov ConfigProvider) {
		prov.EnsureFileCache(name)
	}
}

func (prov *configProvider) EnsureFileCache(fn string) {
	*prov.GetConfig().Storage = Storage{
		Kind:   "sqlite",
		URI:    fn,
		Drop:   false,
		Create: false,
	}
}

func (prov *configProvider) GetConfig() *Config {
	return prov.cfg
}

func (prov *configProvider) GetDynamic() *DynamicConfig {
	return prov.dyn
}

func (prov *configProvider) ToMembership() *memberlist.Config {
	if prov.mls == nil {
		memberlist.DefaultLANConfig()
		cf := memberlist.DefaultLocalConfig()
		cf.Name = uuid.NewString()
		cf.BindPort = int(prov.GetConfig().Members.Member)

		prov.lock.Lock()
		defer prov.lock.Unlock()
		prov.mls = cf
	}
	return prov.mls
}

func (prov *configProvider) HostPort(port uint32) string {
	return net.JoinHostPort(prov.GetConfig().Members.Bind, strconv.Itoa(int(port)))
}

func (prov *configProvider) EnsureList(li *memberlist.Memberlist) error {
	if prov.cfg.List != nil {
		return errExists
	}
	prov.lock.Lock()
	defer prov.lock.Unlock()
	prov.cfg.List = li
	return nil
}

func (prov *configProvider) GetList() *memberlist.Memberlist {
	return prov.cfg.List
}

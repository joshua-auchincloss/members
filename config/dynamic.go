package config

import (
	"members/common"
	"members/utils"
	"sync"

	"github.com/rs/zerolog/log"
	_ "github.com/spf13/viper/remote"
)

type (
	DynamicConfig struct {
		mu *sync.Mutex

		clusters map[common.Service]map[string][]string
	}
)

func (d *DynamicConfig) GetDns(dns string) []string {
	d.mu.Lock()
	defer d.mu.Unlock()
	known := []string{}
	for _, dns_clusters := range d.clusters {
		for tg, addrs := range dns_clusters {
			if tg == dns {
				known = append(known, addrs...)

			}
		}
	}
	return known
}

func (d *DynamicConfig) AllKnown() map[string][]string {
	d.mu.Lock()
	defer d.mu.Unlock()
	known := map[string][]string{}
	for _, dns_clusters := range d.clusters {
		for k, addrs := range dns_clusters {
			ar, ok := known[k]
			if !ok {
				ar = []string{}
			}
			for _, addr := range addrs {
				if !utils.AnyEq(
					ar, addr,
				) {
					ar = append(ar, addr)
				}
				known[k] = ar
			}
		}
	}
	return known
}
func (d *DynamicConfig) Sync(members []*common.Membership) {
	for _, memb := range members {
		// mu will sync this
		go d.AddCluster(memb.Service, memb.Dns, memb.PublicAddress)
	}
}

func (d *DynamicConfig) AddCluster(svc common.Service, dns, addr string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	m, ok := d.clusters[svc]
	if !ok {
		m = make(map[string][]string)
	}
	addrs, ok := m[dns]
	if !ok {
		addrs = []string{}
	}
	if !utils.AnyEq(addrs, addr) {
		addrs = append(addrs, addr)
	}
	m[dns] = addrs
	d.clusters[svc] = m

	log.Info().Str("service", common.ServiceKeys.Get(svc)).Str("dns", dns).Str("addr", addr).Msg("cluster updated")
}

func getDynamic() (*DynamicConfig, error) {
	dyn := &DynamicConfig{
		mu:       new(sync.Mutex),
		clusters: make(map[common.Service]map[string][]string),
	}
	return dyn, nil
}

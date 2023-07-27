package cluster

import (
	"members/common"
	"members/utils"
	"net"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

type (
	dnsList struct {
		mu      *sync.Mutex
		members map[string][]ClusterMember
	}

	DnsList interface {
		Peek() map[string][]ClusterMember
		AvailabilityWalk(wg *sync.WaitGroup, depth int)
		Walk(
			do func(addr ClusterMember) error,
			cap int,
			dns ...string) error
		Candidates(dns ...string) map[string][]string
		MergeWith(svc common.Service, dns string, other []ClusterMember)
		AddToCluster(svc common.Service, dns, proto, addr string)
	}
)

var (
	_ DnsList = ((*dnsList)(nil))
)

func NewDns(dns string, members ...ClusterMember) DnsList {
	start := map[string][]ClusterMember{}
	if len(members) != 0 {
		start[dns] = members
	}
	return &dnsList{
		new(sync.Mutex),
		start,
	}
}

func (d *dnsList) Peek() map[string][]ClusterMember {
	return d.members
}

func (d *dnsList) AvailabilityWalk(wg *sync.WaitGroup, depth int) {
	defer wg.Done()
	if err := d.Walk(
		func(addr ClusterMember) error {
			if !addr.Checked() || addr.Stale() {
				stat := addr.Suspend()
				stat.Check = time.Now()
				address := addr.Address()
				defer addr.Restore(stat)
				log.Info().Str("proto", "tcp").Str("address", address).Send()
				if conn, err := net.DialTimeout("tcp", address, time.Millisecond*1); err != nil {
					stat.Visible = false
					switch err {
					case net.ErrClosed:
					default:
						return err
					}
				} else {
					stat.Visible = true
					defer conn.Close()
				}
			}
			return nil
		},
		depth,
	); err != nil {
		log.Error().Err(err).Send()
	}
}

func (l *dnsList) Walk(
	do func(addr ClusterMember) error,
	cap int,
	dns ...string) error {
	var d string
	if len(dns) != 0 {
		d = dns[0]
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	check := !utils.ZeroStr(d)
	max := cap != -1 && cap > 1
	var seen int
	for dns, addrs := range l.members {
		var this bool
		if check {
			this = dns == d
		} else {
			this = true
		}
		if this {
			for _, addr := range addrs {
				if max {
					seen += 1
					if seen == cap {
						log.Warn().Int("cap", cap).Msg("cap reached")
						return nil
					}
					log.Debug().Int("seen", seen).Int("cap", cap).Send()
				}
				if err := do(addr); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (m *dnsList) Candidates(dns ...string) map[string][]string {
	return filter_list_for(
		m,
		func(addr ClusterMember) bool { return (!addr.Checked() || addr.Visible()) },
		func(addr ClusterMember) string { return addr.Address() },
		dns...,
	)
}

func (m *dnsList) MergeWith(svc common.Service, dns string, other []ClusterMember) {
	m.mu.Lock()
	defer m.mu.Unlock()
	var resolved []ClusterMember
	var ok bool
	if resolved, ok = m.members[dns]; !ok {
		resolved = other
	} else {
		resolved = append(resolved, other...)
	}
	m.members[dns] = resolved
}
func (m *dnsList) AddToCluster(svc common.Service, dns, proto, addr string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	c, ok := m.members[dns]
	if !ok {
		c = []ClusterMember{}
	}
	if !utils.OneOf(c, func(a ClusterMember) bool {
		return a.Address() == addr
	}) {
		c = append(c, &clusterMember{
			new(sync.Mutex),
			svc,
			proto,
			addr,
			nil,
		})
	}
	m.members[dns] = c
}

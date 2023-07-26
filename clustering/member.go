package cluster

import (
	"members/common"
	"sync"
)

type (
	clusterMember struct {
		mu *sync.Mutex

		role  common.Service
		proto string
		addr  string

		checked bool
		visible bool
	}

	ClusterMember interface {
		Visible() bool
		SetVisible(bool)
		Checked() bool
		SetChecked(bool)
		Protocol() string
		Address() string
		Lock()
		Unlock()
	}
)

var (
	_ ClusterMember = ((*clusterMember)(nil))
)

func (m *clusterMember) Address() string {
	return m.addr
}

func (m *clusterMember) Protocol() string {
	return m.proto
}

func (m *clusterMember) Role() common.Service {
	return m.role
}

func (m *clusterMember) Visible() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.visible
}

func (m *clusterMember) SetVisible(is bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.visible = is
}

func (m *clusterMember) Checked() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.checked
}

func (m *clusterMember) SetChecked(is bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.checked = is
}

func (m *clusterMember) Lock() {
	m.mu.Lock()
}
func (m *clusterMember) Unlock() {
	m.mu.Unlock()
}

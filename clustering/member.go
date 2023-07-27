package cluster

import (
	"members/common"
	"sync"
	"time"
)

var (
	stale time.Duration = time.Second * 45
)

type (
	MemberStatus struct {
		Check   time.Time
		Visible bool
	}

	clusterMember struct {
		mu *sync.Mutex

		role  common.Service
		proto string
		addr  string

		status *MemberStatus
	}

	ClusterMember interface {
		Visible() bool
		Checked() bool
		Stale() bool
		Status() *MemberStatus
		Suspend() *MemberStatus
		Restore(*MemberStatus)
		SetStatus(*MemberStatus)
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

func (m *clusterMember) Status() *MemberStatus {
	return m.status
}

func (m *clusterMember) Role() common.Service {
	return m.role
}
func (m *clusterMember) Stale() bool {
	switch {
	case m.status == nil:
		return true
	case time.Now().Add(-stale).After(m.status.Check):
		return true
	}
	return false
}

func (m *clusterMember) Visible() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.status != nil {
		return m.status.Visible
	}
	return false
}

func (m *clusterMember) Checked() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.status != nil
}

func (m *clusterMember) SetStatus(stat *MemberStatus) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.status = stat
}

func (m *clusterMember) Suspend() *MemberStatus {
	m.mu.Lock()
	if m.status == nil {
		return new(MemberStatus)
	}
	return m.status
}

func (m *clusterMember) Restore(stat *MemberStatus) {
	defer m.mu.Unlock()
	m.status = stat
}

func (m *clusterMember) Lock() {
	m.mu.Lock()
}
func (m *clusterMember) Unlock() {
	m.mu.Unlock()
}

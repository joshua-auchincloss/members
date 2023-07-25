package common

type (
	EventKind int
	Event     struct {
		Kind EventKind
	}
)

const (
	EventUnknown          EventKind = iota
	EventMembershipUpdate EventKind = iota + 101
	EventBulkMembershipUpdate
	EventMembershipAvailability
	EventBulkMembershipAvailability
)

var (
	EventKeys = NewKey(
		map[EventKind]string{
			EventMembershipUpdate:     "membership-update",
			EventBulkMembershipUpdate: "bulk-membership-update",
		},
	)
)

package dummy

import "github.com/agile-app/flexdb/internal/sla"

type dummy struct{}

// New creates a new SLA instance.
func Offer() sla.SLA {
	return &dummy{}
}

func (dummy) BorrowSearchQuota(orgID int64) bool {
	return true
}

func (dummy) ReturnSearchQuota(orgID int64) {}

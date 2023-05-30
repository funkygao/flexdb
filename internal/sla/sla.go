package sla

// SLA is the service level agreement of FlexDB.
type SLA interface {

	// BorrowSearchQuota borrows a quota ticket.
	// If quota exhausted, return false.
	BorrowSearchQuota(orgID int64) bool

	// ReturnSearchQuota returns a quota ticket.
	ReturnSearchQuota(orgID int64)
}

// Provider is the default SLA implementation.
var Provider SLA

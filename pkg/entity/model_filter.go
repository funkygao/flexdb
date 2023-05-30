package entity

import (
	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/plugins/filter"
)

// VisibleFor checks if the model is visible for an HTTP request context.
func (m *Model) VisibleFor(c context.RESTContext) bool {
	if m.Feature.ReadRowEnabled() {
		return true
	}

	// not readable for all, check perm
	return filter.SatisfyToBeKilledPermRule(c.PIN())
}

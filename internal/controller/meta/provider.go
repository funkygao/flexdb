package meta

import (
	"github.com/agile-app/flexdb/internal/controller"
)

var (
	mc *metaHandler // singleton
)

type metaHandler struct {
}

func init() {
	mc = &metaHandler{}
}

// Offer provides a singleton MetaController instance.
func Offer() controller.MetaHandler {
	return mc
}

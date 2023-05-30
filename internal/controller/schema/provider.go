package schema

import (
	"github.com/agile-app/flexdb/internal/controller"
)

var (
	sc *schemaHandler // singleton
)

type schemaHandler struct {
}

func init() {
	sc = &schemaHandler{}
}

// Offer provides a singleton MetaController instance.
func Offer() controller.SchemaHandler {
	return sc
}

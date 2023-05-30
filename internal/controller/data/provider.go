package data

import "github.com/agile-app/flexdb/internal/controller"

var (
	dc *dataHandler // singleton
)

type dataHandler struct {
}

func init() {
	dc = &dataHandler{}
}

// Offer provides a singleton DataController instance.
func Offer() controller.DataHandler {
	return dc
}

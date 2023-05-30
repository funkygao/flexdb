package entity

import (
	"encoding/json"

	"github.com/agile-app/flexdb/internal/spec"
)

var (
	_ json.Marshaler = (*App)(nil)
)

// MarshalJSON implements json.Marshaler so that App can have customized json output.
func (a *App) MarshalJSON() ([]byte, error) {
	type alias App // avoid infinite loop with alias type
	return json.Marshal(struct {
		*alias
		StatusLabel string `json:"statusLabel"`
		PrettyCTime string `json:"pctime"`
		PrettyMTime string `json:"pmtime"`
	}{
		alias:       (*alias)(a),
		StatusLabel: a.Status.Label(),
		PrettyCTime: a.CTime.Format(spec.YYYYMMDD),
		PrettyMTime: a.MTime.Format(spec.YYYYMMDD),
	})
}

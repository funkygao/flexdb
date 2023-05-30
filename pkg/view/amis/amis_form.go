package amis

import "github.com/agile-app/flexdb/internal/profile"

type Form struct {
	Type       string      `json:"type,omitempty"`
	Name       string      `json:"name,omitempty"`
	Autofocus  bool        `json:"autoFocus,omitempty"`
	InitApi    interface{} `json:"initApi,omitempty"`
	API        interface{} `json:"api,omitempty"`
	Reload     string      `json:"reload,omitempty"`
	Controls   []*Control  `json:"controls,omitempty"`
	SubmitText string      `json:"submitText,omitempty"`
	Target     string      `json:"target,omitempty"`
	Debug      bool        `json:"debug,omitempty"`
	Mode       string      `json:"mode,omitempty"` // horizontal
	Data       interface{} `json:"data,omitempty"`
}

func NewForm() *Form {
	return &Form{
		Type:     "form",
		Debug:    profile.LocalDebug(),
		Controls: make([]*Control, 0, 5),
	}
}

func (f *Form) AddControl(c *Control) {
	f.Controls = append(f.Controls, c)
}

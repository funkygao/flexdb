package amis

type Dialog struct {
	Type       string      `json:"type,omitempty"`
	Title      string      `json:"title,omitempty"`
	CloseOnEsc bool        `json:"closeOnEsc,omitempty"`
	Size       string      `json:"size,omitempty"`
	Body       interface{} `json:"body,omitempty"`
	Actions    []*Action   `json:"actions,omitempty"`
}

func NewDialog() *Dialog {
	return &Dialog{
		Type:    "dialog",
		Actions: make([]*Action, 0, 3),
	}
}

func (d *Dialog) AddAction(a *Action) {
	d.Actions = append(d.Actions, a)
}

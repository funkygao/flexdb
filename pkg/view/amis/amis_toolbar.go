package amis

type Toolbar struct {
	Type       string    `json:"type,omitempty"`
	ActionType string    `json:"actionType,omitempty"`
	Icon       string    `json:"icon,omitempty"`
	Label      string    `json:"label,omitempty"`
	Primary    bool      `json:"primary,omitempty"`
	Dialog     *Dialog   `json:"dialog,omitempty"`
	Link       string    `json:"link,omitempty"`
	Drawer     *Drawer   `json:"drawer,omitempty"`
	Tooltip    string    `json:"tooltip,omitempty"`
	Buttons    []*Button `json:"buttons,omitempty"`
}

func NewToolbar() *Toolbar {
	return &Toolbar{}
}

func (tb *Toolbar) AddButton(btn *Button) {
	tb.Buttons = append(tb.Buttons, btn)
}

package amis

type Tab struct {
	Title    string      `json:"title,omitempty"`
	Body     interface{} `json:"body,omitempty"`
	Controls []*Control  `json:"controls,omitempty"`
}

func NewTab() *Tab {
	return &Tab{}
}

func (t *Tab) AddControl(c *Control) {
	if t.Controls == nil {
		t.Controls = make([]*Control, 0, 3)
	}

	t.Controls = append(t.Controls, c)
}

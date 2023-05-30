package amis

type Filter struct {
	Title    interface{} `json:"title,omitempty"`
	Controls []*Control  `json:"controls,omitempty"`
	Actions  []*Control  `json:"actions,omitempty"`
}

func NewFilter() *Filter {
	return &Filter{
		Controls: make([]*Control, 0, 5),
		Actions:  make([]*Control, 0, 2),
	}
}

func (f *Filter) AddControl(c *Control) {
	f.Controls = append(f.Controls, c)
}

func (f *Filter) AddAction(c *Control) {
	f.Actions = append(f.Actions, c)
}

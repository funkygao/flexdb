package amis

type Column struct {
	Type       string      `json:"type,omitempty"`
	Name       string      `json:"name,omitempty"`
	Label      string      `json:"label,omitempty"`
	Sortable   bool        `json:"sortable,omitempty"`
	Copyable   interface{} `json:"copyable,omitempty"`
	QuickEdit  bool        `json:"quickEdit,omitempty"`
	Searchable bool        `json:"searchable,omitempty"`
	Body       interface{} `json:"body,omitempty"`
	Href       interface{} `json:"href,omitempty"`
	Blank      bool        `json:"blank,omitempty"`
	Fixed      string      `json:"fixed,omitempty"`
	Toggled    bool        `json:"toggled"`
	Remark     string      `json:"remark,omitempty"`
	PopOver    interface{} `json:"popOver,omitempty"`
	Buttons    []*Button   `json:"buttons,omitempty"`
}

func NewColumn() *Column {
	return &Column{
		Type:    "text",
		Toggled: true,
		Buttons: make([]*Button, 0, 5),
	}
}

func (c *Column) AddButton(b *Button) {
	c.Buttons = append(c.Buttons, b)
}

package amis

// Tpl is amis template.
type Tpl struct {
	Type    string `json:"type,omitempty"`
	Content string `json:"tpl,omitempty"`
	Inline  bool   `json:"inline"`
	CSS     string `json:"className,omitempty`
}

func NewTpl(c string) *Tpl {
	return &Tpl{
		Type:    "tpl",
		Content: c,
	}
}

package amis

type Button struct {
	Type        string      `json:"type,omitempty"`
	Label       string      `json:"label,omitempty"`
	Icon        string      `json:"icon,omitempty"`
	Tooltip     string      `json:"tooltip,omitempty"`
	Body        string      `json:"body,omitempty"`
	Level       string      `json:"level,omitempty"`
	Href        string      `json:"href,omitempty"`
	Content     string      `json:"content,omitempty"`
	ActionType  string      `json:"actionType,omitempty"`
	Size        string      `json:"size,omitempty"`
	VisibleOn   string      `json:"visibleOn,omitempty"`
	Primary     bool        `json:"primary,omitempty"`
	ClassName   string      `json:"className,omitempty"`
	Dialog      *Dialog     `json:"dialog,omitempty"`
	Drawer      *Drawer     `json:"drawer,omitempty"`
	API         interface{} `json:"api,omitempty"`
	ConfirmText string      `json:"confirmText,omitempty"`
}

func NewButton() *Button {
	return &Button{
		Type: "button",
	}
}

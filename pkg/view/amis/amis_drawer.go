package amis

type Drawer struct {
	Position   string      `json:"position,omitempty"`
	Title      string      `json:"title,omitempty"`
	Size       string      `json:"size,omitempty"`
	Resizable  bool        `json:"resizable,omitempty"`
	CloseOnEsc bool        `json:"closeOnEsc,omitempty"`
	Body       interface{} `json:"body,omitempty"`
}

func NewDrawer() *Drawer {
	return &Drawer{}
}

package amis

// Nav is amis nav node.
type Nav struct {
	Type    string `json:"type,omitempty"`
	Stacked bool   `json:"stacked,omitempty"`
	Links   []Link `json:"links,omitempty"`
}

// Link is the link inside a Nav.
type Link struct {
	Label    string `json:"label,omitempty"`
	To       string `json:"to,omitempty"`
	Icon     string `json:"icon,omitempty"`
	Blank    bool   `json:"blank,omitempty"` // _target=blank
	Active   bool   `json:"active,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
	Children []Link `json:"children,omitempty"`
}

func NewNav() *Nav {
	return &Nav{
		Type:  "nav",
		Links: make([]Link, 0, 5),
	}
}

func (n *Nav) AddLink(link Link) {
	n.Links = append(n.Links, link)
}

func NewLink(label, to string) *Link {
	return &Link{
		Label:    label,
		To:       to,
		Children: make([]Link, 0, 5),
	}
}

func (l *Link) AddChild(c Link) {
	l.Children = append(l.Children, c)
}

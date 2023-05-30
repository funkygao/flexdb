package amis

// Aside is amis aside node.
type Aside struct {
	Type string        `json:"type,omitempty"`
	Size string        `json:"size,omitempty"`
	Body []interface{} `json:"body,omitempty"`
}

func (a *Aside) AddBody(b interface{}) {
	a.Body = append(a.Body, b)
}

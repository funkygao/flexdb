package amis

type Action struct {
	Type       string `json:"type,omitempty"`
	ActionType string `json:"actionType,omitempty"`
	Label      string `json:"label,omitempty"`
	Primary    bool   `json:"primary,omitempty"`
}

func NewAction() *Action {
	return &Action{}
}

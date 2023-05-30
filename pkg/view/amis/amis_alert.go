package amis

type Alert struct {
	Type  string `json:"type,omitempty"`
	Body  string `json:"body,omitempty"`
	Level string `json:"level,omitempty"`
}

func NewAlert(body string) *Alert {
	return &Alert{
		Type: "alert",
		Body: body,
	}
}

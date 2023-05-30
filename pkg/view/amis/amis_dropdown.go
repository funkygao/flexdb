package amis

type Dropdown struct {
	Type      string   `json:"type,omitempty"`
	Label     string   `json:"label,omitempty"`
	ClassName string   `json:"className,omitempty"`
	Size      string   `json:"size,omitempty"`
	Buttons   []Button `json:"buttons,omitempty"`
}

func NewDropdown() *Dropdown {
	return &Dropdown{
		Type: "dropdown-button",
	}
}

func (d *Dropdown) AddButton(b Button) {
	d.Buttons = append(d.Buttons, b)
}

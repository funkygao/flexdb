package amis

type Wizard struct {
	Type string      `json:"type,omitempty"`
	Mode string      `json:"mode,omitempty"` // horizontal | vertical
	API  interface{} `json:"api,omitempty"`  // last step save api

	Steps []*Step `json:"steps,omitempty"`
}

func NewWizard() *Wizard {
	return &Wizard{
		Type:  "wizard",
		Steps: make([]*Step, 0, 5),
	}
}

func (w *Wizard) AddStep(step *Step) {
	w.Steps = append(w.Steps, step)
}

type Step struct {
	Title    string      `json:"title,omitempty"`
	API      interface{} `json:"api,omitempty"`  // current step save api
	Mode     string      `json:"mode,omitempty"` // horizontal | vertical
	Controls []*Control  `json:"controls,omitempty"`
}

func NewStep(title string) *Step {
	return &Step{Title: title}
}

func (s *Step) AddControl(c *Control) {
	if s.Controls == nil {
		s.Controls = make([]*Control, 5)
	}

	s.Controls = append(s.Controls, c)
}

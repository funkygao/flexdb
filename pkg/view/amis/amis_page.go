package amis

// Page is amis page node.
// see https://houtai.baidu.com/v2/schemas/page.json
type Page struct {
	Type     string        `json:"type,omitempty"`
	Title    interface{}   `json:"title,omitempty"`
	SubTitle interface{}   `json:"subTitle,omitempty"`
	Remark   interface{}   `json:"remark,omitempty"`
	InitAPI  string        `json:"initApi,omitempty"`
	Aside    *Aside        `json:"aside,omitempty"`
	Body     []interface{} `json:"body,omitempty"`
	Toolbar  []*Toolbar    `json:"toolbar,omitempty"`
}

func NewPage() *Page {
	return &Page{
		Type:    "page",
		Body:    make([]interface{}, 0, 10),
		Toolbar: make([]*Toolbar, 0, 3),
	}
}

func (p *Page) SetAside(typ string) *Aside {
	p.Aside = &Aside{
		Type: typ,
		Body: make([]interface{}, 0, 5),
	}

	return p.Aside
}

func (p *Page) AddBody(b interface{}) {
	p.Body = append(p.Body, b)
}

func (p *Page) AddToolbar(t *Toolbar) {
	p.Toolbar = append(p.Toolbar, t)
}

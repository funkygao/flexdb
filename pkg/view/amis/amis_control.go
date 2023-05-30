package amis

import "github.com/agile-app/flexdb/pkg/dto"

// Control is amis SchemaNode.
type Control struct {
	Type         string      `json:"type,omitempty"`
	Name         string      `json:"name,omitempty"`
	Label        string      `json:"label,omitempty"`
	Value        interface{} `json:"value,omitempty"`
	Placeholder  string      `json:"placeholder,omitempty"`
	API          interface{} `json:"api,omitempty"`
	Disabled     bool        `json:"disabled,omitempty"`
	DisabledOn   string      `json:"disabledOn,omitempty"`
	Tpl          string      `json:"tpl,omitempty"`
	Title        interface{} `json:"title,omitempty"`
	URL          string      `json:"url,omitempty"`
	Body         interface{} `json:"body,omitempty"`
	MultiLine    bool        `json:"multiLine,omitempty"`
	Icon         string      `json:"icon,omitempty"`
	ClassName    string      `json:"className,omitempty"`
	ConfirmText  string      `json:"confirmText,omitempty"`
	Validations  string      `json:"validations,omitempty"`
	Header       interface{} `json:"header,omitempty"` // card
	Required     bool        `json:"required,omitempty"`
	VisibleOn    string      `json:"visibleOn,omitempty"`
	Dialog       interface{} `json:"dialog,omitempty"`
	Remark       string      `json:"remark,omitempty"`
	Desc         string      `json:"desc,omitempty"`    // 与Remark类似，但直接显示出来
	MaxDate      string      `json:"maxDate,omitempty"` // e,g. "maxDate": "${endtime}"
	MinDate      string      `json:"minDate,omitempty"` // e,g. "minDate": "${starttime}"
	ActionType   string      `json:"actionType,omitempty"`
	Level        string      `json:"level,omitempty"`
	Hint         string      `json:"hint,omitempty"`
	Source       string      `json:"source,omitempty"`
	ListItem     interface{} `json:"listItem,omitempty"`
	ColumnsCount int         `json:"columnsCount,omitempty"` // cards
	Inline       bool        `json:"inline,omitempty"`
	Link         string      `json:"link,omitempty"`
	Sortable     bool        `json:"sortable,omitempty"`
	Searchable   bool        `json:"searchable,omitempty"`
	SearchAPI    interface{} `json:"searchApi,omitempty"`
	Size         string      `json:"size,omitempty"`
	Creatable    bool        `json:"creatable,omitempty"`
	AddAPI       interface{} `json:"addApi,omitempty"`
	Accept       string      `json:"accept,omitempty"`   // file
	AsBlob       bool        `json:"asBlob,omitempty"`   // file
	Reciever     interface{} `json:"reciever,omitempty"` // file
	Content      string      `json:"content,omitempty"`  // remark
	Controls     []*Control  `json:"controls,omitempty"`
	AddControls  interface{} `json:"addControls,omitempty"`
	Columns      []*Control  `json:"columns,omitempty"`
	Actions      []*Control  `json:"actions,omitempty"`
	Tabs         []*Tab      `json:"tabs,omitempty"`
	Editable     bool        `json:"editable,omitempty"`
	EditAPI      interface{} `json:"editApi,omitempty"`
	Options      interface{} `json:"options,omitempty"`
	AutoComplete interface{} `json:"autoComplete,omitempty"` // e,g. https://houtai.baidu.com/api/mock2/options/autoComplete?term=$term
	AddOn        dto.H       `json:"addOn,omitempty"`
	SelectMode   string      `json:"selectMode,omitempty"`
}

func NewControl() *Control {
	return &Control{
		Type: "text",
	}
}

func (c *Control) AddControl(ctrl *Control) {
	if c.Controls == nil {
		c.Controls = make([]*Control, 0, 5)
	}
	c.Controls = append(c.Controls, ctrl)
}

func (c *Control) AddTab(tab *Tab) {
	if c.Tabs == nil {
		c.Tabs = make([]*Tab, 0, 1)
	}
	c.Tabs = append(c.Tabs, tab)
}

func (c *Control) AddColumn(col *Control) {
	if c.Columns == nil {
		c.Columns = make([]*Control, 0, 5)
	}
	c.Columns = append(c.Columns, col)
}

func (c *Control) AddAction(col *Control) {
	if c.Actions == nil {
		c.Actions = make([]*Control, 0, 5)
	}
	c.Actions = append(c.Actions, col)
}

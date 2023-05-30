package amis

type Crud struct {
	Type                 string      `json:"type,omitempty"`
	Draggable            bool        `json:"draggable,omitempty"`
	Mode                 string      `json:"mode,omitempty"` // table, list, card
	FilterTogglable      bool        `json:"filterTogglable,omitempty"`
	HeaderToolbar        []string    `json:"headerToolbar,omitempty"`
	FooterToolbar        []string    `json:"footerToolbar,omitempty"`
	API                  interface{} `json:"api,omitempty"`
	QuickSaveAPI         interface{} `json:"quickSaveApi,omitempty"`
	QuickSaveItemAPI     interface{} `json:"quickSaveItemApi,omitempty"`
	SaveOrderAPI         interface{} `json:"saveOrderApi,omitempty"`
	Filter               *Filter     `json:"filter,omitempty"`
	Card                 interface{} `json:"card,omitempty"`
	ColumnsCount         int         `json:"columnsCount,omitempty"` // cards
	Data                 interface{} `json:"data,omitempty"`
	Columns              []*Column   `json:"columns,omitempty"`
	DefaultParams        interface{} `json:"defaultParams,omitempty"`
	PerPageAvailable     []int       `json:"perPageAvailable,omitempty"`
	FilterDefaultVisible bool        `json:"filterDefaultVisible"`
	APIIntervalInMs      int         `json:"interval,omitempty"`
	BulkActions          []*Control  `json:"bulkActions,omitempty"`
}

func NewCrud() *Crud {
	return &Crud{
		Type:                 "crud",
		Columns:              make([]*Column, 0, 10),
		BulkActions:          make([]*Control, 0, 5),
		Draggable:            true,
		FilterTogglable:      true,
		FilterDefaultVisible: false,
		PerPageAvailable:     []int{10, 20, 50, 100, 2000}, // TODO quota
		HeaderToolbar:        []string{"bulk-actions", "filter-toggler", "export-excel", "pagination"},
		FooterToolbar:        []string{"statistics", "switch-per-page", "pagination", "load-more"},
	}
}

func (c *Crud) Reset() {
	c.Columns = nil
	c.BulkActions = nil
	c.Draggable = false
	c.FilterTogglable = false
	c.FilterDefaultVisible = false
	c.PerPageAvailable = nil
	c.HeaderToolbar = nil
	c.FooterToolbar = nil
	c.Mode = "table"
}

func (c *Crud) AddColumn(col *Column) {
	c.Columns = append(c.Columns, col)
}

func (c *Crud) AddBulkAction(ctrl *Control) {
	c.BulkActions = append(c.BulkActions, ctrl)
}

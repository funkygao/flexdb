package view

import (
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/view/amis"
)

// ColumnWidget is in charge of UI for the column.
type ColumnWidget struct {
	*entity.Column
}

// ColumnWidgetOf creates a column widget reference of column.
func ColumnWidgetOf(c *entity.Column) ColumnWidget {
	return ColumnWidget{Column: c}
}

// RenderLabel returns label in CRUD page.
func (c *ColumnWidget) RenderLabel() string {
	if c.Label != "" {
		return c.Label
	}

	return c.Name
}

// ListWidgetType returns column type in CRUD page list view.
func (c *ColumnWidget) ListWidgetType() string {
	if p, yes := c.Plugin().(entity.ListViewWidgetTypeAware); yes {
		return p.ListViewWidgetType()
	}

	return "text"
}

func (c *ColumnWidget) ListViewHrefIfNec() (href interface{}, anchor interface{}) {
	if p, yes := c.Plugin().(entity.ListViewHrefer); yes {
		href, anchor = p.ListViewHref()
		return
	}

	return
}

// EditWidgetType returns the UI widget type of this column in Form.
func (c *ColumnWidget) EditWidgetType() string {
	return editWidgetTypes[c.Kind]
}

// EnrichEditWidget is part of EditWidgetAware.
func (c *ColumnWidget) EnrichEditWidget(ctrl *amis.Control) {
	if p, yes := c.Plugin().(entity.EditWidgetAware); yes {
		p.EnrichEditControl(ctrl)
	}
}

func (c *ColumnWidget) EnrichViewWidget(col *amis.Column) {
	if p, yes := c.Plugin().(entity.ViewWidgetAware); yes {
		p.EnrichViewControl(col)
	}
}

package entity

import "github.com/agile-app/flexdb/pkg/view/amis"

// ListViewExtraCeller renders an extra column data in list view.
// e,g. datetime stored as timestamp, and renders with another {col}_h for human readable format.
type ListViewExtraCeller interface {
	ListViewExtraCell(currentCellVal string) (extraCellColumn, extraCellVal string)
}

// ListViewWidgetTypeAware renders the column cell value as specified type in CRUD list page.
type ListViewWidgetTypeAware interface {
	ListViewWidgetType() string
}

// ListViewHrefer renders href cell value in CRUD list page.
type ListViewHrefer interface {
	ListViewHref() (href interface{}, anchor interface{})
}

// EditWidgetAware is an optional column plugin that renders the widget.
type EditWidgetAware interface {

	// EnrichEditControl enriches the widget extra UI schema basides standard control.
	EnrichEditControl(*amis.Control)
}

// ViewWidgetAware renders extra schema in view page.
type ViewWidgetAware interface {
	EnrichViewControl(*amis.Column)
}

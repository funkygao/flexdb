package controller

import "github.com/agile-app/flexdb/internal/controller/context"

// MetaHandler defines meta related RESTful APIs.
type MetaHandler interface {
	CreateApp(c context.RESTContext)
	UpdateApp(c context.RESTContext)
	FindApps(c context.RESTContext)
	ShowApp(c context.RESTContext)

	UploadModel(c context.RESTContext)
	CreateModel(c context.RESTContext)
	UpdateModel(c context.RESTContext)
	ListModels(c context.RESTContext)
	ShowModel(c context.RESTContext)

	ReorderColumns(c context.RESTContext)
	AddColumn(c context.RESTContext)
	DeprecateColumn(c context.RESTContext)
	UpdateColumn(c context.RESTContext)

	ShowPages(c context.RESTContext)
	ShowPage(c context.RESTContext)

	FindTemplates(c context.RESTContext)
}

// SchemaHandler defines amis page schema related RESTful APIs.
type SchemaHandler interface {
	FindApps(c context.RESTContext)
	ShowApp(c context.RESTContext)
	ModelCRUD(c context.RESTContext)
	ModelSchemaCRUD(c context.RESTContext)
}

// DataHandler defines data related RESTful APIs.
type DataHandler interface {
	ImportRows(c context.RESTContext)

	// CreateRow creates a new row for a specified model.
	CreateRow(c context.RESTContext)

	// RetrieveRow retrieves a single row from a specified model and interprets
	// into user defined model schema.
	RetrieveRow(c context.RESTContext)

	UpdateRow(c context.RESTContext)

	QuickSave(c context.RESTContext)

	DeleteRow(c context.RESTContext)

	// FindRows search rows in pagination.
	FindRows(c context.RESTContext)

	// Lookup used for lookup column.
	Lookup(c context.RESTContext)

	CreatePickItem(c context.RESTContext)
	UpdatePickItem(c context.RESTContext)
	ShowPicklist(c context.RESTContext)
}

type UserHandler interface {
	GetUserInfo(c context.RESTContext)

	RecommendAppUsers(c context.RESTContext)

	ShareAppToUser(c context.RESTContext)
}

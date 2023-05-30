package entity

import "errors"

var (
	ErrAppEmptyOrgID = errors.New("empty app.orgID")
	ErrAppEmptyName  = errors.New("empty app.name")
	ErrAppEmptyID    = errors.New("empty app.ID")

	ErrModelEmptyAppID = errors.New("empty model.appID")
	ErrModelEmptyName  = errors.New("empty model.name")
	ErrModelEmptyKind  = errors.New("empty model.kind")

	ErrColumnEmptyID      = errors.New("empty column.id not allowed")
	ErrColumnEmptyModelID = errors.New("empty column.modelID")
	ErrColumnUnknownKind  = errors.New("unknown column.kind")
	ErrColumnEmptySlot    = errors.New("empty column.slot")
	ErrColumnEmptyName    = errors.New("empty column.name")

	ErrIndexEmptyModelID = errors.New("empty index.modelID")

	ErrRowEmptyID      = errors.New("empty row.id")
	ErrRowEmptyOrgID   = errors.New("empty row.orgID")
	ErrRowEmptyModelID = errors.New("empty row.modelID")

	ErrClobEmptyOrgID = errors.New("empty clob.orgID")
	ErrClobEmptyRowID = errors.New("empty clob.rowID")
)

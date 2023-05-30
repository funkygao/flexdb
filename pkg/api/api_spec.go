// Package api defines spec shared between FlexDB and client.
package api

const (
	ResponseCodeOK  = 0
	ResponseCodeErr = 10

	FlexClientHeader = "X-FlexClient-Ver"

	OrgIDHeader    = "X-Flex-Org-Id"
	OperatorHeader = "X-Flex-Operator"
	RevisionHeader = "X-FlexDB-Revision"

	RowID = "_rowid"

	QueryOrderBy        = "orderBy"  // e,g. mtime
	QueryOrderDirection = "orderDir" // e,g. asc, desc

	Debugger = "_debugger_"
)

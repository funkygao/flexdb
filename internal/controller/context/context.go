package context

import (
	"github.com/agile-app/flexdb/pkg/dto"
	"github.com/gin-gonic/gin"
)

// RESTContext is the RESTful API context that can extract HTTP request params.
type RESTContext interface {
	Gin() *gin.Context

	// AMIS tells whether the request originated from AMIS client.
	// Besides AMIS, it can be flexsdk: FlexDB go client.
	AMIS() bool

	// QueryID returns the id from the url.
	// e,g. http://foo.bar/abc?id=5 QueryID() will return 5
	QueryID() int64

	// T translates a msg by key with HTTP request locale.
	T(key string, args ...string) string

	OrgID() int64
	ModelID() int64
	AppID() int64
	RowID() uint64
	SlotID() int16
	ColumnID() int64
	UUID() string

	SessionID() uint64

	PageIndex() int
	PageSize() int

	// PIN returns current login user Personal Identification Number.
	PIN() string

	SearchCriteria() dto.Criteria

	RenderOK(data interface{})
	RenderOKWithoutData()
	RenderWithMsg(msg string)
	AbortWithError(err error)
}

// Offer provides the default RESTContext instance.
var Offer func(*gin.Context) RESTContext

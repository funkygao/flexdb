package revision

import (
	"github.com/agile-app/flexdb/pkg/api"
	"github.com/funkygao/golib/version"
	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set(api.RevisionHeader, version.Revision)

		c.Next()
	}
}

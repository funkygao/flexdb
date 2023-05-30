package casbin

import (
	"errors"
	"net/http"

	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

var ErrNoPerm = errors.New("no perm")

func Middleware(ctx context.RESTContext, enforcer *casbin.SyncedEnforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		p := c.Request.URL.Path
		m := c.Request.Method
		if ok, err := enforcer.Enforce(ctx.PIN(), p, m); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		} else if !ok {
			c.AbortWithError(http.StatusForbidden, ErrNoPerm)
			return
		}

		c.Next()
	}
}

package perm

import (
	"fmt"

	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	"github.com/gin-gonic/gin"
)

func Authorize(ctx context.RESTContext, obj string, act string, adapter persist.Adapter) gin.HandlerFunc {
	return func(c *gin.Context) {
		// casbin enforces policy
		ok, err := enforce(ctx.PIN(), obj, act, adapter)
		if err != nil {
			ctx.AbortWithError(nil)
			return
		}
		if !ok {
			ctx.AbortWithError(nil)
			return
		}

		c.Next()
	}
}

func enforce(sub, obj, act string, adapter persist.Adapter) (bool, error) {
	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		return false, fmt.Errorf("failed to create casbin enforcer: %w", err)
	}

	// Load policies from DB dynamically
	err = enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}

	enforcer.AddPolicy("bob", "/app", "POST")

	return enforcer.Enforce(sub, obj, act)
}

package router

import (
	"net/http"
	"time"

	"github.com/agile-app/flexdb/internal/profile"
	"github.com/agile-app/flexdb/internal/router/middleware/authz"
	"github.com/agile-app/flexdb/internal/router/middleware/jwt"
	"github.com/agile-app/flexdb/internal/router/middleware/locale"
	"github.com/agile-app/flexdb/internal/router/middleware/revision"
	"github.com/agile-app/flexdb/internal/router/middleware/session"
	"github.com/agile-app/flexdb/pkg/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine // singleton
)

func prepare() {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	router = gin.New()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// Insert cors middleware definition before any routes
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = profile.P.CORSAllowOrigins
	corsConfig.AllowWildcard = true // TODO
	corsConfig.AllowHeaders = []string{api.OrgIDHeader, api.OperatorHeader, "X-Action-Addr", "Content-Type"}
	corsConfig.AllowCredentials = true // allows cross-origin cookie
	corsConfig.MaxAge = time.Hour * 24 // preflight request cache ttl

	// middleware chain
	router.Use(gin.Recovery())
	router.Use(cors.New(corsConfig))
	router.Use(session.Middleware())
	setupJWT(router)
	router.Use(authz.Middleware())
	router.Use(locale.Middleware())
	if !profile.Debug() {
		router.Use(gzip.Gzip(gzip.DefaultCompression))
	}
	router.Use(revision.Middleware())
}

func setupJWT(router *gin.Engine) {
	middleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
	})
	if err != nil {
		panic(err)
	}

	if false {
		router.Use(middleware.MiddlewareFunc())
	}
}

// Offer provides a gin engine.
func Offer() *gin.Engine {
	if router == nil {
		prepare()
	}

	return router
}

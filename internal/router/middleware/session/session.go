package session

import (
	"errors"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/agile-app/flexdb/internal/profile"
	"github.com/agile-app/flexdb/pkg/api"
	"github.com/agile-app/flexdb/pkg/connector"
	"github.com/funkygao/go-metrics"
	"github.com/funkygao/log4go"
	"github.com/gin-gonic/gin"
)

const (
	keyCurrentUser = "_user"
	keyOrgID       = "_org"
	keySessionID   = "_sid"
)

var (
	errInvalidOrgID = errors.New("invalid orgID")
	reqID           int64

	responseSize = metrics.NewRegisteredHistogram("http.reply.sz", nil, metrics.NewExpDecaySample(1028, 0.015))
	latency      = metrics.NewRegisteredHistogram("http.reply.ms", nil, metrics.NewExpDecaySample(1028, 0.015))
)

// Middleware auto set current user pin in the context of the current request.
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t0 := time.Now()

		orgID, err := strconv.ParseInt(c.GetHeader(api.OrgIDHeader), 10, 64)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, errInvalidOrgID)
			return
		}
		// remember org id
		c.Set(keyOrgID, orgID)

		//rid := uuid.New().String() github.com/google/uuid
		rid := atomic.AddInt64(&reqID, 1)
		c.Set(keySessionID, rid)

		var pin string
		if pin = c.Request.Header.Get(api.OperatorHeader); pin == "" { // FIXME will be cheated
			pin, err = connector.SSOConnector(connector.OrgID(orgID)).CurrentUser(c)
			if err != nil {
				if profile.LocalDebug() {
					c.Next() // currently let pass, FIXME block anonymous access

					cost := time.Since(t0)
					latency.Update(cost.Milliseconds())
					responseSize.Update(int64(c.Writer.Size()))
					log4go.Warn("#%d$ %s [%d] sz:%d %s %s %s", rid,
						cost, c.Writer.Status(), c.Writer.Size(),
						c.Request.RemoteAddr, c.Request.Method, c.Request.RequestURI)
					return
				} else {
					log4go.Warn("#%d %s %s %s: %v", rid,
						c.Request.RemoteAddr, c.Request.Method, c.Request.RequestURI, err)
					c.AbortWithError(http.StatusUnauthorized, err)
					return
				}
			}
		}
		// remember pin
		c.Set(keyCurrentUser, pin)

		log4go.Info("#%d %s pin:%s %s %s", rid, c.Request.RemoteAddr,
			pin, c.Request.Method, c.Request.RequestURI)
		c.Next()

		cost := time.Since(t0)
		latency.Update(cost.Milliseconds())
		responseSize.Update(int64(c.Writer.Size()))
		log4go.Info("#%d$ %s [%d] sz:%d %s pin:%s %s %s", rid,
			cost, c.Writer.Status(), c.Writer.Size(),
			c.Request.RemoteAddr, pin, c.Request.Method, c.Request.RequestURI)
	}
}

// CurrentUser returns the current login user of the session.
func CurrentUser(c *gin.Context) string {
	return c.GetString(keyCurrentUser)
}

// OrgID returns org id of the session.
func OrgID(c *gin.Context) int64 {
	return c.GetInt64(keyOrgID)
}

// ID returns id of the request session.
func ID(c *gin.Context) int64 {
	return c.GetInt64(keySessionID)
}

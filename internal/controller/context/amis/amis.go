package amis

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/internal/i18n"
	"github.com/agile-app/flexdb/internal/profile"
	"github.com/agile-app/flexdb/internal/router/middleware/session"
	"github.com/agile-app/flexdb/pkg/api"
	"github.com/agile-app/flexdb/pkg/dto"
	"github.com/funkygao/log4go"
	"github.com/gin-gonic/gin"
)

type amis struct {
	*gin.Context
}

func Of(c *gin.Context) context.RESTContext {
	return &amis{Context: c}
}

func (c *amis) AMIS() bool {
	if c.Gin().Request.Header.Get(api.FlexClientHeader) != "" {
		return false
	}

	return true
}

func (c *amis) T(key string, args ...string) string {
	return i18n.B.TranslationsForRequest(c.Context.Request).Value(key, args...)
}

func (c *amis) SessionID() uint64 {
	return uint64(session.ID(c.Gin()))
}

func (c *amis) OrgID() int64 {
	return session.OrgID(c.Gin())
}

func (c *amis) Gin() *gin.Context {
	return c.Context
}

func (c *amis) PIN() string {
	if profile.LocalDebug() {
		return api.Debugger
	}

	return session.CurrentUser(c.Context)
}

func (c *amis) ModelID() int64 {
	id, err := strconv.ParseInt(c.Param("model"), 10, 64)
	if err != nil {
		return 0
	}
	return id
}

func (c *amis) AppID() int64 {
	id, err := strconv.ParseInt(c.Param("app"), 10, 64)
	if err != nil {
		return 0
	}

	return id
}

func (c *amis) ColumnID() int64 {
	id, err := strconv.ParseInt(c.Param("column"), 10, 64)
	if err != nil {
		return 0
	}
	return id
}

func (c *amis) RowID() uint64 {
	id, err := strconv.ParseUint(c.Param("row"), 10, 64)
	if err != nil {
		return 0
	}

	return id
}

func (c *amis) SlotID() int16 {
	id, err := strconv.ParseInt(c.Param("slot"), 10, 64)
	if err != nil {
		return 0
	}

	return int16(id)
}

func (c *amis) UUID() string {
	if uuid, present := c.GetQuery("uuid"); present {
		return uuid
	}

	return "_invalid_"
}

func (c *amis) QueryID() int64 {
	if p, present := c.GetQuery("id"); present {
		id, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
			return 0
		}

		return id
	}

	return 0
}

func (c *amis) PageIndex() int {
	if p, present := c.GetQuery("page"); present {
		page, err := strconv.Atoi(p)
		if err != nil {
			return 1
		}

		if profile.P.PaginationMaxRowsScan > 0 &&
			page*c.PageSize() > profile.P.PaginationMaxRowsScan {
			// deep pagination disallowed
			return 1
		}

		return page
	}

	return 1
}

func (c *amis) PageSize() int {
	if p, present := c.GetQuery("page"); present {
		page, err := strconv.Atoi(p)
		if err != nil {
			return 10
		}

		if p, present := c.GetQuery("perPage"); present {
			pageSize, err := strconv.Atoi(p)
			if err != nil {
				return 10
			}

			if profile.P.PaginationMaxRowsScan > 0 &&
				pageSize*page > profile.P.PaginationMaxRowsScan {
				// deep pagination disallowed
				return 1
			}

			return pageSize
		}
	}

	return 10
}

func (c *amis) SearchCriteria() (criteria dto.Criteria) {
	g := c.Gin()
	criteria = make(dto.Criteria, 0, 5)
	for key, vals := range g.Request.URL.Query() {
		if len(vals) == 0 || vals[0] == "" {
			continue
		}
		if _, excluded := excludedCriteria[key]; excluded {
			continue
		}

		item := dto.CriteriaItem{
			Key: key,
			Op:  "=",
			Val: strings.TrimSpace(vals[0]),
		}
		tuple := strings.Split(item.Val, " ")
		if _, present := criteriaOperators[tuple[0]]; present {
			item.Op = tuple[0]
			item.Val = strings.Join(tuple[1:], "")
			if item.Op == "like" {
				item.Val = item.Val + "%" // only prefix matching supported
			}
		}

		criteria = criteria.Append(item)
	}

	return criteria
}

func (c *amis) RenderOK(data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": api.ResponseCodeOK,
		"msg":    "OK",
		"data":   data,
	})
}

func (c *amis) RenderOKWithoutData() {
	c.JSON(http.StatusOK, gin.H{
		"status": api.ResponseCodeOK,
		"msg":    "OK",
	})
}

func (c *amis) RenderWithMsg(msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status": api.ResponseCodeOK,
		"msg":    msg,
	})
}

func (c *amis) AbortWithError(err error) {
	log4go.Error("#%d pin:%s %s %s %s %v",
		c.SessionID(), c.PIN(),
		c.Request.RemoteAddr, c.Request.Method, c.Request.RequestURI,
		err)

	c.Gin().AbortWithStatusJSON(http.StatusOK, gin.H{
		"status": api.ResponseCodeErr,
		"msg":    err.Error(),
	})
}

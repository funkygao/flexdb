package user

import (
	"strings"
	"time"

	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/pkg/connector"
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/store"
	"github.com/gin-gonic/gin"
)

func (uc *userHandler) GetUserInfo(c context.RESTContext) {
	c.RenderOK(gin.H{
		"nickname": c.PIN(),
		"avatar":   "/static/images/default_avatar.jpg",
	})
}

func (uc *userHandler) RecommendAppUsers(c context.RESTContext) {
	term := c.Gin().Query("term")
	r, err := connector.SSOConnector(connector.OrgID(c.OrgID())).SearchUser(term)
	if err != nil {
		c.AbortWithError(err)
		return
	}

	c.RenderOK(gin.H{
		"options": r,
	})
}

func (uc *userHandler) ShareAppToUser(c context.RESTContext) {
	var form struct {
		ShareTo string `json:"shareTo"`
	}
	if err := c.Gin().ShouldBindJSON(&form); err != nil {
		c.AbortWithError(err)
		return
	}

	// TODO perm validation
	os := store.Provider.InferOrgStore(c)
	for _, s := range strings.Split(form.ShareTo, ",") {
		share := &entity.Share{
			AppID:   c.AppID(),
			Subject: s,
			CTime:   time.Now(),
			CUser:   c.PIN(),
		}

		if err := os.ShareApp(nil, share); err != nil {
			c.AbortWithError(err)
			return
		}
	}

	c.RenderOKWithoutData()
}

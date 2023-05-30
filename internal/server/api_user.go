package server

import (
	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/gin-gonic/gin"
)

func (s *Server) exportUserAPIs() {
	userGroup := s.router.Group("/api/v0.1/user")

	userGroup.GET("/info", func(c *gin.Context) {
		s.uc.GetUserInfo(context.Offer(c))
	})
	userGroup.GET("/recommend/:app", func(c *gin.Context) {
		s.uc.RecommendAppUsers(context.Offer(c))
	})
	userGroup.POST("/share/:app", func(c *gin.Context) {
		s.uc.ShareAppToUser(context.Offer(c))
	})
}

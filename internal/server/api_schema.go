package server

import (
	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/gin-gonic/gin"
)

func (s *Server) exportSchemaAPIs() {
	schemaGroup := s.router.Group("/api/v0.1/schema")

	schemaGroup.GET("/App", s.appsHandler)
	schemaGroup.GET("/App/:app", s.appsHandler)
	schemaGroup.GET("/App/:app/Schema/:model", func(c *gin.Context) {
		s.sc.ModelSchemaCRUD(context.Offer(c))
	})
	schemaGroup.GET("/App/:app/Model/:model", func(c *gin.Context) {
		s.sc.ModelCRUD(context.Offer(c))
	})
	schemaGroup.GET("/App/:app/Launch", func(c *gin.Context) {
		s.sc.ModelCRUD(context.Offer(c))
	})
}

func (s *Server) appsHandler(c *gin.Context) {
	appIDStr := c.Param("app")
	if appIDStr != "" {
		s.sc.ShowApp(context.Offer(c))
		return
	}

	s.sc.FindApps(context.Offer(c))
}

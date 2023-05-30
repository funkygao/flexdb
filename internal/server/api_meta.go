package server

import (
	"net/http"
	"strings"

	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/gin-gonic/gin"
)

func (s *Server) exportMetaAPIs() {
	metaGroup := s.router.Group("/api/v0.1/meta")

	metaGroup.GET("/Template", func(c *gin.Context) {
		s.mc.FindTemplates(context.Offer(c))
	})

	appGroup := metaGroup.Group("/App")
	appGroup.GET("/", func(c *gin.Context) {
		s.mc.FindApps(context.Offer(c))
	})
	appGroup.POST("/", s.appHandler)
	appGroup.GET("/:app", s.appHandler)
	appGroup.PUT("/:app", func(c *gin.Context) {
		s.mc.UpdateApp(context.Offer(c))
	})
	appGroup.GET("/:app/Page", s.appHandler)
	appGroup.GET("/:app/Page/:page", s.appHandler)

	metaGroup.POST("/App/:app/UploadModel", func(c *gin.Context) {
		s.mc.UploadModel(context.Offer(c))
	})

	modelGroup := metaGroup.Group("/App/:app/Model")
	modelGroup.POST("/", func(c *gin.Context) {
		s.mc.CreateModel(context.Offer(c))
	})
	modelGroup.PUT("/:model", func(c *gin.Context) {
		s.mc.UpdateModel(context.Offer(c))
	})
	modelGroup.GET("/", func(c *gin.Context) {
		s.mc.ListModels(context.Offer(c))
	})
	modelGroup.GET("/:model", func(c *gin.Context) {
		s.mc.ShowModel(context.Offer(c))
	})
	modelGroup.POST(":model/Column/", func(c *gin.Context) {
		s.mc.AddColumn(context.Offer(c))
	})
	modelGroup.DELETE(":model/Column/", func(c *gin.Context) {
		s.mc.DeprecateColumn(context.Offer(c))
	})
	modelGroup.PUT(":model/Column/", func(c *gin.Context) {
		s.mc.ReorderColumns(context.Offer(c))
	})
	modelGroup.PUT(":model/Column/:column", func(c *gin.Context) {
		s.mc.UpdateColumn(context.Offer(c))
	})
}

func (s *Server) appHandler(c *gin.Context) {
	if c.Request.Method == http.MethodPost {
		s.mc.CreateApp(context.Offer(c))
		return
	}

	if strings.HasSuffix(c.Request.RequestURI, "Page") {
		s.mc.ShowPages(context.Offer(c))
		return
	}

	pageIDStr := c.Param("page")
	if pageIDStr != "" {
		s.mc.ShowPage(context.Offer(c))
		return
	}

	s.mc.ShowApp(context.Offer(c))
}

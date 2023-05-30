package server

import (
	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/gin-gonic/gin"
)

func (s *Server) exportDataAPIs() {
	dataGroup := s.router.Group("/api/v0.1/data")

	dataGroup.GET("/lookup/Model/:model/Slot/:slot", func(c *gin.Context) {
		s.dc.Lookup(context.Offer(c))
	})

	dataGroup.POST("/App/:app", func(c *gin.Context) {
		s.dc.ImportRows(context.Offer(c))
	})

	dataGroup.POST("/Model/:model/", func(c *gin.Context) {
		s.dc.CreateRow(context.Offer(c))
	})
	dataGroup.PUT("/Model/:model/", func(c *gin.Context) {
		s.dc.QuickSave(context.Offer(c))
	})
	dataGroup.GET("/Model/:model/Row/:row", s.rowHandler)
	dataGroup.GET("/Model/:model/Row", s.rowHandler)
	dataGroup.PUT("/Model/:model/Row/:row", func(c *gin.Context) {
		s.dc.UpdateRow(context.Offer(c))
	})
	dataGroup.DELETE("/Model/:model/Row/:row", func(c *gin.Context) {
		s.dc.DeleteRow(context.Offer(c))
	})

	dataGroup.POST("/Model/:model/Slot/:slot/Picklist", func(c *gin.Context) {
		s.dc.CreatePickItem(context.Offer(c))
	})
	dataGroup.PUT("/Model/:model/Slot/:slot/Picklist", func(c *gin.Context) {
		s.dc.UpdatePickItem(context.Offer(c))
	})
	dataGroup.GET("/Model/:model/Slot/:slot/Picklist", func(c *gin.Context) {
		s.dc.ShowPicklist(context.Offer(c))
	})
}

// layered routing
func (s *Server) rowHandler(c *gin.Context) {
	row := c.Param("row")
	// list rows by filter
	if row == "" {
		s.dc.FindRows(context.Offer(c))
		return
	}

	// retrieve a single row
	s.dc.RetrieveRow(context.Offer(c))
}

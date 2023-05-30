package data

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/pkg/dto"
	"github.com/agile-app/flexdb/pkg/store"
	"github.com/funkygao/log4go"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func (dc *dataHandler) CreateRow(c context.RESTContext) {
	var rd dto.RowData
	if err := c.Gin().ShouldBindJSON(&rd); err != nil {
		c.AbortWithError(err)
		return
	}

	ms, err := store.Provider.InferAppStore(c).LoadModel(c.ModelID(), "Triggers", "Slots")
	if err != nil {
		c.AbortWithError(err)
		return
	}

	rd.SetCUser(c.PIN())
	rd.SetMUser(c.PIN())
	rowID, err := ms.DS().CreateRow(rd)
	if err != nil {
		// https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html
		if me, yes := err.(*mysql.MySQLError); yes && me.Number == 1062 {
			c.AbortWithError(errDupEntry)
			return
		}

		c.AbortWithError(err)
		return
	}

	c.RenderOK(rowID)
}

func (dc *dataHandler) ImportRows(c context.RESTContext) {
	// import rows by uploading excel
	g := c.Gin()
	file, err := g.FormFile("excel")
	if err != nil {
		c.AbortWithError(err)
		return
	}

	f, err := file.Open()
	if err != nil {
		c.AbortWithError(err)
		return
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f) // TODO OOM
	if err != nil {
		c.AbortWithError(err)
		return
	}

	os := store.Provider.InferOrgStore(c)
	app, err := os.LoadApp(c.AppID(), "Models.Slots")
	if err != nil {
		c.AbortWithError(err)
		return
	}

	ok, fail, err := dc.importDataFromExcel(b, app, os.AppStoreOf(app.ID))
	if err != nil {
		c.AbortWithError(err)
		return
	}

	c.RenderWithMsg(fmt.Sprintf("成功导入：%d，失败：%d 条", ok, fail))
}

func (dc *dataHandler) QuickSave(c context.RESTContext) {
	var qs dto.QuickSave
	if err := c.Gin().ShouldBindJSON(&qs); err != nil {
		c.AbortWithError(err)
		return
	}

	log4go.Info("%+v", qs)
	c.RenderOKWithoutData()
}

func (dc *dataHandler) UpdateRow(c context.RESTContext) {
	var rd dto.RowData
	if err := c.Gin().ShouldBindJSON(&rd); err != nil {
		c.AbortWithError(err)
		return
	}

	ms, err := store.Provider.InferAppStore(c).LoadModel(c.ModelID(), "Slots")
	if err != nil {
		c.AbortWithError(err)
		return
	}

	if err = ms.DS().UpdateRow(rd); err != nil {
		c.AbortWithError(err)
		return
	}

	c.RenderOKWithoutData()
}

func (dc *dataHandler) DeleteRow(c context.RESTContext) {
	ms, err := store.Provider.InferAppStore(c).LoadModel(c.ModelID(), "Slots")
	if err != nil {
		c.AbortWithError(err)
		return
	}

	if err = ms.DS().DeleteRow(c.RowID()); err != nil {
		c.AbortWithError(err)
		return
	}

	c.RenderOKWithoutData()
}

func (dc *dataHandler) RetrieveRow(c context.RESTContext) {
	ms, err := store.Provider.InferAppStore(c).LoadModel(c.ModelID(), "Slots")
	if err != nil {
		c.AbortWithError(err)
	}

	withChildren := false
	if c.Gin().Query("details") != "" {
		withChildren = true
	}
	md, err := ms.RetrieveRow(c.RowID(), withChildren)
	if err != nil {
		c.AbortWithError(err)
		return
	}

	c.RenderOK(md.Master)
}

func (dc *dataHandler) FindRows(c context.RESTContext) {
	ms, err := store.Provider.InferAppStore(c).LoadModel(c.ModelID(), "Slots")
	if err != nil {
		log4go.Error(err.Error())
		c.AbortWithError(err)
		return
	}

	criteria := c.SearchCriteria()
	var selectFields []string
	for _, c := range criteria {
		if c.Key == "_select" { // TODO will use explicit select fields in REST
			selectFields = strings.Split(c.Val, ",")
		}
	}

	rows, err := ms.DS().FindRows(selectFields, criteria, c.PageIndex(), c.PageSize())
	if err != nil {
		log4go.Error(err.Error())
		c.AbortWithError(err)
		return
	}

	if c.AMIS() {
		c.RenderOK(gin.H{"items": rows, "hasNext": len(rows) == c.PageSize()})
	} else {
		c.RenderOK(rows)
	}

}

func (dc *dataHandler) Lookup(c context.RESTContext) {
	ms, err := store.Provider.InferAppStore(c).LoadModel(c.ModelID(), "Slots")
	if err != nil {
		c.AbortWithError(err)
		return
	}

	slotID, err := strconv.ParseInt(c.Gin().Param("slot"), 10, 64)
	if err != nil {
		c.AbortWithError(err)
		return
	}

	var slotName string
	for _, c := range ms.EntityModel().Slots {
		if c.Slot == int16(slotID) {
			slotName = c.Name
			break
		}
	}

	rows, err := ms.DS().FindRows(nil, nil, 1, 200)
	if err != nil {
		c.AbortWithError(err)
		return
	}
	options := make([]dto.LabelValue, len(rows))
	for i, r := range rows {
		options[i] = dto.LabelValue{
			Label: r.StrValueOf(slotName),
			Value: strconv.FormatUint(r.ID(), 10),
		}
	}

	c.RenderOK(gin.H{"options": options})
}

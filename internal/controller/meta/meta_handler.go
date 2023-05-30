package meta

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/pkg/dto"
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/store"
	"github.com/funkygao/log4go"
	"github.com/gin-gonic/gin"
)

// TODO transaction
func (mc *metaHandler) CreateApp(c context.RESTContext) {
	// create app by uploading excel
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

	visibility, err := entity.ToVisibility(g.Request.FormValue("visibility"))
	app := &entity.App{
		Name:        g.Request.FormValue("name"),
		Description: g.Request.FormValue("description"),
		Logo:        g.Request.FormValue("logo"),
		Visibility:  visibility,
	}
	os := store.Provider.InferOrgStore(c)
	if err := os.CreateApp(app); err != nil {
		c.AbortWithError(err)
		return
	}

	as := os.AppStoreOf(app.ID)
	if err = mc.createModelsAndColumnsFromExcel(b, app, as, c.PIN()); err != nil {
		c.AbortWithError(err)
		return
	}

	if err = mc.createPresetModels(app, as, c); err != nil {
		c.AbortWithError(err)
		return
	}

	c.RenderOK(app)
}

func (mc *metaHandler) UpdateApp(c context.RESTContext) {
	var dto entity.App
	if err := c.Gin().ShouldBindJSON(&dto); err != nil {
		c.AbortWithError(err)
		return
	}

	os := store.Provider.InferOrgStore(c)
	app, err := os.LoadApp(dto.ID, "Models.Slots")
	if err != nil {
		c.AbortWithError(err)
		return
	}

	// merge with dto
	app.Name = dto.Name
	app.Description = dto.Description
	app.MTime = time.Now()
	app.MUser = c.PIN()
	app.Visibility = dto.Visibility
	if err := os.UpdateApp(app); err != nil {
		c.AbortWithError(err)
		return
	}

	c.RenderOKWithoutData()
}

func (mc *metaHandler) FindApps(c context.RESTContext) {
	cond := make(dto.Criteria, 0)
	if q, present := c.Gin().GetQuery("q"); present {
		switch q {
		case "private":
			cond = cond.Append(dto.CriteriaItem{
				Key: "cuser",
				Op:  "=",
				Val: c.PIN(),
			})
			cond = cond.Append(dto.CriteriaItem{
				Key: "visibility",
				Op:  "=",
				Val: entity.PrivateVisibility.String(),
			})

		case "shared":

		case "recent":

		case "public":
			cond = cond.Append(dto.CriteriaItem{
				Key: "visibility",
				Op:  "=",
				Val: entity.PublicVisibility.String(),
			})
		}
	}
	apps, err := store.Provider.InferOrgStore(c).
		FindApps(cond, c.PageIndex(), c.PageSize())
	if err != nil {
		c.AbortWithError(err)
		return
	}

	c.RenderOK(gin.H{"items": apps})
}

func (mc *metaHandler) ShowApp(c context.RESTContext) {
	app, err := store.Provider.InferOrgStore(c).LoadApp(c.AppID(), "Models.Slots")
	if err != nil {
		c.AbortWithError(err)
		return
	}

	c.RenderOK(gin.H{"app": app})
}

func (mc *metaHandler) ShowPages(c context.RESTContext) {

}

func (mc *metaHandler) ShowPage(c context.RESTContext) {
}

func (mc *metaHandler) CreateModel(c context.RESTContext) {
	req := c.Gin().Request
	now := time.Now()
	model := entity.Model{
		Name:  req.Header.Get("TableName"),
		CTime: now,
		MTime: now,
		CUser: c.PIN(),
		AppID: c.AppID(),
		Ver:   1,
	}
	if err := store.Provider.InferAppStore(c).CreateModel(&model); err != nil {
		c.AbortWithError(err)
		return
	}

	c.RenderOK(model)
}

func (mc *metaHandler) UpdateModel(c context.RESTContext) {
	var model entity.Model
	if err := c.Gin().ShouldBindJSON(&model); err != nil {
		c.AbortWithError(err)
		return
	}

	as := store.Provider.InferAppStore(c)
	if as == nil {
		c.AbortWithError(fmt.Errorf("app:%d not found", c.AppID()))
		return
	}

	model.MTime = time.Now()
	model.MUser = c.PIN()
	model.AppID = c.AppID()
	features := map[string]func(bool){
		"createRow": model.Feature.EnableCreateRow,
		"readRow":   model.Feature.EnableReadRow,
		"updateRow": model.Feature.EnableUpdateRow,
		"deleteRow": model.Feature.EnableDeleteRow,
	}
	for k, f := range features {
		if v, present := model.H[k]; present {
			f(v.(bool))
		}
	}

	if err := as.UpdateModel(&model); err != nil {
		c.AbortWithError(err)
		return
	}

	c.RenderOKWithoutData()
}

func (mc *metaHandler) UploadModel(c context.RESTContext) {
	req := c.Gin().Request
	file, header, err := req.FormFile("file1") // file{N}
	if err != nil {
		c.AbortWithError(errors.New("file1 not present"))
		return
	}

	b := make([]byte, header.Size)
	io.ReadFull(file, b)
	var model entity.Model
	json.Unmarshal(b, &model)

	// enrich model attributes
	model.Name = req.Header.Get("TableName") // TODO
	now := time.Now()
	model.CTime = now
	model.MTime = now
	model.Deleted = false
	model.AppID = c.AppID()
	model.CUser = c.PIN()
	model.Ver = 1
	model.Feature.EnableCRUD()

	as := store.Provider.InferAppStore(c)
	if err := as.CreateModel(&model); err != nil {
		c.AbortWithError(err)
		return
	}

	// create columns
	ms := as.ModelStoreOf(&model)
	slots := model.Slots // save uploaded slots
	// reset model.slots: otherwise, slots will be added twice
	model.Slots = make([]*entity.Column, 0, len(slots))
	if err := ms.AddColumns(slots); err != nil {
		c.AbortWithError(err)
		return
	}

	log4go.Info("model[%s] created, id:%d", model.Name, model.ID)
	c.RenderOK(model)
}

func (mc *metaHandler) ShowModel(c context.RESTContext) {
	model, err := store.Provider.InferAppStore(c).LoadModel(c.ModelID(), "Slots")
	if err != nil {
		c.AbortWithError(err)
		return
	}

	c.RenderOK(model)
}

func (mc *metaHandler) ListModels(c context.RESTContext) {
	models, err := store.Provider.InferAppStore(c).LoadModels()
	if err != nil {
		c.AbortWithError(err)
		return
	}

	c.RenderOK(models)
}

func (mc *metaHandler) AddColumn(c context.RESTContext) {
	var col entity.Column
	if err := c.Gin().ShouldBindJSON(&col); err != nil {
		c.AbortWithError(err)
		return
	}

	col.CUser = c.PIN()
	col.CTime = time.Now()

	if col.Relational() {
		tuple := strings.SplitN(col.H["ref"].(string), ",", 2)
		if len(tuple) != 2 {
			c.AbortWithError(fmt.Errorf("Invalid ref model info"))
			return
		}

		refModelID, err := strconv.ParseInt(tuple[0], 10, 64)
		if err != nil {
			c.AbortWithError(err)
			return
		}
		refSlotID, err := strconv.ParseInt(tuple[1], 10, 64)
		if err != nil {
			c.AbortWithError(err)
			return
		}

		if refModelID < 1 || refSlotID < 1 {
			c.AbortWithError(fmt.Errorf("invalid refModel or refSlot"))
			return
		}

		col.RefModelID = refModelID
		col.RefSlot = int16(refSlotID)
	}

	ms, err := store.Provider.InferAppStore(c).LoadModel(c.ModelID(), "Slots")
	if err != nil {
		c.AbortWithError(err)
		return
	}

	if err = ms.AddColumns([]*entity.Column{&col}); err != nil {
		c.AbortWithError(err)
		return
	}

	c.RenderOKWithoutData()
}

func (mc *metaHandler) UpdateColumn(c context.RESTContext) {
	var dto entity.Column
	if err := c.Gin().ShouldBindJSON(&dto); err != nil {
		c.AbortWithError(err)
		return
	}

	ms, err := store.Provider.InferAppStore(c).LoadModel(c.ModelID(), "Slots")
	if err != nil {
		c.AbortWithError(err)
		return
	}

	col := ms.EntityModel().SlotByColumnID(c.ColumnID())
	if col == nil {
		c.AbortWithError(errors.New("invalid column"))
		return
	}

	// merge DTO with storage
	col.Label = dto.Label
	col.Name = dto.Name
	col.Choices = dto.Choices
	col.Remark = dto.Remark
	col.ReadOnly = dto.ReadOnly
	col.Required = dto.Required
	col.Choices = dto.Choices

	// audit
	col.MUser = c.PIN()
	col.MTime = time.Now()
	if err = ms.UpdateColumn(col); err != nil {
		c.AbortWithError(err)
		return
	}

	c.RenderOKWithoutData()
}

func (mc *metaHandler) ReorderColumns(c context.RESTContext) {
	type postData struct {
		IDs string `json:"ids"`
	}
	var dto postData
	if err := c.Gin().ShouldBindJSON(&dto); err != nil {
		c.AbortWithError(err)
		return
	}

	ms, err := store.Provider.InferAppStore(c).LoadModel(c.ModelID())
	if err != nil {
		c.AbortWithError(err)
		return
	}

	// builtin columns not able to order

	log4go.Info("%+v", dto)

	ids := strings.Split(dto.IDs, ",")
	columns := make([]entity.Column, 0, len(ids))
	for seq, id := range ids {
		if id == "0" {
			// builtin column
			continue
		}

		columnID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.AbortWithError(err)
			return
		}

		columns = append(columns, entity.Column{
			ID:      columnID,
			ModelID: c.ModelID(),
			Ordinal: int16(seq) + 1,
			MUser:   c.PIN(),
			MTime:   time.Now(),
		})
	}

	if err := ms.ReorderColumns(columns); err != nil {
		c.AbortWithError(err)
		return
	}

	c.RenderOKWithoutData()
}

func (mc *metaHandler) DeprecateColumn(c context.RESTContext) {
	ms, err := store.Provider.InferAppStore(c).LoadModel(c.ModelID(), "Slots")
	if err != nil {
		c.AbortWithError(err)
		return
	}

	col := ms.EntityModel().SlotByColumnID(c.QueryID())
	if col == nil {
		c.AbortWithError(errors.New("invalid column"))
		return
	}

	col.Deprecated = true
	col.MTime = time.Now()
	col.MUser = c.PIN()
	if err = ms.DeprecateColumn(col); err != nil {
		c.AbortWithError(err)
		return
	}

	c.RenderOKWithoutData()
}

func (mc *metaHandler) FindTemplates(c context.RESTContext) {
	type template struct {
		Name     string `json:"name"`
		Desc     string `json:"desc"`
		Creator  string `json:"creator"`
		Size     string `json:"size"`
		Link     string `json:"link"`
		Terminal string `json:"terminal"`
	}
	templates := []template{
		{
			Name:     "从零开始搭建应用",
			Desc:     "能执行的Excel！请阅读该Excel模版的说明，从零开始设计你的应用程序",
			Creator:  "FunkyGao",
			Size:     "1.31MB",
			Link:     "/static/images/template.xlsx",
			Terminal: "PC Mobile",
		},
	}

	c.RenderOK(gin.H{"items": templates})
}

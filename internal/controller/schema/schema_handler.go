package schema

import (
	"errors"

	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/store"
	"github.com/agile-app/flexdb/pkg/view/design"
	"github.com/agile-app/flexdb/pkg/view/runtime"
)

func (mc *schemaHandler) FindApps(c context.RESTContext) {
	page := design.RenderApps(c)
	c.RenderOK(page)
}

func (mc *schemaHandler) ShowApp(c context.RESTContext) {
	app, err := store.Provider.InferOrgStore(c).LoadApp(c.AppID(), "Models.Slots", "Shares")
	if err != nil {
		c.AbortWithError(err)
		return
	}

	if app.UUID != c.UUID() {
		c.AbortWithError(errors.New("denied"))
		return
	}

	page := design.RenderApp(c, app)
	c.RenderOK(page)
}

func (mc *schemaHandler) ModelSchemaCRUD(c context.RESTContext) {
	app, err := store.Provider.InferOrgStore(c).LoadApp(c.AppID(), "Models.Slots")
	if err != nil {
		c.AbortWithError(err)
		return
	}

	as := store.Provider.InferAppStore(c)
	models, err := as.LoadModels("Slots")
	if err != nil {
		c.AbortWithError(err)
		return
	}

	var model *entity.Model
	for _, m := range models {
		if m.ID == c.ModelID() {
			model = &m
			break
		}
	}

	page := design.RenderModel(c, app, models, model)
	c.RenderOK(page)
}

func (mc *schemaHandler) ModelCRUD(c context.RESTContext) {
	app, err := store.Provider.InferOrgStore(c).LoadApp(c.AppID(), "Models.Slots")
	if err != nil {
		c.AbortWithError(err)
		return
	}

	as := store.Provider.InferAppStore(c)
	models, err := as.LoadModels("Slots")
	if err != nil {
		c.AbortWithError(err)
		return
	}

	filteredModels := models[:0]
	for _, m := range models {
		if !m.VisibleFor(c) {
			continue
		}

		filteredModels = append(filteredModels, m)
	}

	var (
		modelID = c.QueryID() // TODO bad design
		model   *entity.Model
	)
	if modelID > 0 {
		for _, m := range models {
			if m.ID == modelID {
				model = &m
				break
			}
		}
	} else {
		model = &filteredModels[0]
	}

	page := runtime.RenderModel(c, app, filteredModels, model)
	c.RenderOK(page)
}

package meta

import (
	"time"

	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/store"
)

func (m *metaHandler) createPresetModels(app *entity.App, as store.AppStore, ctx context.RESTContext) (err error) {
	for _, model := range entity.PermModels {
		now := time.Now()
		model.CTime = now
		model.MTime = now
		model.CUser = ctx.PIN()
		model.AppID = app.ID
		if err = as.CreateModel(&model); err != nil {
			return err
		}

		columns := make([]*entity.Column, len(model.Slots))
		for i, c := range model.Slots {
			c.CTime = now
			c.MTime = now
			c.CUser = model.CUser
			c.ID = 0 // TODO reset id, otherwise will raise dup key err
			columns[i] = c
		}
		ms := as.ModelStoreOf(&model)
		if err = ms.AddColumns(columns); err != nil {
			return err
		}
	}

	return
}

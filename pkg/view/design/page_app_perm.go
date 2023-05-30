package design

import (
	"strconv"

	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/pkg/api"
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/view"
	"github.com/agile-app/flexdb/pkg/view/amis"
)

func renderAppPerm(ctx context.RESTContext, app *entity.App, page *amis.Page, modelID string) {
	var model entity.Model
	id, _ := strconv.ParseInt(modelID, 10, 64)
	for _, m := range app.PermModels() {
		if m.ID == id {
			model = m
			break
		}
	}

	if model.ID == 0 {
		// bad request
		return
	}

	newBtn := amis.NewButton()
	page.AddBody(newBtn)
	newBtn.ActionType = "drawer"
	newBtn.Label = ctx.T("crud.button.create")
	newBtn.Icon = "fa fa-plus pull-left"
	newBtn.ClassName = "m-b-sm"

	newBtnDrawer := amis.NewDrawer()
	newBtn.Drawer = newBtnDrawer
	newBtnDrawer.CloseOnEsc = true
	newBtnDrawer.Size = "lg"

	createForm := amis.NewForm()
	createForm.Autofocus = true
	createForm.API = api.CreateRowAPI(model.ID)
	newBtnDrawer.Body = createForm

	crud := view.CRUD(ctx, &model, createForm)
	page.AddBody(crud)
}

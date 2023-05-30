package runtime

import (
	"fmt"

	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/pkg/api"
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/view"
	"github.com/agile-app/flexdb/pkg/view/amis"
)

// RenderModel renders https://baidu.github.io/amis/zh-CN/docs/components/table?page=1
func RenderModel(ctx context.RESTContext, app *entity.App, models []entity.Model, model *entity.Model) *amis.Page {
	page := amis.NewPage()
	titleIcon := amis.NewControl()
	titleIcon.Type = "icon"
	titleIcon.Icon = "database"
	titleIcon.ClassName = "text-info text-lg"
	titleText := amis.NewControl()
	titleText.Type = "tpl"
	titleText.ClassName = "m-l-sm"
	if model == nil {
		page.Title = "无权限" // filtered by upper layer
		return page
	}

	titleText.Tpl = model.Name
	page.Title = []interface{}{titleIcon, titleText}
	if model.Remark != "" {
		page.SubTitle = model.Remark
	}

	// aside nav
	aside := page.SetAside("wrapper")
	aside.Size = "xs"
	nav := amis.NewNav()
	nav.Stacked = true
	aside.AddBody(nav)

	for _, m := range models {
		if m.Builtin() || !m.VisibleFor(ctx) {
			continue
		}

		link := amis.NewLink(m.Name, api.RowCRUDPage(app.UUID, app.ID, m.ID))
		if m.ID == model.ID {
			link.Active = true
		}

		nav.AddLink(*link)
	}

	// toolbar
	toolbar := amis.NewToolbar()
	toolbar.Type = "button-group"
	page.AddToolbar(toolbar)

	var createForm *amis.Form
	if model.Feature.CreateRowEnabled() {
		newBtn := amis.NewButton()
		toolbar.AddButton(newBtn)
		newBtn.ActionType = "drawer"
		newBtn.Label = ctx.T("crud.button.create")
		newBtn.Icon = "fa fa-plus pull-left"
		newBtn.Primary = true

		// toolbar.addBtn.drawer
		newBtnDrawer := amis.NewDrawer()
		newBtn.Drawer = newBtnDrawer
		newBtnDrawer.Title = fmt.Sprintf("%s 数据录入", model.Name)
		newBtnDrawer.CloseOnEsc = true
		newBtnDrawer.Size = "lg"

		// toolbar.addBtn.drawer.form
		createForm = amis.NewForm()
		createForm.Autofocus = true
		createForm.API = api.CreateRowAPI(model.ID)
		newBtnDrawer.Body = createForm
	}

	crud := view.CRUD(ctx, model, createForm)
	page.AddBody(crud)

	return page
}

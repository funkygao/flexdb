package design

import (
	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/pkg/api"
	"github.com/agile-app/flexdb/pkg/dto"
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/view/amis"
)

func RenderApps(ctx context.RESTContext) *amis.Page {
	page := amis.NewPage()
	page.Title = ctx.T("design.apps.page.title")
	page.Remark = ctx.T("design.apps.page.remark")
	page.SubTitle = ctx.T("design.apps.page.subTitle")
	page.InitAPI = api.FindAppsAPI("recent")

	// create app button
	toolbar := amis.NewToolbar()
	page.AddToolbar(toolbar)
	toolbar.Type = amis.ATButton
	toolbar.ActionType = amis.ATDialog
	toolbar.Label = ctx.T("design.apps.page.create")
	toolbar.Icon = "fa fa-plus pull-left"
	toolbar.Primary = true

	// create app dialog
	createAppDialog := amis.NewDialog()
	createAppDialog.Title = ctx.T("design.apps.page.create")
	createAppDialog.Size = "lg"
	createAppDialog.CloseOnEsc = true
	toolbar.Dialog = createAppDialog
	// create app form
	createAppForm := amis.NewForm()
	createAppDialog.Body = createAppForm
	createAppForm.Autofocus = true
	createAppForm.API = api.UploadAppAPI()
	createAppForm.Reload = "window"
	// create app form controls
	controls := []dto.H{
		{"type": amis.ATText, "name": "name", "label": "应用名称", "desc": "支持中文名称，但不能包含特殊符号和空格", "required": true, "placeholder": ""},
		{"type": amis.ATTextArea, "name": "description", "label": "应用描述", "desc": "请用简短的语言描述这个应用的用途", "required": true},

		{"type": amis.ATRadios, "name": "visibility", "label": "可见性", "required": true, "value": entity.PrivateVisibility,
			"options": []dto.H{
				{"label": ctx.T("app.visibility.public.label"), "value": entity.PublicVisibility, "description": ctx.T("app.visibility.public.desc")},
				{"label": ctx.T("app.visibility.private.label"), "value": entity.PrivateVisibility, "description": ctx.T("app.visibility.private.desc")},
			}},
		{"type": amis.ATDivider},
		{"type": amis.ATImage, "name": "logo", "label": "Logo"},
		{"type": amis.ATFile, "name": "excel", "label": "设计模版", "remark": "就是你在应用市场下载的模版编辑后的Excel文件", "asBlob": true, "required": true, "accept": ".xlsx"},
	}
	for _, c := range controls {
		ctrl := amis.NewControl()
		createAppForm.AddControl(ctrl)
		ctrl.Type = c.S("type")
		if s := c.S("name"); s != "" {
			ctrl.Name = s
		}
		if s := c.S("desc"); s != "" {
			ctrl.Desc = s
		}
		if s := c.S("label"); s != "" {
			ctrl.Label = s
		}
		if s := c.S("placeholder"); s != "" {
			ctrl.Placeholder = s
		}
		if b := c.B("required"); b {
			ctrl.Required = b
		}
		if b := c.B("asBlob"); b {
			ctrl.AsBlob = b
		}
		if v, present := c.V("options"); present {
			ctrl.Options = v
		}
		if v, present := c.V("value"); present {
			ctrl.Value = v
		}
		if s := c.S("remark"); s != "" {
			ctrl.Remark = s
		}
		if s := c.S("accept"); s != "" {
			ctrl.Accept = s
		}
	}

	// page body tabs
	tabsCtrl := amis.NewControl()
	tabsCtrl.Type = amis.ATTabs
	page.AddBody(tabsCtrl)
	tabs := []dto.H{
		{"title": "最近打开"}, // has no api, its default api source
		{"api": api.FindAppsAPI("private"), "title": "我创建的"},
		{"api": api.FindAppsAPI("shared"), "title": ctx.T("app.visibility.shared.label")},
		{"api": api.FindAppsAPI("public"), "title": ctx.T("app.visibility.public.label")},
	}
	for _, t := range tabs {
		tab := amis.NewTab()
		tab.Title = t.S("title")
		tabsCtrl.AddTab(tab)
		crud := amis.NewCrud()
		tab.Body = crud
		crud.Reset()

		crud.Mode = amis.ATCards
		crud.ColumnsCount = 3
		crud.PerPageAvailable = nil
		crud.FilterDefaultVisible = false
		if s := t.S("api"); s != "" {
			crud.API = t.S("api")
		}

		card := amis.NewControl()
		crud.Card = card
		card.Title = t.S("title")
		card.Type = amis.ATCard
		card.Header = dto.H{
			"title":           "${name}",
			"avatar":          "/static/images/app.png",
			"avatarClassName": "pull-left thumb b-3x m-r",
		}
		card.Body = []dto.H{
			{"name": "${statusLabel}", "label": "状态"},
			{"name": "${cuser}", "label": "创建人"},
			{"name": "${pctime}", "label": "创建时间"},
			{"name": "${muser}", "label": "修改人"},
			{"name": "${pmtime}", "label": "最近修改"},
		}

		// card actions: launch app
		launch := amis.NewControl()
		card.AddAction(launch)
		launch.Label = "启动应用 "
		launch.Type = amis.ATAction
		launch.ActionType = amis.ATUrl
		launch.Icon = "fa fa-external-link"
		launch.URL = "/crud?app=${id}&uuid=${uuid}" // TODO kill app.id

		// card actions: build app
		design := amis.NewControl()
		card.AddAction(design)
		design.Label = "设计应用 "
		design.Type = amis.ATAction
		design.ActionType = amis.ATUrl
		design.Icon = "fa fa-code"
		design.URL = "/design?id=${id}&menu=models&uuid=${uuid}" // TODO kill app.id
	}

	return page
}

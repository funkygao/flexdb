package design

import (
	"fmt"
	"strings"

	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/pkg/api"
	"github.com/agile-app/flexdb/pkg/dto"
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/view/amis"
)

func RenderApp(ctx context.RESTContext, app *entity.App) *amis.Page {
	page := amis.NewPage()
	titleIcon := amis.NewControl()
	titleIcon.Type = "icon"
	titleIcon.Icon = "sitemap"
	titleIcon.ClassName = "text-info text-lg"
	titleText := amis.NewControl()
	titleText.Type = "tpl"
	titleText.ClassName = "m-l-sm"
	titleText.Tpl = app.Name
	page.Title = []interface{}{titleIcon, titleText}
	page.SubTitle = app.Description

	// toobar 应用设置
	toolbar := amis.NewToolbar()
	page.AddToolbar(toolbar)
	toolbar.Type = amis.ATButton
	toolbar.ActionType = "dialog"
	toolbar.Label = "应用设置"
	toolbar.Icon = "fa fa-cog pull-left"

	// app edit dialog
	appEditDialog := amis.NewDialog()
	appEditDialog.Title = "应用设置"
	appEditDialog.Size = "lg"
	appEditDialog.CloseOnEsc = true
	toolbar.Dialog = appEditDialog
	// app edit form
	appEditForm := amis.NewForm()
	appEditDialog.Body = appEditForm
	appEditForm.API = api.UpdateAppAPI(app.ID)
	appEditForm.Reload = "window"
	appEditForm.Data = dto.H{
		"id":         app.ID,
		"name":       app.Name,
		"desc":       app.Description,
		"visibility": app.Visibility,
	}
	// edit form tabs
	tabs := amis.NewControl()
	tabs.Type = amis.ATTabs
	appEditForm.AddControl(tabs)

	tab := amis.NewTab()
	tabs.AddTab(tab)
	tab.Title = "基本设置"
	ctrl := amis.NewControl()
	ctrl.Name = "name"
	ctrl.Type = amis.ATText
	ctrl.Label = "应用名称"
	ctrl.Required = true
	tab.AddControl(ctrl)
	ctrl = amis.NewControl()
	ctrl.Name = "desc"
	ctrl.Type = amis.ATTextArea
	ctrl.Label = "应用说明"
	ctrl.Required = true
	tab.AddControl(ctrl)
	ctrl = amis.NewControl()
	ctrl.Name = "visibility"
	ctrl.Type = amis.ATRadios
	ctrl.Value = app.Visibility
	ctrl.Options = []dto.H{
		{
			"description": ctx.T("app.visibility.public.desc"),
			"label":       ctx.T("app.visibility.public.label"),
			"value":       entity.PublicVisibility,
		},
		{
			"description": ctx.T("app.visibility.private.desc"),
			"label":       ctx.T("app.visibility.private.label"),
			"value":       entity.PrivateVisibility,
		},
	}
	ctrl.Label = ctx.T("app.visibility.label")
	ctrl.Required = true
	tab.AddControl(ctrl)
	ctrl = amis.NewControl()
	ctrl.Name = "id"
	ctrl.Type = "hidden"
	tab.AddControl(ctrl)

	tab = amis.NewTab()
	tab.Title = "应用成员"
	tabs.AddTab(tab)
	ctrl = amis.NewControl()
	ctrl.Name = "x"
	ctrl.Type = "text"
	ctrl.Placeholder = "请输入erp"
	ctrl.Label = "主管理员"
	tab.AddControl(ctrl)
	ctrl = amis.NewControl()
	ctrl.Name = "y"
	ctrl.Type = "text"
	ctrl.Label = "角色"
	tab.AddControl(ctrl)
	ctrl = amis.NewControl()
	ctrl.Name = "z"
	ctrl.Type = "text"
	ctrl.Placeholder = "请输入erp"
	ctrl.Label = "数据管理员"
	tab.AddControl(ctrl)
	ctrl = amis.NewControl()
	ctrl.Name = "z1"
	ctrl.Type = "text"
	ctrl.Placeholder = "开发成员"
	ctrl.Label = "数据管理员"
	tab.AddControl(ctrl)

	tab = amis.NewTab()
	tab.Title = "接口对接"
	tabs.AddTab(tab)
	ctrl = amis.NewControl()
	ctrl.Name = "z5"
	ctrl.Type = "text"
	ctrl.Label = "AppKey"
	tab.AddControl(ctrl)
	ctrl = amis.NewControl()
	ctrl.Name = "z9"
	ctrl.Type = "text"
	ctrl.Label = "AppSecret"
	tab.AddControl(ctrl)

	// toobar 创建新表
	toolbar = amis.NewToolbar()
	page.AddToolbar(toolbar)
	toolbar.Type = amis.ATButton
	toolbar.ActionType = "dialog"
	toolbar.Icon = "fa fa-plus pull-left"
	toolbar.Label = "建新表"
	toolbar.Tooltip = "请上传Excel作为新表模版"
	addModelDialog := amis.NewDialog()
	toolbar.Dialog = addModelDialog
	addModelDialog.CloseOnEsc = true
	addModelDialog.Size = "lg"
	addModelDialog.Title = "创建新表"
	addModelForm := amis.NewForm()
	addModelDialog.Body = addModelForm
	addModelForm.API = api.UpdateAppAPI(app.ID) // TODO
	ctrl = amis.NewControl()
	ctrl.Type = amis.ATFile
	ctrl.Name = "excel"
	ctrl.Required = true
	ctrl.Accept = ".xlsx"
	ctrl.Remark = "请认真核对Excel后操作"
	ctrl.AsBlob = true
	ctrl.Label = "Excel模版文件"
	addModelForm.AddControl(ctrl)
	ctrl = amis.NewControl()
	ctrl.Type = amis.ATHidden
	ctrl.Name = "addModel"
	ctrl.Value = true
	addModelForm.AddControl(ctrl)

	// toobar 下载模版
	toolbar = amis.NewToolbar()
	page.AddToolbar(toolbar)
	toolbar.Type = amis.ATButton
	toolbar.ActionType = amis.AATUrl
	toolbar.Link = ""
	toolbar.Icon = "fa fa-download pull-left"
	toolbar.Label = "下载模版"
	toolbar.Tooltip = "一次性下载当前应用的所有建表模版"

	// toobar 导入数据
	toolbar = amis.NewToolbar()
	page.AddToolbar(toolbar)
	toolbar.Type = amis.ATButton
	toolbar.ActionType = "dialog"
	toolbar.Icon = "fa fa-upload pull-left"
	toolbar.Label = "导入数据"
	toolbar.Tooltip = "上传包含_mapping_的Excel文件"
	uploadDialog := amis.NewDialog()
	toolbar.Dialog = uploadDialog
	uploadDialog.CloseOnEsc = true
	uploadDialog.Title = "导入数据"
	uploadForm := amis.NewForm()
	uploadDialog.Body = uploadForm
	uploadForm.API = api.ImportRowsAPI(app.ID)
	ctrl = amis.NewControl()
	ctrl.Type = "static"
	ctrl.Label = "注意事项"
	ctrl.Tpl = "数据文件通过sheet: _mapping_ 来定义转换规则<br/><br/>_mapping_ sheet必须是第一个(最左)sheet<br/>每个数据表sheet的表头占2行"
	uploadForm.AddControl(ctrl)
	ctrl = amis.NewControl()
	ctrl.Type = amis.ATFile
	ctrl.Name = "excel"
	ctrl.Required = true
	ctrl.Accept = ".xlsx"
	ctrl.Remark = "请认真核对Excel后操作"
	ctrl.AsBlob = true
	ctrl.Label = "数据文件"
	uploadForm.AddControl(ctrl)

	// aside nav
	aside := page.SetAside("wrapper")
	aside.Size = "xs"
	basicNav := amis.NewNav()
	basicNav.Stacked = true
	lk := amis.NewLink(ctx.T("design.app.menu.share"), api.AppDesignPage(app.UUID, app.ID, fmt.Sprintf("share:%d", app.ID)))
	lk.Icon = "fa fa-share-alt"
	basicNav.AddLink(*lk)
	lk = amis.NewLink(ctx.T("design.app.menu.preview"), "")
	lk.Icon = "fa fa-eye"
	lk.To = api.RowCRUDPage(app.UUID, app.ID, 0)
	basicNav.AddLink(*lk)
	lk = amis.NewLink(ctx.T("design.app.menu.api"), "")
	lk.Icon = "fa fa-book"
	lk1 := amis.NewLink("元数据", api.AppDesignPage(app.UUID, app.ID, "api:meta"))
	lk.AddChild(*lk1)
	lk1 = amis.NewLink("业务数据", api.AppDesignPage(app.UUID, app.ID, "api:data"))
	lk.AddChild(*lk1)
	basicNav.AddLink(*lk)
	aside.AddBody(basicNav)

	aside.AddBody(amis.NewDivider())

	nav := amis.NewNav()
	aside.AddBody(nav)
	nav.Stacked = true

	renderAppDesignMenu(ctx, app, nav)
	renderAppDesignBody(ctx, app, page)

	return page
}

func renderAppDesignMenu(ctx context.RESTContext, app *entity.App, nav *amis.Nav) {
	models := amis.NewLink(ctx.T("design.app.menu.tables"), api.AppDesignPage(app.UUID, app.ID, "models"))
	models.Icon = "fa fa-database"
	nav.AddLink(*models)

	perms := amis.NewLink(ctx.T("design.app.menu.perm"), "")
	perms.Icon = "fa fa-lock"
	for _, m := range app.PermModels() {
		lk := amis.NewLink(m.Name, api.AppDesignPage(app.UUID, app.ID, fmt.Sprintf("perm:%d", m.ID)))
		perms.AddChild(*lk)
	}
	nav.AddLink(*perms)

	pages := amis.NewLink(ctx.T("design.app.menu.pages"), "")
	pages.Icon = "fa fa-film"
	for _, p := range app.Pages {
		link := amis.NewLink(p.Name, "?id=x")
		link.Disabled = true
		pages.AddChild(*link)
	}
	nav.AddLink(*pages)

	charts := amis.NewLink(ctx.T("design.app.menu.charts"), "")
	charts.Icon = "fa fa-line-chart"
	for _, c := range app.Charts {
		link := amis.NewLink(c.Name, "")
		charts.AddChild(*link)
	}
	nav.AddLink(*charts)

	mobile := amis.NewLink("京ME入口", "")
	mobile.Icon = "fa fa-mobile"
	mobile.Disabled = true
	nav.AddLink(*mobile)

	flow := amis.NewLink("流程自动", "")
	flow.Icon = "fa fa-tasks"
	flow.Disabled = true
	nav.AddLink(*flow)

	triggers := amis.NewLink("触发器管理", "")
	triggers.Icon = "fa fa-random"
	nav.AddLink(*triggers)

	addons := amis.NewLink("插件管理", "")
	addons.Icon = "fa fa-plug"
	addon := amis.NewLink("字段类型", "")
	addons.AddChild(*addon)
	addon = amis.NewLink("触发器", "")
	addons.AddChild(*addon)
	addon = amis.NewLink("连接器", "")
	addons.AddChild(*addon)
	addon = amis.NewLink("过滤器", "")
	addons.AddChild(*addon)
	nav.AddLink(*addons)

	tools := amis.NewLink("平台工具", "")
	tools.Icon = "fa fa-cube"
	lc := amis.NewLink("发布管理", "")
	tools.AddChild(*lc)
	lc = amis.NewLink("数据源开发", "")
	tools.AddChild(*lc)
	lc = amis.NewLink("集成对接", "#")
	tools.AddChild(*lc)
	lc = amis.NewLink("存储空间使用情况", "#")
	tools.AddChild(*lc)
	lc = amis.NewLink("访问统计", "#")
	tools.AddChild(*lc)
	lc = amis.NewLink("出错报告", "#")
	tools.AddChild(*lc)
	lc = amis.NewLink("系统监控", "#")
	tools.AddChild(*lc)
	nav.AddLink(*tools)
}

func renderAppDesignBody(ctx context.RESTContext, app *entity.App, page *amis.Page) {
	tuple := strings.Split(ctx.Gin().Query("menu"), ":")
	if f, present := appMenuRenderers[tuple[0]]; present {
		if len(tuple) > 1 {
			f(ctx, app, page, tuple[1])
		} else {
			f(ctx, app, page, "")
		}
	}
}

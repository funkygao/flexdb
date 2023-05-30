package design

import (
	"fmt"

	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/pkg/api"
	"github.com/agile-app/flexdb/pkg/dto"
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/view/amis"
)

func renderAppModels(ctx context.RESTContext, app *entity.App, page *amis.Page, arg string) {
	tabs := amis.NewControl()
	tabs.Type = amis.ATTabs
	page.AddBody(tabs)

	tab := amis.NewTab()
	tabs.AddTab(tab)
	tab.Title = "Tables"

	const cols = 2
	var models *amis.Control
	i := -1
	for _, m := range app.Models {
		if m.Builtin() {
			continue
		}

		i++
		if i%cols == 0 {
			models = amis.NewControl()
			models.Type = "grid"
			page.AddBody(models)
		}
		panel := amis.NewControl()
		panel.Type = amis.ATPanel
		titleText := amis.NewControl()
		titleText.Type = "tpl"
		titleText.ClassName = "m-l-sm"
		titleText.Tpl = m.Name
		titleIcon := amis.NewControl()
		titleIcon.Type = "icon"
		titleIcon.Icon = "database"
		panel.Title = []interface{}{titleIcon, titleText}
		panel.ClassName = "Panel--primary"
		ds := m.StoreEngineLabel()

		crud := amis.NewCrud()
		crud.Reset()
		panel.Body = crud
		// data
		items := make([]interface{}, 0, len(m.Slots))
		for _, c := range m.SortedSlots() {
			if c.Deprecated {
				continue
			}

			items = append(items, dto.H{
				"name":     c.Name,
				"kind":     c.KindLabel(),
				"required": c.Required,
				"indexed":  c.Indexed,
			})
		}
		crud.Data = dto.H{"items": items}
		c1 := amis.NewColumn()
		c1.Label = ctx.T("app.model.columns.name.label")
		c1.Name = "name"
		crud.AddColumn(c1)
		c1 = amis.NewColumn()
		c1.Label = "字段类型"
		c1.Name = "kind"
		crud.AddColumn(c1)
		c1 = amis.NewColumn()
		c1.Label = ctx.T("app.model.columns.required.label")
		c1.Name = "required"
		c1.Type = "status"
		crud.AddColumn(c1)
		c1 = amis.NewColumn()
		c1.Label = ctx.T("app.model.columns.indexed.label")
		c1.Name = "indexed"
		c1.Type = "status"
		crud.AddColumn(c1)

		summary := amis.NewTpl(fmt.Sprintf("表类型：%v，存储引擎：%v，ID:%v<br/>总字段：%d个，数据权限：%s",
			m.KindLabel(), ds, m.ID,
			m.TotalColumnsN(),
			m.Feature.Label(),
		))

		panel.Body = []interface{}{summary, amis.NewDivider(), crud}

		action := amis.NewControl()
		panel.AddAction(action)
		action.Type = amis.ATButton
		action.Icon = "fa fa-edit"
		action.Label = "设计"
		action.ActionType = "url"
		action.Link = api.ModelSchemaDesignPage(app.UUID, m.AppID, m.ID)

		action = amis.NewControl()
		panel.AddAction(action)
		action.Type = amis.ATButton
		action.Icon = "fa fa-cog"
		action.Label = "设置"
		action.ActionType = "dialog"

		// model edit dialog
		modelEditDialog := amis.NewDialog()
		modelEditDialog.Title = "数据表设置"
		modelEditDialog.Size = "lg"
		modelEditDialog.CloseOnEsc = true
		action.Dialog = modelEditDialog
		// model edit form
		modelEditForm := amis.NewForm()
		modelEditDialog.Body = modelEditForm
		modelEditForm.API = api.UpdateModelAPI(app.ID, m.ID)
		modelEditForm.Data = dto.H{
			"id":     m.ID,
			"name":   m.Name,
			"remark": m.Remark,
		}
		ctrl := amis.NewControl()
		ctrl.Name = "name"
		ctrl.Type = amis.ATText
		ctrl.Label = ctx.T("app.model.name.label")
		ctrl.Required = true
		modelEditForm.AddControl(ctrl)
		ctrl = amis.NewControl()
		ctrl.Name = "remark"
		ctrl.Type = amis.ATText
		ctrl.Label = ctx.T("app.model.remark.label")
		ctrl.Desc = ctx.T("app.model.remark.desc")
		modelEditForm.AddControl(ctrl)

		// page operations
		combo := amis.NewControl()
		modelEditForm.AddControl(combo)
		combo.Type = amis.ATCombo
		combo.MultiLine = false
		combo.Name = "h"
		combo.Label = ctx.T("app.model.feature.buttons.label")
		combo.Desc = ctx.T("app.model.feature.buttons.desc")
		// feature: create row enabled
		ctrl = amis.NewControl()
		combo.AddControl(ctrl)
		ctrl.Type = amis.ATCheckbox
		ctrl.Label = ctx.T("app.model.feature.button.create.label")
		ctrl.Name = "createRow"
		ctrl.Value = m.Feature.CreateRowEnabled()
		// feature: read row enabled
		ctrl = amis.NewControl()
		ctrl.Name = "readRow"
		ctrl.Type = amis.ATCheckbox
		ctrl.Value = m.Feature.ReadRowEnabled()
		ctrl.Label = ctx.T("app.model.feature.button.read.label")
		combo.AddControl(ctrl)
		// feature: update row enabled
		ctrl = amis.NewControl()
		ctrl.Name = "updateRow"
		ctrl.Type = amis.ATCheckbox
		ctrl.Value = m.Feature.UpdateRowEnabled()
		ctrl.Label = ctx.T("app.model.feature.button.update.label")
		combo.AddControl(ctrl)
		// feature: delete row enabled
		ctrl = amis.NewControl()
		ctrl.Type = amis.ATCheckbox
		ctrl.Label = ctx.T("app.model.feature.button.delete.label")
		ctrl.Name = "deleteRow"
		ctrl.Value = m.Feature.DeleteRowEnabled()
		combo.AddControl(ctrl)

		ctrl = amis.NewControl()
		ctrl.Name = "x"
		ctrl.Type = amis.ATSelect
		ctrl.Label = "权限模型"
		ctrl.Options = []dto.H{
			{
				"label": "ACL",
				"value": "acl",
			},
			{
				"label": "RBAC",
				"value": "rbac",
			},
			{
				"label": "ABAC",
				"value": "abac",
			},
		}
		modelEditForm.AddControl(ctrl)
		ctrl = amis.NewControl()
		ctrl.Name = "x"
		ctrl.Type = amis.ATTextArea
		ctrl.Label = "权限规则"
		modelEditForm.AddControl(ctrl)
		ctrl = amis.NewControl()
		ctrl.Name = "id"
		ctrl.Type = "hidden"
		modelEditForm.AddControl(ctrl)

		if true { // TODO
			action = amis.NewControl()
			panel.AddAction(action)
			action.Type = amis.ATButton
			action.Icon = "fa fa-times text-danger"
			action.Label = "弃用"
			action.ConfirmText = fmt.Sprintf("您确认要弃用数据表 [%s] ?", m.Name)
		}

		models.AddColumn(panel)
	}

	tab = amis.NewTab()
	tabs.AddTab(tab)
	tab.Title = "Settings"

	tab = amis.NewTab()
	tabs.AddTab(tab)
	tab.Title = app.Status.Label()
}

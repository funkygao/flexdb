package design

import (
	"fmt"

	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/pkg/api"
	"github.com/agile-app/flexdb/pkg/dto"
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/view"
	"github.com/agile-app/flexdb/pkg/view/amis"
)

func RenderModel(ctx context.RESTContext, app *entity.App, models []entity.Model, model *entity.Model) *amis.Page {
	page := amis.NewPage()
	titleIcon := amis.NewControl()
	titleIcon.Type = "icon"
	titleIcon.Icon = "database"
	titleIcon.ClassName = "text-info text-lg"
	titleText := amis.NewControl()
	titleText.Type = "tpl"
	titleText.ClassName = "m-l-sm"
	titleText.Tpl = model.Name
	if model.Remark != "" {
		titleRemark := amis.NewControl()
		titleRemark.Type = amis.ATRemark
		titleRemark.Content = model.Remark
		page.Title = []interface{}{titleIcon, titleText, titleRemark}
	} else {
		page.Title = []interface{}{titleIcon, titleText}
	}
	page.SubTitle = fmt.Sprintf("类型：%s，存储引擎：%s，特性：%s",
		model.Kind.Label(), model.StoreEngineLabel(), model.Feature.Label())

	// toolbar
	toolbar := amis.NewToolbar()
	page.AddToolbar(toolbar)
	toolbar.Type = amis.ATButton
	toolbar.ActionType = amis.AATUrl
	toolbar.Link = api.RowCRUDPage(app.UUID, model.AppID, model.ID)
	toolbar.Label = "预览"
	toolbar.Tooltip = "数据预览"
	toolbar.Icon = "fa fa-eye pull-left"

	toolbar = amis.NewToolbar()
	page.AddToolbar(toolbar)
	toolbar.Type = amis.ATButton
	toolbar.ActionType = amis.AATDrawer
	toolbar.Label = "添加字段"
	toolbar.Icon = "fa fa-plus pull-left"
	toolbar.Primary = true

	// toolbar.drawer
	toolbar.Drawer = amis.NewDrawer()
	//toolbar.Drawer.Size = "lg"
	toolbar.Drawer.Title = fmt.Sprintf("%s 添加字段", model.Name)
	toolbar.Drawer.CloseOnEsc = true

	// toolbar.drawer.form
	addColumnForm := amis.NewForm()
	addColumnForm.Autofocus = true
	addColumnForm.Reload = "window"
	addColumnForm.API = api.AddColumnAPI(model.AppID, model.ID)
	toolbar.Drawer.Body = addColumnForm
	renderAddColumnForm(ctx, addColumnForm, models, model)

	// aside nav
	aside := page.SetAside("wrapper")
	aside.Size = "xs"
	nav := amis.NewNav()
	nav.Stacked = true
	aside.AddBody(nav)
	virtualModelsNav := amis.NewNav()
	virtualModelsNav.Stacked = true

	for _, m := range models {
		link := amis.NewLink(m.Name, api.ModelSchemaDesignPage(app.UUID, app.ID, m.ID))
		if m.Virtual() {
			virtualModelsNav.AddLink(*link)
		} else {
			nav.AddLink(*link)
		}
	}

	if len(virtualModelsNav.Links) > 0 {
		if len(nav.Links) > 0 {
			aside.AddBody(amis.NewDivider())
		}

		aside.AddBody(virtualModelsNav)
	}

	// crud
	crud := amis.NewCrud()
	crud.Reset()
	page.AddBody(crud)
	crud.SaveOrderAPI = api.SaveColumnsOrderAPI(model.AppID, model.ID)
	crud.DefaultParams = dto.H{
		"perPage": 100,
	}

	// crud.column.op
	op := amis.NewColumn()
	op.Type = amis.ATOperation
	op.Label = "操作"
	op.Fixed = "right" // fixed, so that user can see it without scrolling
	op.Toggled = true

	// crud.column.op.edit
	editBtn := amis.NewButton()
	op.AddButton(editBtn)
	editBtn.Type = amis.ATButton
	editBtn.Tooltip = "修改"
	editBtn.VisibleOn = "this.builtin == false"
	editBtn.ActionType = amis.AATDrawer
	editBtn.Icon = "fa fa-pencil"

	editDrawer := amis.NewDrawer()
	editDrawer.Position = "right"
	editDrawer.Title = "修改"
	editDrawer.CloseOnEsc = true
	editBtn.Drawer = editDrawer
	editForm := amis.NewForm()
	editForm.Reload = "window"
	editForm.API = api.UpdateColumnAPI(model.AppID, model.ID)
	editBtn.Drawer.Body = editForm

	// TODO delete should reload window

	// editForm.control
	ctrl := amis.NewControl()
	ctrl.Type = "text"
	ctrl.Name = "name"
	ctrl.Label = ctx.T("app.model.columns.name.label")
	ctrl.Required = true
	editForm.AddControl(ctrl)
	ctrl = amis.NewControl()
	ctrl.Type = "text"
	ctrl.Name = "label"
	ctrl.Label = ctx.T("app.model.columns.label.label")
	editForm.AddControl(ctrl)
	ctrl = amis.NewControl()
	ctrl.Type = "text"
	ctrl.Name = "remark"
	ctrl.Label = ctx.T("app.model.columns.remark.label")
	editForm.AddControl(ctrl)
	ctrl = amis.NewControl()
	ctrl.Type = "textarea"
	ctrl.Name = "choices"
	ctrl.Label = "可选项"
	ctrl.Desc = "可选项间通过逗号分割"
	ctrl.VisibleOn = fmt.Sprintf("this.rawkind == %v", entity.ColumnChoice)
	editForm.AddControl(ctrl)
	ctrl = amis.NewControl()
	ctrl.Type = amis.ATCheckbox
	ctrl.Name = "required"
	ctrl.Label = ctx.T("app.model.columns.required.label")
	editForm.AddControl(ctrl)
	ctrl = amis.NewControl()
	ctrl.Type = amis.ATCheckbox
	ctrl.Name = "ro"
	ctrl.Label = ctx.T("app.model.columns.ro.label")
	ctrl.Desc = ctx.T("app.model.columns.ro.desc")
	editForm.AddControl(ctrl)

	// crud.column.op.del
	delBtn := amis.NewButton()
	op.AddButton(delBtn)
	delBtn.Icon = "fa fa-times text-danger"
	delBtn.ActionType = "ajax"
	delBtn.Tooltip = "弃用"
	delBtn.VisibleOn = "this.builtin == false"
	delBtn.API = api.DeprecateColumnAPI(model.AppID, model.ID)
	delBtn.ConfirmText = "您确认要弃用 <b>$name</b> 字段 ?"

	// data
	items := make([]interface{}, 0, model.TotalColumnsN())
	for i, c := range append(model.SortedSlots(), model.BuiltinColumns()...) {
		if c.Deprecated {
			continue
		}

		items = append(items, dto.H{
			"id":         c.ID,
			"name":       c.Name,
			"remark":     c.Remark,
			"kind":       c.KindLabel(),
			"rawkind":    c.Kind,
			"label":      c.Label,
			"required":   c.Required,
			"slot":       c.Slot,
			"indexed":    c.Indexed,
			"builtin":    i >= len(model.Slots),
			"sortable":   c.Sortable,
			"choices":    c.Choices,
			"ro":         c.ReadOnly,
			"uniq":       c.Unique,
			"cuser":      c.CUser,
			"ctime":      c.CTime,
			"muser":      c.MUser,
			"mtime":      c.MTime,
			"relational": c.Relational(),
		})
	}
	crud.Data = dto.H{"items": items}
	columns := []map[string]string{
		{"label": "S", "name": "slot", "remark": "系统内部调试值，可忽略"},
		{"label": "字段名称", "name": "name", "remark": "支持中文，但不能包含特殊符号和空格"},
		{"label": "显示名称", "name": "label", "remark": "页面表单上该字段的输入提示"},
		{"label": "字段类型", "name": "kind"},
		{"label": "字段备注", "name": "remark", "remark": "在输入表单里，会给出提示"},
		{"label": "预留", "name": "builtin", "remark": "所有表都会有的系统提供的字段", "type": amis.ATStatus},
		{"label": "必填", "name": "required", "type": amis.ATStatus},
		{"label": "可搜", "name": "indexed", "type": amis.ATStatus},
		{"label": "可排序", "name": "sortable", "type": amis.ATStatus},
		{"label": "只读", "name": "ro", "type": amis.ATStatus, "remark": "录入后不可修改"},
		{"label": "唯一", "name": "uniq", "type": amis.ATStatus},
		{"label": "可搜", "name": "indexed", "type": amis.ATStatus},
		{"label": ctx.T("column.builtin.cuser.label"), "name": "cuser", "off": ""},
		{"label": ctx.T("column.builtin.ctime.label"), "name": "ctime", "type": amis.ATDatetime, "off": ""},
		{"label": ctx.T("column.builtin.muser.label"), "name": "muser", "off": ""},
		{"label": ctx.T("column.builtin.mtime.label"), "name": "mtime", "type": amis.ATDatetime, "off": ""},
	}

	for _, c := range columns {
		col := amis.NewColumn()
		col.Label = c["label"]
		col.Name = c["name"]
		if r, present := c["remark"]; present {
			col.Remark = r
		}
		if t, present := c["type"]; present {
			col.Type = t
		}
		if _, off := c["off"]; off {
			col.Toggled = false
		}
		crud.AddColumn(col)
	}

	// add op column at the end
	crud.AddColumn(op)

	return page
}

func renderAddColumnForm(ctx context.RESTContext, addColumnForm *amis.Form, models []entity.Model, model *entity.Model) {
	ctrl := amis.NewControl()
	ctrl.Name = "name"
	ctrl.Type = "text"
	ctrl.Label = ctx.T("app.model.columns.name.label")
	ctrl.Required = true
	addColumnForm.AddControl(ctrl)

	ctrl = amis.NewControl()
	ctrl.Name = "label"
	ctrl.Label = ctx.T("app.model.columns.label.label")
	ctrl.Desc = ctx.T("app.model.columns.label.desc")
	ctrl.Type = "text"
	addColumnForm.AddControl(ctrl)

	ctrl = amis.NewControl()
	ctrl.Name = "kind"
	ctrl.Label = "字段类型"
	ctrl.Type = "select"
	ctrl.Required = true
	ctrl.Inline = true
	kinds := make([]dto.H, 0, len(view.ExcelColumnKindMap))
	for k, v := range view.ExcelColumnKindMap {
		kinds = append(kinds, dto.H{
			"label": k,
			"value": v,
		})
	}
	ctrl.Options = kinds
	addColumnForm.AddControl(ctrl)

	ctrl = amis.NewControl()
	ctrl.Type = amis.ATTreeSelect
	ctrl.Name = "h.ref"
	ctrl.Label = "目标字段"
	ctrl.VisibleOn = fmt.Sprintf("this.kind == %v || this.kind == %v", entity.ColumnLookup, entity.ColumnMany2One)
	refModels := []dto.H{}
	// render all model.slots that the column can reference
	for _, m := range models {
		if m.ID == model.ID {
			// self excluded
			continue
		}

		m1 := dto.H{}
		m1["label"] = m.Name
		cols := []dto.H{}
		for _, c := range m.SortedSlots() {
			if c.Referable() {
				cols = append(cols, dto.H{
					"label": c.Name,
					"value": fmt.Sprintf("%d,%d", m.ID, c.Slot),
				})
			}
		}
		m1["children"] = cols
		refModels = append(refModels, m1)
	}
	ctrl.Options = refModels
	addColumnForm.AddControl(ctrl)

	ctrl = amis.NewControl()
	ctrl.Name = "remark"
	ctrl.Label = ctx.T("app.model.columns.remark.label")
	ctrl.Type = "text"
	addColumnForm.AddControl(ctrl)

	ctrl = amis.NewControl()
	ctrl.Type = "textarea"
	ctrl.Name = "choices"
	ctrl.Label = "可选项"
	ctrl.Desc = "通过逗号分割可选项，注意不要使用全角逗号"
	ctrl.VisibleOn = fmt.Sprintf("this.kind == %v", entity.ColumnChoice)
	addColumnForm.AddControl(ctrl)

	ctrl = amis.NewControl()
	ctrl.Name = "dftval"
	ctrl.Label = "默认值"
	ctrl.Type = "text"
	addColumnForm.AddControl(ctrl)

	ctrl = amis.NewControl()
	ctrl.Name = "rule"
	ctrl.Label = "业务规则"
	ctrl.Type = "text"
	ctrl.Desc = "具体使用目前需要联系研发，后期可以自助"
	addColumnForm.AddControl(ctrl)

	ctrl = amis.NewControl()
	ctrl.Type = amis.ATCheckbox
	ctrl.Label = ctx.T("app.model.columns.required.label")
	ctrl.Name = "required"
	addColumnForm.AddControl(ctrl)

	ctrl = amis.NewControl()
	ctrl.Type = amis.ATCheckbox
	ctrl.Label = ctx.T("app.model.columns.indexed.label")
	ctrl.Desc = ctx.T("app.model.columns.indexed.desc")
	ctrl.Name = "indexed"
	addColumnForm.AddControl(ctrl)

	ctrl = amis.NewControl()
	ctrl.Type = amis.ATCheckbox
	ctrl.Label = ctx.T("app.model.columns.sortable.label")
	ctrl.Name = "sortable"
	addColumnForm.AddControl(ctrl)

	ctrl = amis.NewControl()
	ctrl.Label = ctx.T("app.model.columns.unique.label")
	ctrl.Type = amis.ATCheckbox
	ctrl.Name = "uniq"
	addColumnForm.AddControl(ctrl)

	ctrl = amis.NewControl()
	ctrl.Label = ctx.T("app.model.columns.unique.label")
	ctrl.Name = "ro"
	ctrl.Type = amis.ATCheckbox
	ctrl.Desc = "录入后不可修改"
	addColumnForm.AddControl(ctrl)
}

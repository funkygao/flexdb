package view

import (
	"strings"

	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/pkg/api"
	"github.com/agile-app/flexdb/pkg/dto"
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/view/amis"
	"github.com/agile-app/flexdb/plugins/filter"
)

// CRUD creates a amis crud control for a model.
func CRUD(ctx context.RESTContext, model *entity.Model, createForm *amis.Form) *amis.Crud {
	crud := amis.NewCrud()
	crud.FilterTogglable = false
	crud.Mode = "table"
	crud.Draggable = false
	crud.API = api.FindRowsAPI(model.ID) // TODO merge into data
	crud.QuickSaveAPI = api.RowQuickSaveAPI(model.ID)
	crud.DefaultParams = dto.H{
		"perPage": 20,
	}

	var (
		op                   *amis.Column
		editForm             *amis.Form
		singleRowInputerMode = model.SingleRowKind() && !filter.SatisfyToBeKilledPermRule(ctx.PIN())
	)

	if !singleRowInputerMode && model.Feature.DeleteRowEnabled() {
		bulkAction := amis.NewControl()
		bulkAction.Label = ctx.T("crud.bulkDel")
		bulkAction.Type = amis.ATButton
		bulkAction.Level = "danger"
		bulkAction.ActionType = amis.AAAjax
		bulkAction.ConfirmText = "您确认要批量删除记录 ? <br/>WIP"
		bulkAction.API = "delete http://localhost:8000/excel/?ids=${ids|raw}"
		crud.AddBulkAction(bulkAction)
	}

	// crud.column.op
	op = amis.NewColumn()
	op.Type = amis.ATOperation
	op.Label = ctx.T("crud.op.label")
	op.Fixed = "right" // fixed, so that user can see it without scrolling
	op.Toggled = true

	// crud.column.op.edit
	if model.Feature.UpdateRowEnabled() {
		editBtn := amis.NewButton()
		op.AddButton(editBtn)
		editBtn.Type = "button"
		editBtn.Tooltip = ctx.T("crud.op.edit")
		editBtn.ActionType = "drawer"
		editBtn.Icon = "fa fa-pencil"

		editDrawer := amis.NewDrawer()
		editDrawer.Position = "right"
		editDrawer.Title = "修改数据"
		editDrawer.Size = "lg"
		editDrawer.CloseOnEsc = true
		editBtn.Drawer = editDrawer
		editForm = amis.NewForm()
		editForm.SubmitText = "修改"
		editForm.InitApi = api.EditRowInitAPI(model.ID)
		editForm.API = api.UpdateRowAPI(model.ID)
		editBtn.Drawer.Body = editForm
	}

	if model.Feature.ChangeAuditEnabled() {
		logBtn := amis.NewButton()
		op.AddButton(logBtn)
		logBtn.Type = "button"
		logBtn.Tooltip = ctx.T("crud.op.log")
		logBtn.ActionType = "drawer"
		logBtn.Icon = "fa fa-flag"
	}

	// crud.column.op.del
	if model.Feature.DeleteRowEnabled() {
		delBtn := amis.NewButton()
		op.AddButton(delBtn)
		delBtn.Icon = "fa fa-times text-danger"
		delBtn.ActionType = "ajax"
		delBtn.Tooltip = ctx.T("crud.op.delete")
		delBtn.API = api.DeleteRowAPI(model.ID)
		delBtn.ConfirmText = "您确认要删除 ID为 $id 的记录 ?"
	}

	// ID column is required in amis
	idCol := amis.NewColumn()
	crud.AddColumn(idCol)
	if model.IdentName != "" {
		idCol.Name = model.IdentName
		idCol.Label = model.IdentName
		idCol.Copyable = dto.H{
			"content": "$" + model.IdentName,
		}
	} else {
		idCol.Name = "id"
		idCol.Label = "ID"
	}
	idCol.Copyable = dto.H{
		"content": "$id",
	}

	idCol.Remark = ctx.T("crud.id.remark")
	idCol.Toggled = true

	if !singleRowInputerMode {
		// filter
		crud.Filter = amis.NewFilter()
		crud.FilterTogglable = true

		// filter title
		titleIcon := amis.NewControl()
		titleIcon.Type = "icon"
		titleIcon.Icon = "search"
		titleIcon.ClassName = "text-info text-lg"
		titleText := amis.NewControl()
		titleText.Type = "tpl"
		titleText.ClassName = "m-l-sm"
		titleText.Tpl = ctx.T("crud.filter.title")
		titleRemark := amis.NewControl()
		titleRemark.Type = amis.ATRemark
		titleRemark.Content = ctx.T("crud.filter.remark")
		crud.Filter.Title = []interface{}{titleIcon, titleText, titleRemark}

		// filter actions
		action := amis.NewControl()
		action.Type = "reset"
		action.Label = ctx.T("crud.filter.action.reset")
		crud.Filter.AddAction(action)
		action = amis.NewControl()
		action.Type = "submit"
		action.Label = ctx.T("crud.filter.action.submit")
		action.Level = "primary"
		crud.Filter.AddAction(action)

		// id filter
		idFilter := amis.NewControl()
		idFilter.Name = api.RowID
		idFilter.Placeholder = ctx.T("crud.filter.id.placeholder")
		idFilter.Type = "text"
		idFilter.Validations = amis.VTInteger
		crud.Filter.AddControl(idFilter)
	} else {
		// current user can only enter a single row
		crud.HeaderToolbar = crud.HeaderToolbar[:0]
		crud.FooterToolbar = crud.FooterToolbar[:0]
	}

	// slots
	for i, c := range model.SortedSlots() {
		if c.Deprecated {
			continue
		}

		slot := ColumnWidgetOf(c)
		// crud.filter.control
		if crud.Filter != nil && slot.Indexed {
			c := amis.NewControl()
			c.Name = slot.Name
			c.Placeholder = slot.Name
			crud.Filter.AddControl(c)
		}

		// crud.column
		if slot.ListViewable() {
			col := amis.NewColumn()
			col.Name = slot.Name
			col.Type = slot.ListWidgetType()
			col.Label = slot.RenderLabel()
			href, anchor := slot.ListViewHrefIfNec()
			if href != nil {
				col.Href = href
				col.Type = amis.AATLink
				if anchor != nil {
					col.Body = anchor
				}
			}
			slot.EnrichViewWidget(col)
			if slot.Remark != "" {
				col.Remark = slot.Remark
			}
			if slot.Sortable && !singleRowInputerMode {
				col.Sortable = true
			}
			if slot.Indexed && !singleRowInputerMode {
				col.Searchable = true
			}
			if !slot.ReadOnly && !singleRowInputerMode {
				//col.QuickEdit = true
			}
			if i > 8 {
				col.Toggled = false
			}
			crud.AddColumn(col)
		}

		if !slot.Virtual() {
			// createForm.control
			if createForm != nil {
				ctrl := amis.NewControl()
				ctrl.Type = slot.EditWidgetType()
				ctrl.Label = slot.RenderLabel()
				ctrl.Name = slot.Name
				if slot.Remark != "" {
					ctrl.Desc = slot.Remark
				}
				if model.OnlyPinCanInput(slot.Column) {
					ctrl.Value = ctx.PIN()
					//ctrl.Disabled = true
				}
				slot.EnrichEditWidget(ctrl)
				if slot.Required {
					ctrl.Required = true
				}
				createForm.AddControl(ctrl)
			}

			// editForm.control TODO can merge with createForm control: needn't create again
			if model.Feature.UpdateRowEnabled() {
				ctrl := amis.NewControl()
				ctrl.Type = slot.EditWidgetType()
				ctrl.Name = slot.Name
				ctrl.Label = slot.RenderLabel()
				slot.EnrichEditWidget(ctrl)
				if slot.Remark != "" {
					ctrl.Desc = slot.Remark
				}
				if slot.Required {
					ctrl.Required = true
				}
				if slot.ReadOnly || model.OnlyPinCanInput(slot.Column) {
					ctrl.Disabled = true
				}
				editForm.AddControl(ctrl)
			}
		}
	}

	if model.Feature.UpdateRowEnabled() {
		for _, c := range model.BuiltinColumns() {
			ctrl := amis.NewControl()
			slot := ColumnWidgetOf(c)
			ctrl.Name = strings.ToLower(slot.Name)
			ctrl.Label = c.Label
			if !c.ReadOnly {
				if createForm != nil {
					ctrl.Type = slot.EditWidgetType()
					createForm.AddControl(ctrl)
				}
			} else {
				ctrl.Type = "static"
			}
			editForm.AddControl(ctrl)
		}
	}

	// builtin columns
	if !singleRowInputerMode && !model.Virtual() {
		// ctime
		builtinCol := amis.NewColumn()
		builtinCol.Name = dto.KeyCTime
		builtinCol.Label = ctx.T("column.builtin.ctime.label")
		builtinCol.Toggled = false
		crud.AddColumn(builtinCol)
		// cuser
		builtinCol = amis.NewColumn()
		builtinCol.Name = dto.KeyCUser
		builtinCol.Label = ctx.T("column.builtin.cuser.label")
		builtinCol.Toggled = false
		crud.AddColumn(builtinCol)
		// mtime
		builtinCol = amis.NewColumn()
		builtinCol.Name = dto.KeyMTime
		builtinCol.Label = ctx.T("column.builtin.mtime.label")
		if !singleRowInputerMode {
			builtinCol.Sortable = true
		}
		crud.AddColumn(builtinCol)
		builtinCol = amis.NewColumn()
		builtinCol.Name = dto.KeyMUser
		builtinCol.Label = ctx.T("column.builtin.muser.label")
		builtinCol.Toggled = false
		crud.AddColumn(builtinCol)
	}

	// op column is always the last fixed column
	if model.Feature.UpdateRowEnabled() || model.Feature.DeleteRowEnabled() {
		crud.AddColumn(op)
	}

	return crud
}

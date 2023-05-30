package design

import (
	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/internal/spec"
	"github.com/agile-app/flexdb/pkg/api"
	"github.com/agile-app/flexdb/pkg/dto"
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/view/amis"
)

func renderAppShare(ctx context.RESTContext, app *entity.App, page *amis.Page, appID string) {
	shareBtn := amis.NewButton()
	page.AddBody(shareBtn)
	shareBtn.ActionType = amis.ATDialog
	shareBtn.Label = ctx.T("design.app.menu.share")
	shareBtn.Icon = "fa fa-plus pull-left"
	shareBtn.ClassName = "m-b-sm"

	shareDialog := amis.NewDialog()
	shareBtn.Dialog = shareDialog
	shareDialog.Title = ctx.T("design.app.menu.share")
	shareDialog.CloseOnEsc = true
	shareDialog.Size = "lg"

	createForm := amis.NewForm()
	createForm.Autofocus = true
	createForm.API = api.ShareAppToUser(app.ID)
	shareDialog.Body = createForm

	ctrl := amis.NewControl()
	createForm.AddControl(ctrl)
	ctrl.Type = amis.ATTransfer
	ctrl.Inline = false
	ctrl.Searchable = true
	ctrl.SearchAPI = api.AutocompleteShareRecommend(app.ID)
	ctrl.Name = "shareTo"
	ctrl.Label = ctx.T("design.app.share.column.subject.label")

	crud := amis.NewCrud()
	crud.Reset()
	page.AddBody(crud)
	// data
	items := make([]interface{}, 0, len(app.Shares))
	for _, s := range app.Shares {
		items = append(items, dto.H{
			"id":      s.AppID,
			"subject": s.Subject,
			"ctime":   s.CTime.Format(spec.SimpleDateFormat),
		})
	}
	crud.Data = dto.H{"items": items}

	columns := []*dto.H{
		{"label": ctx.T("design.app.share.column.subject.label"), "name": "subject"},
		{"label": ctx.T("design.app.share.column.ctime.label"), "name": "ctime"},
	}
	for _, c := range columns {
		col := amis.NewColumn()
		col.Label = c.S("label")
		col.Name = c.S("name")
		crud.AddColumn(col)
	}

	op := amis.NewColumn()
	crud.AddColumn(op)
	op.Type = amis.ATOperation
	op.Label = ctx.T("crud.op.label")
	op.Fixed = "right" // fixed, so that user can see it without scrolling
	op.Toggled = true

	revokeBtn := amis.NewButton()
	op.AddButton(revokeBtn)
	revokeBtn.Icon = "fa fa-times text-danger"
	revokeBtn.ActionType = "ajax"
	revokeBtn.Tooltip = ctx.T("crud.op.delete")
	revokeBtn.ConfirmText = "您确认要停止共享 ?"

}

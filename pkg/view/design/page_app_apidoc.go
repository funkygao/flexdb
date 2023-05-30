package design

import (
	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/view/amis"
)

func renderAppAPIDocs(ctx context.RESTContext, app *entity.App, page *amis.Page, arg string) {
	tpl := amis.NewControl()
	tpl.Type = "tpl"
	tpl.Tpl = arg
	page.AddBody(tpl)
}

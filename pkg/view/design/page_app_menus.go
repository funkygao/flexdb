package design

import (
	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/view/amis"
)

var (
	appMenuRenderers = map[string]func(context.RESTContext, *entity.App, *amis.Page, string){
		"models": renderAppModels,
		"api":    renderAppAPIDocs,
		"perm":   renderAppPerm,
		"share":  renderAppShare,
	}
)

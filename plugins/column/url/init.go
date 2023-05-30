package url

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnURL,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &url{Column: c, ctx: ctx}
		})
}

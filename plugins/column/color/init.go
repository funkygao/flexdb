package color

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnColor,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &color{Column: c, ctx: ctx}
		})
}

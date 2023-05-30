package boolean

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnBoolean,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &boolean{Column: c, ctx: ctx}
		})
}

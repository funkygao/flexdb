package city

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnCity,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &city{Column: c, ctx: ctx}
		})
}

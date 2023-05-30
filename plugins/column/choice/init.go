package choice

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnChoice,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &choice{Column: c, ctx: ctx}
		})
}

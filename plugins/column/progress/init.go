package progress

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnProgress,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &progress{Column: c, ctx: ctx}
		})
}

package text

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnText,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &text{Column: c, ctx: ctx}
		})
}

package sequence

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnAutoNumber,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &sequence{Column: c, ctx: ctx}
		})
}

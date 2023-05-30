package integer

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnInteger,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &integer{Column: c, ctx: ctx}
		})
}

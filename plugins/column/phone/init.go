package phone

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnPhone,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &phone{Column: c, ctx: ctx}
		})
}

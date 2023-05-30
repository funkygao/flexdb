package password

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnPassword,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &password{Column: c, ctx: ctx}
		})
}

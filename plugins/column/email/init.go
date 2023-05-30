package email

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnEmail,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &email{Column: c, ctx: ctx}
		})
}

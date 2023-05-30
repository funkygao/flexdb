package datetime

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnDatetime,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &datetime{Column: c, ctx: ctx}
		})
}

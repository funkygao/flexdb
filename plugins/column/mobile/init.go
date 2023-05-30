package mobile

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnMobile,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &mobile{Column: c, ctx: ctx}
		})
}

package image

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnImage,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &image{Column: c, ctx: ctx}
		})
}

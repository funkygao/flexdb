package file

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnFile,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &file{Column: c, ctx: ctx}
		})
}

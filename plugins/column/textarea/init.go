package textarea

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnTextArea,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &textarea{Column: c, ctx: ctx}
		})
}

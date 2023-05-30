package formula

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnFormula,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &formula{Column: c, ctx: ctx}
		})
}

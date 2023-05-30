package erp

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnERP,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &erp{Column: c, ctx: ctx}
		})
}

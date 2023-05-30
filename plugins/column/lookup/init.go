package lookup

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnLookup,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &lookup{Column: c, ctx: ctx}
		},
		reservedColumn)
}

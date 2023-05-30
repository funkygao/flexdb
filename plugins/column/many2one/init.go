package many2one

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnMany2One,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &many2one{Column: c, ctx: ctx}
		},
		reservedColumn)
}

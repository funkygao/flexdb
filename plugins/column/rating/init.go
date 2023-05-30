package rating

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnRating,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &rating{Column: c, ctx: ctx}
		})
}

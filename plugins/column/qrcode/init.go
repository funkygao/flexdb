package qrcode

import "github.com/agile-app/flexdb/pkg/entity"

func init() {
	entity.RegisterColumnPlugin(
		entity.ColumnQRCode,
		func(c *entity.Column, ctx entity.PluginContext) entity.ColumnPlugin {
			return &qrcode{Column: c, ctx: ctx}
		})
}

package lookup

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/agile-app/flexdb/internal/profile"
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/view/amis"
	"github.com/gin-gonic/gin"
)

// plugin compliance check
var (
	_ entity.ColumnPlugin       = (*lookup)(nil)
	_ entity.Indexable          = (*lookup)(nil)
	_ entity.ReferenceValidator = (*lookup)(nil)
	_ entity.EditWidgetAware    = (*lookup)(nil)
	_ entity.ViewWidgetAware    = (*lookup)(nil)
)

type lookup struct {
	*entity.Column
	ctx entity.PluginContext
}

func (c *lookup) ValidateReference(m *entity.Model, toAddColumns []*entity.Column) error {
	return nil // TODO
}

func (c *lookup) ValidateCell(val string) error {
	refRowID, err := c.EvaluateCell(val, nil)
	if err != nil {
		return err
	}

	if id := refRowID.(uint64); id < 1 {
		return fmt.Errorf("empty value on column:%s", c.Name)
	}

	// TODO lookup ref table row to see whether it exists

	return nil
}

func (lookup) EvaluateCell(val string, row *entity.Row) (interface{}, error) {
	return strconv.ParseUint(val, 10, 64)
}

func (c *lookup) CreateIndex(val string) (entity.Index, error) {
	if !c.Indexed {
		return nil, nil
	}

	idx := &entity.IndexStr{Slot: c.Slot}
	idx.Val = val
	return idx, nil
}

func (lookup) IndexKind() entity.Index {
	return &entity.IndexInt{}
}

func (c lookup) EnrichEditControl(ctrl *amis.Control) {
	ctrl.Searchable = true
	ctrl.Source = fmt.Sprintf("%s/api/v0.1/data/lookup/Model/%d/Slot/%d",
		profile.P.APIBaseEndpoint,
		c.RefModelID, c.RefSlot)

}

func (c lookup) EnrichViewControl(col *amis.Column) {
	col.Remark = fmt.Sprintf("关联表：%d %d", c.RefModelID, c.RefSlot)
	col.PopOver = gin.H{
		"mode":  "popOver",
		"type":  "panel",
		"title": "查看详情",
		"body": []gin.H{
			{
				"type": "form",
				"initApi": fmt.Sprintf("%s/api/v0.1/data//Model/%d/Row/$%s",
					profile.P.APIBaseEndpoint,
					c.RefModelID, url.QueryEscape(c.Name)),
				"controls": []gin.H{
					{
						"label": c.Name,
						"type":  "static",
						"name":  "商家名称", // TODO
					},
				},
			},
		},
	}
}

package choice

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/agile-app/flexdb/internal/profile"
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/view/amis"
	"github.com/gin-gonic/gin"
)

// plugin compliance check
var (
	_ entity.ColumnPlugin    = (*choice)(nil)
	_ entity.Indexable       = (*choice)(nil)
	_ entity.Introspector    = (*choice)(nil)
	_ entity.EditWidgetAware = (*choice)(nil)
)

type choice struct {
	*entity.Column
	ctx entity.PluginContext
}

func (c *choice) ValidateCell(val string) error {
	if c.Choices != "" {
		for _, choice := range c.ChoiceOptions() {
			if val == choice {
				// bingo!
				return nil
			}
		}

		return fmt.Errorf("column:%s has invalid choice: %s", c.Name, val)
	}

	valueID, err := c.EvaluateCell(val, nil)
	if err != nil {
		return err
	}

	m, err := c.PluginContext.ModelAccessorOf(c.ModelID)
	if err != nil {
		return err
	}

	pickList, err := m.SlotPickList(c.Slot)
	if err != nil {
		return err
	}

	for _, item := range pickList {
		if item.ID == valueID {
			// bingo!
			return nil
		}
	}

	return fmt.Errorf("column:%s has invalid choice: %d", c.Name, valueID)
}

func (c choice) EvaluateCell(val string, row *entity.Row) (interface{}, error) {
	if c.Choices != "" {
		return val, nil
	}

	return strconv.ParseInt(val, 10, 64)
}

func (c *choice) Introspect() error {
	if strings.Contains(c.Choices, "，") {
		return errors.New("可选项里包含了全角逗号，请换成半角逗号分隔选项")
	}

	return nil
}

func (c *choice) CreateIndex(val string) (entity.Index, error) {
	if !c.Indexed {
		return nil, nil
	}

	if c.Choices != "" {
		return &entity.IndexStr{Slot: c.Slot, Val: val}, nil
	}

	// picklist choice column always stores ID as value
	idx := &entity.IndexInt{Slot: c.Slot}
	v, err := c.EvaluateCell(val, nil)
	if err != nil {
		return nil, err
	}

	idx.Val = v.(int64)
	return idx, nil
}

func (c choice) IndexKind() entity.Index {
	if c.Choices != "" {
		return &entity.IndexStr{}
	}

	return &entity.IndexInt{}
}

func (c choice) EnrichEditControl(ctrl *amis.Control) {
	updatable := true
	if c.Choices != "" {
		updatable = false

		if c.Required {
			options := c.ChoiceOptions()
			if len(options) == 1 {
				// auto choose
				ctrl.Value = options[0]
			}
		}
	}
	ctrl.Creatable = updatable
	ctrl.Editable = updatable
	ctrl.Searchable = true

	ctrl.AddControls = []gin.H{
		{
			"type":  "text",
			"name":  "label",
			"label": "选项标题",
		},
		{
			"type":  "text",
			"name":  "value",
			"label": "选项值",
		},
	}
	if updatable {
		ctrl.AddAPI = map[string]string{
			"url": fmt.Sprintf("%s/api/v0.1/data/Model/%d/Slot/%d/Picklist",
				profile.P.APIBaseEndpoint,
				c.ModelID, c.Slot),
			"method": "post",
		}
		ctrl.EditAPI = map[string]string{
			"url": fmt.Sprintf("%s/api/v0.1/data/Model/%d/Slot/%d/Picklist?id=$val",
				profile.P.APIBaseEndpoint,
				c.ModelID, c.Slot),
			"method": "put",
		}
	}
	ctrl.Source = fmt.Sprintf("%s/api/v0.1/data/Model/%d/Slot/%d/Picklist",
		profile.P.APIBaseEndpoint,
		c.ModelID, c.Slot)
}

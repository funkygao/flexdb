package data

import (
	"github.com/agile-app/flexdb/internal/controller/context"
	"github.com/agile-app/flexdb/pkg/dto"
	"github.com/agile-app/flexdb/pkg/store"
	"github.com/gin-gonic/gin"
)

func (dc *dataHandler) CreatePickItem(c context.RESTContext) {
	ms, err := store.Provider.InferAppStore(c).LoadModel(c.ModelID())
	if err != nil {
		c.AbortWithError(err)
		return
	}

	var lv dto.LabelValue
	if err := c.Gin().ShouldBindJSON(&lv); err != nil {
		c.AbortWithError(err)
		return
	}

	item, err := ms.CreateSlotPickItem(c.SlotID(), lv.Label)
	if err != nil {
		c.AbortWithError(err)
		return
	}

	c.RenderOK(item.ID)
}

func (dc *dataHandler) UpdatePickItem(c context.RESTContext) {
	ms, err := store.Provider.InferAppStore(c).LoadModel(c.ModelID())
	if err != nil {
		c.AbortWithError(err)
		return
	}

	var lv dto.LabelValue
	if err := c.Gin().ShouldBindJSON(&lv); err != nil {
		c.AbortWithError(err)
		return
	}

	if err := ms.UpdateSlotPickItem(c.QueryID(), lv.Label); err != nil {
		c.AbortWithError(err)
		return
	}

	c.RenderOKWithoutData()
}

func (dc *dataHandler) ShowPicklist(c context.RESTContext) {
	ms, err := store.Provider.InferAppStore(c).LoadModel(c.ModelID(), "Slots")
	if err != nil {
		c.AbortWithError(err)
		return
	}

	slot := ms.EntityModel().SlotBySlotID(c.SlotID())
	if slot != nil {
		if options := slot.ChoiceOptions(); options != nil {
			// extract choices from slot instead of picklist
			c.RenderOK(gin.H{"options": options})
			return
		}
	}

	picklist, err := ms.SlotPickList(c.SlotID())
	if err != nil {
		c.AbortWithError(err)
		return
	}

	c.RenderOK(gin.H{"options": picklist})
}

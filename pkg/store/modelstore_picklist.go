package store

import (
	"time"

	"github.com/agile-app/flexdb/pkg/entity"
	"gorm.io/gorm/clause"
)

func (m *modelStore) CreateSlotPickItem(slot int16, val string) (item *entity.PickItem, err error) {
	item = &entity.PickItem{ModelID: m.ID, Slot: slot, Val: val, CUser: m.c.PIN()}
	now := time.Now()
	item.CTime = now
	item.MTime = now
	item.CUser = m.c.PIN()
	err = m.store.mdb().Omit(clause.Associations).Create(item).Error
	return
}

func (m *modelStore) UpdateSlotPickItem(id int64, val string) (err error) {
	item := entity.PickItem{ID: id}
	updates := map[string]interface{}{
		"val":   val,
		"muser": m.c.PIN(),
		"mtime": time.Now(),
	}
	return m.store.mdb().Model(&item).Updates(updates).Error
}

func (m *modelStore) SlotPickList(slot int16) (picklist []*entity.PickItem, err error) {
	err = m.store.mdb().Order("id asc").Where("model_id=? and slot=?", m.ID, slot).Find(&picklist).Error
	return
}

package store

import (
	"github.com/agile-app/flexdb/internal/cache"
	"github.com/agile-app/flexdb/internal/profile"
	"github.com/agile-app/flexdb/pkg/entity"
	"gorm.io/gorm/clause"
)

// DeprecateColumn simulates 'ALTER TABLE DROP COLUMN'.
// DeprecateColumn will stop using the column:
// new data will not use this column, but existing data unaffected.
func (m *modelStore) DeprecateColumn(c *entity.Column) (err error) {
	updates := map[string]interface{}{
		"deprecated": c.Deprecated,
		"muser":      c.MUser,
		"mtime":      c.MTime,
	}
	if err = m.store.mdb().Omit(clause.Associations).Model(c).Updates(updates).Error; err != nil {
		return
	}

	c.Deprecated = true
	cache.Provider.Evict(modelStoreCacheHintID, c.ModelID)
	return
}

func (m *modelStore) UpdateColumn(c *entity.Column) (err error) {
	if err = m.Model.UpdateSlot(c, m.as); err != nil {
		return
	}

	updates := map[string]interface{}{
		// from form
		"name":     c.Name,
		"label":    c.Label,
		"remark":   c.Remark,
		"required": c.Required,
		"ro":       c.ReadOnly,

		"muser": c.MUser,
		"mtime": c.MTime,
	}
	if c.Choices != "" {
		updates["choices"] = c.Choices
	}

	err = m.store.mdb().Omit(clause.Associations).
		Where("id=? and model_id=?", c.ID, m.EntityModel().ID).
		Table(c.TableName()).
		Limit(1).
		Updates(updates).Error
	if err == nil {
		cache.Provider.Evict(modelStoreCacheHintID, c.ModelID)
	}

	return
}

func (m *modelStore) ReorderColumns(columns []entity.Column) (err error) {
	// all done in a tx
	tx := m.store.mdb().Begin()
	if profile.Debug() {
		tx = tx.Debug()
	}
	if err = tx.Error; err != nil {
		// connection pool exhausted?
		return
	}

	for _, c := range columns {
		updates := map[string]interface{}{
			"ordinal": c.Ordinal,
			"muser":   c.MUser,
			"mtime":   c.MTime,
		}
		if err = tx.
			Omit(clause.Associations).
			Model(&c).
			Updates(updates).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit().Error
	return
}

// AddColumns simulates 'ALTER TABLE ADD COLUMN'.
func (m *modelStore) AddColumns(columns []*entity.Column) (err error) {
	if err = m.Model.AddSlots(columns, m.as); err != nil {
		return err
	}

	// add columns within a tx
	tx := m.store.mdb().Begin()
	if profile.Debug() {
		tx = tx.Debug()
	}
	if err = tx.Error; err != nil {
		// connection pool exhausted?
		return
	}

	// TODO bulk insert for perf
	model := m.EntityModel()
	for _, c := range columns {
		if err = tx.Create(c).Error; err != nil {
			tx.Rollback()
			return err
		}

		if c.Relational() {
			rel := &entity.Relation{
				FromModelID: model.ID,
				FromSlot:    c.Slot,
				RefKind:     c.Kind,
				ToModelID:   c.RefModelID,
				ToSlot:      c.RefSlot,
			}
			if err = tx.Create(rel).Error; err != nil {
				tx.Rollback()
				return err
			}
		}

		model.Slots = append(model.Slots, c)
	}

	if err = tx.Commit().Error; err == nil {
		cache.Provider.Evict(appCacheHintID, model.AppID) // TODO in cluster, it does not work
	}

	return
}

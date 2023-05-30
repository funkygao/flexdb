package store

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/agile-app/flexdb/internal/profile"
	"github.com/agile-app/flexdb/internal/sla"
	"github.com/agile-app/flexdb/pkg/dto"
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/plugins/filter"
	"github.com/funkygao/log4go"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (m *modelStore) CreateRow(rd dto.RowData) (uint64, error) {
	row, clob, indexes, err := m.PrepareCreateRow(rd, m.as)
	if err != nil {
		return 0, err
	}

	now := time.Now()
	row.CTime = now
	row.MTime = now
	row.CUser = m.c.PIN()

	ctx := newTriggerContext(m)
	if err = m.InvokeTriggers(ctx, row, entity.TriggerBeforeCreate); err != nil {
		return 0, err
	}

	// within a tx, insert (1 row, [1] clob, N indexes)
	tx := m.store.ddb(m.AppID).Begin()
	if profile.Debug() {
		tx = tx.Debug()
	}

	// persist the row itself
	//if err := tx.Omit(clause.Associations).Table(row.TableName()).Create(row).Error; err != nil {
	if err := tx.Omit(clause.Associations).Create(row).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	// persist clob ifNec
	if clob != nil {
		clob.RowID = row.ID // row.ID generated within this tx
		//if err := tx.Omit(clause.Associations).Table(clob.TableName()).Create(clob).Error; err != nil {
		if err := tx.Omit(clause.Associations).Create(clob).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	// persist the indexes TODO perf in batch
	for _, idx := range indexes { // indexes should never be nil
		idx.SetRowID(row.ID) // row.ID generated within this tx
		//if err := tx.Omit(clause.Associations).Table(idx.TableName()).Create(idx).Error; err != nil {
		if err := tx.Omit(clause.Associations).Create(idx).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	// bingo!
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	if err = m.InvokeTriggers(ctx, row, entity.TriggerAfterCreate); err != nil {
		return 0, err
	}

	return row.ID, nil
}

func (m *modelStore) QuickSave(qs dto.QuickSave) error {
	return nil
}

func (m *modelStore) UpdateRow(rd dto.RowData) error {
	// within a tx, update (1 row, [1] clob, N indexes)
	tx := m.store.ddb(m.AppID).Begin()
	if profile.Debug() {
		tx = tx.Debug()
	}

	// fetch before update so that we can find diff
	snapshot := &entity.Row{ID: rd.ID()}
	if err := tx.Preload("Clob").Take(snapshot).Error; err != nil {
		tx.Rollback()
		return err
	}

	rowUpdates, row, clobRow, toAddindexes, toUpdateIndexes, toDeleteIndexes, err := m.PrepareUpdateRow(snapshot, rd, m.as)
	if err != nil {
		tx.Rollback()
		return err
	}

	if len(rowUpdates) == 0 && (clobRow == nil || clobRow.IsZero()) {
		tx.Rollback()
		return nil
	}

	ctx := newTriggerContext(m)
	if err = m.InvokeTriggers(ctx, row, entity.TriggerBeforeUpdate); err != nil {
		return err
	}

	// Update Selected Fields of the row
	if len(rowUpdates) > 0 {
		rowUpdates["ver"] = (snapshot.Ver + 1) % math.MaxInt16
		result := tx.
			Where("id=? and model_id=? and ver=?", row.ID, row.ModelID, snapshot.Ver).
			Table(row.TableName()).
			Updates(map[string]interface{}(rowUpdates))
		if err = result.Error; err != nil {
			tx.Rollback()
			return err
		}

		if result.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("optimistic lock fails: concurrent update")
		}
	}

	// update clob ifNec
	if clobRow != nil {
		if snapshot.HasClob() {
			if clobRow.IsZero() {
				// delete clob
				if err = tx.Omit(clause.Associations).
					Where("row_id=?", rd.ID()).
					Delete(clobRow).Error; err != nil {
					tx.Rollback()
					return err
				}
			} else {
				// update clob
				if err = tx.Omit(clause.Associations).
					Where("row_id=?", rd.ID()).
					Updates(clobRow).Error; err != nil {
					tx.Rollback()
					return err
				}
			}
		} else {
			// insert clob
			if err = tx.Omit(clause.Associations).Create(clobRow).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	for _, idx := range toAddindexes {
		if err = tx.Omit(clause.Associations).Create(idx).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, idx := range toDeleteIndexes {
		if err = tx.Omit(clause.Associations).
			Where("model_id=? and slot=? and val=? and row_id=?", m.ID, idx.SlotID(), idx.Value(), idx.GetRowID()).
			Delete(idx).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, idx := range toUpdateIndexes {
		if err = tx.Omit(clause.Associations).
			Model(idx.Index).
			Where("model_id=? and slot=? and val=? and row_id=?", m.ID, idx.SlotID(), idx.OriginalVal, idx.GetRowID()).
			Update("val", idx.Value()).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if m.Feature.ChangeAuditEnabled() {
		// DataChangelog
	}

	// bingo!
	if err := tx.Commit().Error; err != nil {
		return err
	}

	if err = m.InvokeTriggers(ctx, row, entity.TriggerAfterUpdate); err != nil {
		return err
	}

	return nil
}

func (m *modelStore) DeleteRow(rowID uint64) (err error) {
	row := &entity.Row{ID: rowID}
	if err = m.store.ddb(m.AppID).Preload("Clob", func(db *gorm.DB) *gorm.DB {
		// skip clob fields for performance
		return db.Select("row_id")
	}).Take(row).Error; err != nil {
		return err
	}

	indexes, clobRow, fakeDelete, err := m.PrepareDeleteRow(row, m.as)
	if err != nil {
		return err
	}

	ctx := newTriggerContext(m)
	if err = m.InvokeTriggers(ctx, row, entity.TriggerBeforeDelete); err != nil {
		return err
	}

	model := m.EntityModel()
	tx := m.store.ddb(m.AppID).Begin()
	if profile.Debug() {
		tx = tx.Debug()
	}
	for slot, index := range indexes {
		if err = tx.Table(index.TableName()).
			Where("model_id=? and slot=? and val=?", model.ID, slot, index.Value()).
			Delete(&index).Error; err != nil {
			tx.Rollback()
			return
		}
	}
	if clobRow != nil {
		if err = tx.Where("row_id=?", rowID).Delete(&clobRow).Error; err != nil {
			tx.Rollback()
			return
		}
	}

	if fakeDelete {
		// TODO
	}

	if m.Feature.ChangeAuditEnabled() {
		// DataChangelog
	}

	if err = tx.Delete(&row).Error; err != nil {
		tx.Rollback()
		return
	}

	if err = tx.Commit().Error; err != nil {
		return err
	}

	if err = m.InvokeTriggers(ctx, row, entity.TriggerAfterDelete); err != nil {
		return err
	}

	return
}

func (m *modelStore) RetrieveRow(rowID uint64, withChildren bool) (*dto.MasterDetail, error) {
	row := entity.Row{ID: rowID}
	if err := m.store.ddb(m.AppID).
		Preload("Clob").
		Take(&row).Error; err != nil {
		return nil, err
	}

	if withChildren { // TODO
	}

	return &dto.MasterDetail{Master: m.RenderRow(row, m.as, false)}, nil // TODO details
}

// TODO use row.TableName()
func (m *modelStore) FindRows(selectFields []string, cond dto.Criteria, pageIndex, pageSize int) (rds []dto.RowData, err error) {
	var (
		rowID     = cond.RowID()
		hasFilter = cond.Size() > 0
		rows      []entity.Row
	)

	if pageSize*pageIndex == profile.P.PaginationMaxRowsScan {
		// excel dump TODO
	}

	if hasFilter {
		log4go.Info("Find: %v", cond)
	}

	// 数据查询的同时获得总行数
	// db.Table(表名).Where(条件).Count(&totalSize).Sort(排序).Offset(offset).Limit(limit).Find(values)

	db := m.store.ddb(m.AppID).
		Limit(pageSize).Offset((pageIndex - 1) * pageSize)
	if len(selectFields) > 0 {
		db = db.Select(selectFields)
	}

	pin := m.c.PIN()
	if m.EntityModel().SingleRowKind() && !filter.SatisfyToBeKilledPermRule(pin) {
		err = db.
			Order("id desc").
			Where("model_id=? and cuser=?", m.ID, pin).
			Find(&rows).Error
	} else if !hasFilter {
		// will not use index, query row directly
		err = db.
			Order("id desc").
			Where("model_id=?", m.ID).
			Find(&rows).Error
	} else if rowID > 0 {
		err = db.Where("model_id=? AND id=?", m.ID, rowID).
			Find(&rows).Error
	} else {
		// MySQL merge index
		if !sla.Provider.BorrowSearchQuota(m.orgID) {
			return nil, errQuotaExhausted
		}

		condHit := false
		for i, ci := range cond {
			if slot := m.SlotByName(ci.Key); slot != nil {
				if index := slot.Indexer(); index != nil {
					alias := joinAliasTab[i]
					db = db.Joins(fmt.Sprintf(
						"JOIN %s %s ON %s.model_id=? AND %s.slot=? AND %s.val %s ?",
						index.TableName(),
						alias, alias, alias, alias, ci.Op,
					), m.ID, slot.Slot, ci.Val)
					db = db.Where(fmt.Sprintf("Data.id=%s.row_id", alias))
					condHit = true
				}
			}
		}
		if !condHit {
			db = db.Where("model_id=?", m.ID)
		}

		orderBy := "Data.id desc"
		if ob := cond.OrderBy(); ob != "" {
			for _, col := range m.BuiltinColumns() {
				if col.Sortable && ob == col.Name {
					orderBy = fmt.Sprintf("Data.%s %s", ob, cond.OrderDirection())
					break
				}
			}
		}

		// using MySQL merge index
		err = db.Order(orderBy).Find(&rows).Error
		sla.Provider.ReturnSearchQuota(m.orgID)
	}

	rds = make([]dto.RowData, 0, len(rows))
	for _, r := range rows {
		rds = append(rds, m.RenderRow(r, m.as, true))
	}

	return
}

package usf

import (
	"github.com/agile-app/flexdb/pkg/dto"
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/store"
)

type usf struct {
	*entity.Model
}

var rows map[int64][]dto.RowData = make(map[int64][]dto.RowData)

func init() {
	store.RegisterStoreEngine(entity.StoreEngineUSF, func(model *entity.Model) store.DataStore {
		return &usf{Model: model}
	})
}

func (m *usf) CreateRow(rd dto.RowData) (uint64, error) {
	if _, present := rows[m.ID]; !present {
		rows[m.ID] = make([]dto.RowData, 0, 5)
	}
	rows[m.ID] = append(rows[m.ID], rd)
	return 0, nil
}

func (m *usf) QuickSave(qs dto.QuickSave) (err error) {
	return
}

func (m *usf) UpdateRow(rd dto.RowData) (err error) {
	return
}

func (m *usf) DeleteRow(rowID uint64) (err error) {
	return
}

func (m *usf) RetrieveRow(rowID uint64, withChildren bool) (*dto.MasterDetail, error) {
	return nil, nil
}

func (m *usf) FindRows(selectFields []string, cond dto.Criteria, pageIndex, pageSize int) (rds []dto.RowData, err error) {
	rds = rows[m.ID]
	if rds == nil {
		rds = make([]dto.RowData, 0)
	}
	return
}

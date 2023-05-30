package upstream

import (
	"github.com/agile-app/flexdb/pkg/dto"
)

func (m *upstream) CreateRow(rd dto.RowData) (uint64, error) {
	return 0, nil
}

func (m *upstream) QuickSave(qs dto.QuickSave) (err error) {
	return
}

func (m *upstream) UpdateRow(rd dto.RowData) (err error) {
	return
}

func (m *upstream) DeleteRow(rowID uint64) (err error) {
	return
}

func (m *upstream) RetrieveRow(rowID uint64, withChildren bool) (*dto.MasterDetail, error) {
	return nil, nil
}

func (m *upstream) FindRows(selectFields []string, cond dto.Criteria, pageIndex, pageSize int) (rds []dto.RowData, err error) {
	return
}

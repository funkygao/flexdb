package mqsink

import (
	"github.com/agile-app/flexdb/pkg/connector"
	"github.com/agile-app/flexdb/pkg/dto"
)

func (m *mqSink) CreateRow(rd dto.RowData) (uint64, error) {
	err := connector.MQConnector(1).Produce(m.topic, rd)
	return 0, err
}

func (m *mqSink) QuickSave(qs dto.QuickSave) (err error) {
	return
}

func (m *mqSink) UpdateRow(rd dto.RowData) (err error) {
	return
}

func (m *mqSink) DeleteRow(rowID uint64) (err error) {
	return
}

func (m *mqSink) RetrieveRow(rowID uint64, withChildren bool) (*dto.MasterDetail, error) {
	return nil, nil
}

func (m *mqSink) FindRows(selectFields []string, cond dto.Criteria, pageIndex, pageSize int) (rds []dto.RowData, err error) {
	return
}

package data

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	excel "github.com/360EntSecGroup-Skylar/excelize"
	es "github.com/agile-app/flexdb/internal/spec"
	"github.com/agile-app/flexdb/pkg/dto"
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/store"
	"github.com/funkygao/log4go"
)

// TODO mapping横着放，这样便于防错；cuser逻辑；错误记录单独存放.
func (mc *dataHandler) importDataFromExcel(b []byte, app *entity.App, as store.AppStore) (ok, fail int, err error) {
	f, err := excel.OpenReader(bytes.NewReader(b))
	if err != nil {
		return 0, 0, err
	}

	// user guide sheets: _xxx_

	// mapping sheet
	// =============
	// Description row
	// modelID, sourceCol, targetCol

	// data sheets

	sheetMap := f.GetSheetMap()
	log4go.Info("sheets: %+v", sheetMap)

	var (
		mappings = make(map[int64]map[int]string, len(sheetMap)-1)
		models   = make(map[string]store.ModelStore, len(sheetMap)-1) // name -> store
		modelID  int64
	)

	// TODO 时间格式转换 transaction

	for i := es.ExcelSheetStartID; i <= len(sheetMap); i++ {
		sheetName := f.GetSheetName(i)
		if sheetName == es.DataMappingSheetName {
			rows := f.GetRows(sheetName)

			// magic validation: answer naughty users
			if rows[0][0] != "说明" {
				return 0, 0, fmt.Errorf("Invalid Excel schema")
			}

			// build the mappings
			// modelID, sourceCol, targetCol
			for _, mapping := range rows[es.DataMappingSheetHeaderRows:] {
				if strings.TrimSpace(mapping[0]) == "" {
					// empty or comment line
					continue
				}

				modelID, err = strconv.ParseInt(mapping[0], 10, 64)
				if err != nil {
					return 0, 0, fmt.Errorf("sheet:%s row:%v %v", sheetName, mapping, err)
				}

				modelFound := false
				for _, model := range app.Models {
					if model.ID == modelID {
						modelFound = true
						break
					}
				}
				if !modelFound {
					return 0, 0, fmt.Errorf("Table with id=%d not defined yet", modelID)
				}

				if _, present := mappings[modelID]; !present {
					mappings[modelID] = make(map[int]string, len(rows)-es.DataMappingSheetHeaderRows)
				}

				// {modelID, {sourceCol, targetCol}}
				idx := mapping[1][0] - 'A'
				mappings[modelID][int(idx)] = mapping[2]
			}

			for mid := range mappings {
				ms, err := as.LoadModel(mid, "Slots")
				if err != nil {
					return 0, 0, err
				}

				models[ms.EntityModel().Name] = ms
			}

			// now, the model mapping is ready
			log4go.Debug("mappings: %+v, models: %v", mappings, models)
			continue
		}

		if strings.HasPrefix(sheetName, "_") && strings.HasSuffix(sheetName, "_") {
			// sheet of manual guide, ignored
			continue
		}

		if len(mappings) == 0 {
			return 0, 0, fmt.Errorf("%s sheet MUST be before all data sheets", es.DataMappingSheetName)
		}

		ms, present := models[sheetName]
		if !present {
			//log4go.Warn("sheet: %s not defined in mapping, ignored", sheetName)
			continue
		}

		log4go.Info("importing data sheet: %s...", sheetName)
		modelID = ms.EntityModel().ID

		for _, row := range f.GetRows(sheetName)[es.DataSheetHeaderRows:] {
			log4go.Debug("excel: %s %+v", sheetName, row)

			rd := make(dto.RowData)
			for i, cell := range row {
				columnName := strings.TrimSpace(mappings[modelID][i])
				if columnName == "" {
					//log4go.Warn("sheet:%s column:%d not mapped, cell:%v ignored", sheetName, i, cell)
					continue
				}

				// if col is builtin, it will be nil
				if col := ms.EntityModel().SlotByName(columnName); col != nil && col.Kind == entity.ColumnDatetime {
					// TODO should reside inside pkg entity itself: logic leakage
					if t, err := time.Parse("01-02-06", cell); err == nil {
						rd[columnName] = t.Unix()
					}
				} else {
					rd[columnName] = cell // memo must be 'memo' in _mapping_
				}
			}

			log4go.Info("row: %v", rd)

			if _, err = ms.DS().CreateRow(rd); err != nil {
				log4go.Error("sheet:%s row:%v %v", sheetName, rd, err)
				fail++
			} else {
				ok++
			}
		}
	}

	return ok, fail, nil
}

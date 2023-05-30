package meta

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	excel "github.com/360EntSecGroup-Skylar/excelize"
	es "github.com/agile-app/flexdb/internal/spec"
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/store"
	"github.com/agile-app/flexdb/pkg/view"

	"github.com/funkygao/log4go"
)

func (mc *metaHandler) createModelsAndColumnsFromExcel(b []byte, app *entity.App, as store.AppStore, cuser string) error {
	f, err := excel.OpenReader(bytes.NewReader(b))
	if err != nil {
		return err
	}

	// FIXME sheet之间有顺序要求：如果有lookup类型字段，需要被ref的sheet放左侧
	log4go.Info("will create %d models", f.SheetCount)

	sheetMap := f.GetSheetMap()
	models := make(map[string]*entity.Model, len(sheetMap))
	for i := es.ExcelSheetStartID; i <= len(sheetMap); i++ {
		sheetName := f.GetSheetName(i)
		if sheetName == "" {
			log4go.Warn("empty sheet name")
			continue
		}

		if strings.HasPrefix(sheetName, "_") && strings.HasSuffix(sheetName, "_") {
			// sheet of manual guide
			continue
		}

		rows := f.GetRows(sheetName)

		// smoketest validity of the sheet, top left cell is the magic
		if rows[0][0] != "表类型" {
			return fmt.Errorf("Invalid Excel format")
		}

		log4go.Info("creating model[%s] %d columns: %+v", sheetName, len(rows), rows[es.ModelsTemplateHeaderSize:])

		var model entity.Model
		model.Name = sheetName
		now := time.Now()
		model.CTime = now
		model.MTime = now
		model.CUser = cuser
		model.Deleted = false
		model.Kind = view.ModelKindFromExcelTemplate(rows[0][1])
		model.AppID = app.ID
		model.Ver = 1
		model.Feature.EnableCRUD()
		if err := as.CreateModel(&model); err != nil {
			return err
		}

		log4go.Info("model: %s kind: %v, %v", sheetName, rows[0][1], model.Kind)

		// create columns
		ms := as.ModelStoreOf(&model)
		models[model.Name] = ms.EntityModel()
		columns := make([]*entity.Column, 0, len(rows)-es.ModelsTemplateHeaderSize)
		for i, row := range rows[es.ModelsTemplateHeaderSize:] {
			name := strings.TrimSpace(row[0])
			if name == "" {
				break
			}

			kind := strings.TrimSpace(row[1])
			if kind == "" {
				return fmt.Errorf("%s %s empty kind", sheetName, name)
			}

			c := &entity.Column{
				Name:    name,
				Kind:    view.ColumnKindFromExcelTemplate(kind),
				CUser:   cuser,
				CTime:   now,
				MTime:   now,
				Remark:  strings.TrimSpace(row[2]), // C
				Label:   strings.TrimSpace(row[3]), // D
				Default: strings.TrimSpace(row[8]), // I
			}
			c.Unique = strings.TrimSpace(row[4]) != ""          // E
			c.Indexed = strings.TrimSpace(row[5]) != ""         // F
			c.Required = strings.TrimSpace(row[6]) != ""        // G
			c.ReadOnly = strings.TrimSpace(row[7]) != ""        // H
			if c.Relational() && len(row) > 9 && row[9] != "" { // J
				// tuple: [sheet1, A8]
				tuple := strings.SplitN(f.GetCellFormula(sheetName, fmt.Sprintf("J%d", i+2)), "!", 2)
				refModel := models[tuple[0]]
				if refModel == nil {
					return fmt.Errorf("Ref: %s corresponding model not found", tuple[0])
				}

				c.RefModelID = refModel.ID
				refSlotName := f.GetCellValue(tuple[0], tuple[1])
				c.RefSlot = refModel.SlotByName(refSlotName).Slot
			}
			if len(row) > 10 && row[10] != "" { // K
				c.Sortable = strings.TrimSpace(row[10]) != ""
			}
			if len(row) > 11 && row[11] != "" { // L
				c.Choices = strings.TrimSpace(row[11])
			}

			columns = append(columns, c)
		}

		model.Slots = make([]*entity.Column, 0, len(columns))
		if err := ms.AddColumns(columns); err != nil {
			return err
		}
	}

	return nil
}

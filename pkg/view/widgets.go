package view

import (
	"github.com/agile-app/flexdb/pkg/entity"
	"github.com/agile-app/flexdb/pkg/view/amis"
)

var editWidgetTypes = map[entity.ColumnKind]string{
	entity.ColumnText:       amis.ATText,
	entity.ColumnEmail:      amis.ATEmail,
	entity.ColumnChoice:     amis.ATSelect,
	entity.ColumnURL:        amis.ATUrl,
	entity.ColumnFile:       amis.ATFile,
	entity.ColumnImage:      amis.ATImage,
	entity.ColumnInteger:    amis.ATNumber,
	entity.ColumnDatetime:   amis.ATDate,
	entity.ColumnMobile:     amis.ATText,
	entity.ColumnColor:      amis.ATColor,
	entity.ColumnPhone:      amis.ATText,
	entity.ColumnTextArea:   amis.ATTextArea,
	entity.ColumnAutoNumber: amis.ATStatic,
	entity.ColumnERP:        amis.ATText,
	entity.ColumnMany2One:   amis.ATSelect,
	entity.ColumnLookup:     amis.ATSelect,
	entity.ColumnFormula:    amis.ATFormula,
	entity.ColumnBoolean:    amis.ATCheckbox,
	entity.ColumnProgress:   amis.ATRange,
	entity.ColumnRating:     amis.ATRating,
	entity.ColumnCity:       amis.ATCity,
	entity.ColumnQRCode:     amis.ATQRCode,
	entity.ColumnPassword:   amis.ATPassword,
}

var (
	ExcelColumnKindMap = map[string]entity.ColumnKind{
		"text":     entity.ColumnText,
		"textarea": entity.ColumnTextArea,
		"date":     entity.ColumnDatetime,
		"choice":   entity.ColumnChoice,
		"lookup":   entity.ColumnLookup,
		"email":    entity.ColumnEmail,
		"color":    entity.ColumnColor,
		"progress": entity.ColumnProgress,
		"url":      entity.ColumnURL,
		"integer":  entity.ColumnInteger,
		"boolean":  entity.ColumnBoolean,
		"rating":   entity.ColumnRating,
		"many2one": entity.ColumnMany2One,
		"erp":      entity.ColumnERP,
		"phone":    entity.ColumnText,
		"city":     entity.ColumnCity,
		"qrcode":   entity.ColumnQRCode,
		"password": entity.ColumnPassword,
	}

	excelModelKindMap = map[string]entity.ModelKind{
		"singlerow": entity.ModelSingleRowPerUser,
		"custom":    entity.ModelCustom,
	}
)

// ColumnKindFromExcelTemplate returns a column kind from excel template file.
func ColumnKindFromExcelTemplate(ui string) entity.ColumnKind {
	return ExcelColumnKindMap[ui]
}

// ModelKindFromExcelTemplate returns a model kind from excel template file.
func ModelKindFromExcelTemplate(s string) entity.ModelKind {
	if k, present := excelModelKindMap[s]; present {
		return k
	}

	return entity.ModelNormal
}

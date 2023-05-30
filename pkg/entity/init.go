package entity

import (
	"github.com/agile-app/flexdb/internal/i18n"
	"github.com/agile-app/flexdb/pkg/dto"
)

var (
	appStatusLabels  map[AppStatus]string
	modelKindLabels  map[ModelKind]string
	columnKindLabels map[ColumnKind]string

	builtinColumns []*Column
	virtualColumns = map[ColumnKind]struct{}{ // TODO
		ColumnFormula: {},
	}

	alwaysIndexColumns = map[ColumnKind]struct{}{ // TODO
		ColumnMany2One: {},
	}

	clobColumns = map[ColumnKind]struct{}{
		ColumnTextArea: {},
	}

	illegalNames = []string{
		"$", "#", "*", "*", "^", "&", "@", "!", "~", "^", ",", "?", " ",
		"\"", "/",
		"{", "}",
		"[", "]",
		"(", ")",
		"<", ">",
	}
)

// Prepare prepares entity states.
func Prepare() {
	builtinColumns = []*Column{
		{Name: dto.KeyMemo, Kind: ColumnTextArea, Label: i18n.T("column.builtin.memo.label"), Indexed: false, ReadOnly: false},
		{Name: dto.KeyID, Kind: ColumnInteger, Label: i18n.T("column.builtin.id.label"), Indexed: true, ReadOnly: true},
		{Name: dto.KeyCTime, Kind: ColumnDatetime, Label: i18n.T("column.builtin.ctime.label"), Indexed: false, ReadOnly: true},
		{Name: dto.KeyCUser, Kind: ColumnERP, Label: i18n.T("column.builtin.cuser.label"), Indexed: false, ReadOnly: true},
		{Name: dto.KeyMTime, Kind: ColumnDatetime, Label: i18n.T("column.builtin.mtime.label"), Indexed: true, Sortable: true, ReadOnly: true},
		{Name: dto.KeyMUser, Kind: ColumnERP, Label: i18n.T("column.builtin.muser.label"), Indexed: false, ReadOnly: true},
	}

	// merge builtin column names into reservedColumnName
	for _, c := range builtinColumns {
		reservedColumnName[c.Name] = struct{}{}
	}

	columnKindLabels = map[ColumnKind]string{
		ColumnText:       "text",
		ColumnEmail:      "email",
		ColumnChoice:     "choice",
		ColumnURL:        "url",
		ColumnFile:       "file",
		ColumnImage:      "image",
		ColumnInteger:    "integer",
		ColumnDatetime:   "datetime",
		ColumnMobile:     "mobile",
		ColumnColor:      "color",
		ColumnPhone:      "phone",
		ColumnTextArea:   "textarea",
		ColumnAutoNumber: "autoseq",
		ColumnERP:        "erp",
		ColumnMany2One:   "many2one",
		ColumnLookup:     "lookup",
		ColumnFormula:    "formula",
		ColumnBoolean:    "boolean",
		ColumnProgress:   "progress",
		ColumnRating:     "rating",
		ColumnCity:       "city",
		ColumnQRCode:     "qrcode",
		ColumnPassword:   "password",
	}

	// TODO i18n
	modelKindLabels = map[ModelKind]string{
		ModelNormal:           "普通表",
		ModelSingleRowPerUser: "每人一行记录表",
		ModelCustom:           "定制表",
		ModelPerm:             "权限表",
	}

	// TODO i18n
	appStatusLabels = map[AppStatus]string{
		AppStatusInit:    "未上线",
		AppStatusOnline:  "已上线",
		AppStatusOffline: "已下线",
	}
}

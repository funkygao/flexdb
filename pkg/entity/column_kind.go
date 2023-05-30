package entity

// ColumnKind declares kind of a column.
// see https://help.salesforce.com/articleView?id=custom_field_types.htm&type=5
// Rich set of column kinds provided:
// classical (varchar, integer, boolean, ...)
// relational (many2one, many2many)
// functional
type ColumnKind int16

// Label returns human readable kind label.
func (c ColumnKind) Label() string {
	return columnKindLabels[c]
}

const (
	unknownColumnKind ColumnKind = 0

	ColumnText       ColumnKind = 1
	ColumnEmail      ColumnKind = 2
	ColumnChoice     ColumnKind = 3
	ColumnURL        ColumnKind = 4
	ColumnFile       ColumnKind = 5
	ColumnImage      ColumnKind = 6
	ColumnInteger    ColumnKind = 7
	ColumnDatetime   ColumnKind = 8
	ColumnMobile     ColumnKind = 9
	ColumnPhone      ColumnKind = 10
	ColumnTextArea   ColumnKind = 11
	ColumnAutoNumber ColumnKind = 12
	ColumnERP        ColumnKind = 13
	ColumnBoolean    ColumnKind = 14
	ColumnColor      ColumnKind = 15
	ColumnProgress   ColumnKind = 16
	ColumnRating     ColumnKind = 17
	ColumnCity       ColumnKind = 18
	ColumnQRCode     ColumnKind = 19
	ColumnPassword   ColumnKind = 22

	ColumnMany2One ColumnKind = 31
	ColumnLookup   ColumnKind = 32

	ColumnFormula ColumnKind = 41
)

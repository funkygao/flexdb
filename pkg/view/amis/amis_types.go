package amis

const (
	// AT means amis type
	ATRemark       = "remark"
	ATPage         = "page"
	ATCrud         = "crud"
	ATDialog       = "dialog"
	ATDrawer       = "drawer"
	ATNav          = "nav"
	ATText         = "text"
	ATTextArea     = "textarea"
	ATEmail        = "email"
	ATDivider      = "divider"
	ATSelect       = "select"
	ATTreeSelect   = "tree-select"
	ATNestedSelect = "nested-select"
	ATUrl          = "url"
	ATFile         = "file"
	ATHidden       = "hidden"
	ATImage        = "image"
	ATNumber       = "number"
	ATDate         = "date"
	ATTransfer     = "transfer"
	ATDatetime     = "datetime"
	ATDateRange    = "date-range"
	ATTime         = "time"
	ATColor        = "color"
	ATStatic       = "static"
	ATRadios       = "radios"
	ATStatus       = "status"
	ATPanel        = "panel"
	ATFormula      = "formula"
	ATSwitch       = "switch"
	ATCheckbox     = "checkbox"
	ATCheckboxes   = "checkboxes"
	ATRange        = "range"
	ATProgress     = "progress"
	ATRating       = "rating"
	ATCity         = "city"
	ATQRCode       = "qr-code"
	ATRepeat       = "repeat"
	ATPicker       = "picker"
	ATTag          = "tag"
	ATEditor       = "editor" // language: Javascript | html | json | css
	ATPassword     = "password"
	ATMapping      = "mapping"
	ATList         = "list"

	ATCombo         = "combo"
	ATButtonToolbar = "button-toolbar" // https://ufologist.github.io/page-schema/_demo/index.html?schema=https://ufologist.github.io/page-schema/_demo/form-reaction.json
	ATButton        = "button"
	ATCard          = "card"
	ATCards         = "cards"
	ATIcon          = "icon"
	ATAction        = "action"
	ATOperation     = "operation"
	ATTabs          = "tabs"
	ATTpl           = "tpl"
	ATFieldSet      = "fieldSet"

	// amis action type
	AATDrawer = "drawer"
	AATUrl    = "url"
	AATLink   = "link"
	AAAjax    = "ajax"

	// VT means validation type
	VTNumber         = "isNumeric"                 // 数字
	VTLetterOrNumber = "isAlphanumeric"            // 字母或数字
	VTInteger        = "isInt"                     // 整形
	VTLengthLimit    = "minLength:%d,maxLength:%d" // 长度限制
	VTNumberLimit    = "maximum:%d,minimum:%d"     // 数值限制
	VTRegex          = "matchRegexp:/^abc/"        // 正则 TODO
)

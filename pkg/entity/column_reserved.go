package entity

import (
	"strings"
)

var (
	// reserved columns already reside in row: should not waste resource
	reservedColumnName = map[string]struct{}{
		"csince":      {},
		"msince":      {},
		"ts":          {},
		"yn":          {},
		"is_delete":   {},
		"deleted":     {},
		"createtime":  {},
		"create_time": {},
		"updatetime":  {},
		"update_time": {},
		"createuser":  {},
		"create_user": {},
		"updateuser":  {},
		"update_user": {},

		"org_id": {},
		"orgid":  {},

		"modelid":  {},
		"model_id": {},

		"ident": {},
		"slug":  {},
	}
)

func columnNameReserved(name string) bool {
	if _, present := reservedColumnName[strings.ToLower(name)]; present {
		return true
	}

	return false
}

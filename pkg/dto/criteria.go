package dto

import (
	"encoding/json"
	"strconv"

	"github.com/agile-app/flexdb/pkg/api"
)

// Criteria is the row data find condition. TODO might use OData
type Criteria []CriteriaItem

// Valid validate the criteria.
func (c Criteria) Valid() bool {
	for _, item := range c {
		if !item.valid() {
			return false
		}
	}

	return true
}

func (c Criteria) Append(i CriteriaItem) Criteria {
	return append(c, i)
}

func (c Criteria) Size() int {
	return len(c)
}

func (c Criteria) OrderBy() string {
	for _, item := range c {
		if item.Key == api.QueryOrderBy {
			return item.Val
		}
	}

	return ""
}

func (c Criteria) OrderDirection() string {
	for _, item := range c {
		if item.Key == api.QueryOrderDirection {
			return item.Val
		}
	}

	return ""
}

func (c Criteria) RowID() int64 {
	for _, item := range c {
		if item.Key == api.RowID {
			id, _ := strconv.ParseInt(item.Val, 10, 64)
			return id
		}
	}

	return 0
}

func (c Criteria) JSONString() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// CriteriaItem is a criteria item.
type CriteriaItem struct {
	Key string
	Op  string
	Val string
}

func (i CriteriaItem) valid() bool {
	if i.Key == "" || i.Op == "" || i.Val == "" {
		return false
	}

	if _, present := validCriteriaOps[i.Op]; !present {
		return false
	}

	return true
}

var validCriteriaOps = map[string]struct{}{
	"=":  {},
	">":  {},
	"<":  {},
	">=": {},
	"<=": {},
}

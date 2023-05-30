package amis

var excludedCriteria = map[string]struct{}{
	"app":     {}, // app id
	"aid":     {}, // app id
	"mid":     {}, // model id
	"perPage": {},
	"page":    {},
	"id":      {}, // id from url, rowid is in filter
}

var criteriaOperators = map[string]struct{}{
	"=":    {},
	">=":   {},
	"<=":   {},
	">":    {},
	"<":    {},
	"like": {},
}

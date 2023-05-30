package profile

import "errors"

var (
	ErrorEmptyOrgName   = errors.New("empty org.id")
	ErrorEmptyOrgID     = errors.New("empty org.id")
	ErrorEmptyOrgSecret = errors.New("empty org.secret")
)

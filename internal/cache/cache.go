// Package cache provides acceleration for the meta framework.
package cache

import "time"

// Cache accelerates entity accessing.
type Cache interface {
	Get(hintID, id int64) (interface{}, error)
	SetWithExpire(hintID, id int64, value interface{}, expiration time.Duration) error
	Evict(hintID, id int64)
}

// Provider is the default cache implementation.
var Provider Cache

package blackhole

import (
	"errors"
	"time"

	"github.com/agile-app/flexdb/internal/cache"
)

var errMiss = errors.New("cache missed")

type blackhole struct {
}

func Offer() cache.Cache {
	return &blackhole{}
}

func (c *blackhole) Get(hintID, id int64) (interface{}, error) {
	return nil, errMiss
}

func (c *blackhole) Evict(hintID, id int64) {
}

func (c *blackhole) SetWithExpire(hintID, id int64, value interface{}, expiration time.Duration) error {
	return nil
}

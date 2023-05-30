package gcache

import (
	"time"

	"github.com/agile-app/flexdb/internal/cache"
	"github.com/bluele/gcache"
)

type gocache struct {
	gcache.Cache
}

func Offer() cache.Cache {
	return &gocache{
		Cache: gcache.New(10 << 10).
			LFU().
			Build(),
	}
}

func (c *gocache) Get(hintID, id int64) (interface{}, error) {
	return c.Cache.Get(generateMergeID(hintID, id))
}

func (c *gocache) Evict(hintID, id int64) {
	c.Cache.Remove(generateMergeID(hintID, id))
}

func (c *gocache) SetWithExpire(hintID, id int64, value interface{}, expiration time.Duration) error {
	return c.Cache.SetWithExpire(generateMergeID(hintID, id), value, expiration)
}

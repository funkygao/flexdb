package bcache

import (
	"errors"
	"time"

	"github.com/agile-app/flexdb/internal/cache"
	"github.com/allegro/bigcache"
)

var errMiss = errors.New("cache missed")

type bcache struct {
	*bigcache.BigCache
}

func Offer() cache.Cache {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	return &bcache{BigCache: cache}
}

func (c *bcache) Get(hintID, id int64) (interface{}, error) {
	return nil, errMiss
}

func (c *bcache) Evict(hintID, id int64) {
}

func (c *bcache) SetWithExpire(hintID, id int64, value interface{}, expiration time.Duration) error {
	return nil
}

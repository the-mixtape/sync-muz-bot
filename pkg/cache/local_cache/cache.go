package local_cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type BotCache struct {
	localCache *cache.Cache
}

func NewBotCache(defaultExpiration time.Duration, cleanupInterval time.Duration) *BotCache {
	c := cache.New(defaultExpiration, cleanupInterval)
	return &BotCache{localCache: c}
}

func (c *BotCache) Get(key string) interface{} {
	value, ok := c.localCache.Get(key)
	if ok {
		return value
	}
	return nil
}

func (c *BotCache) Set(key string, value interface{}) {
	c.localCache.Set(key, value, cache.DefaultExpiration)
}

func (c *BotCache) Delete(key string) {
	c.localCache.Delete(key)
}

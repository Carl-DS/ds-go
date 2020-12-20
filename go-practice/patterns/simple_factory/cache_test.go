package simple_factory

import (
	"log"
	"testing"
)

func TestCacheFactory_Create(t *testing.T) {
	cacheFactory := &CacheFactory{}
	cache, err := cacheFactory.Create(redis)
	if err != nil {
		t.Error(err)
	}
	cache.Set("redisKey", "redis value 缓存值")
	log.Print(cache.Get("redisKey"))

	memCache, error := cacheFactory.Create(mem)
	if error != nil {
		t.Error(error)
	}
	memCache.Set("memCache", "memCache value 缓存值")
	log.Print(memCache.Get("memCache"))
}

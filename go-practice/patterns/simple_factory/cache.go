package simple_factory

import "errors"

// 定义一个 Cache 接口，作为父类
type Cache interface {
	Get(key string) string
	Set(key, value string)
}

// 实现具体的 Cache：RedisCache
type RedisCache struct {
	data map[string]string
}

func (redis *RedisCache) Set(key, value string) {
	redis.data[key] = value
}

func (redis *RedisCache) Get(key string) string {
	return "redis " + redis.data[key]
}

// 实现具体的 Cache：MemCache
type MemCache struct {
	data map[string]string
}

func (mem *MemCache) Set(key, value string) {
	mem.data[key] = value
}

func (mem *MemCache) Get(key string) string {
	return "mem " + mem.data[key]
}

type cacheType int

const (
	redis cacheType = iota
	mem
)

// 实现 Cache 的简单工厂
type CacheFactory struct {
}

func (factory *CacheFactory) Create(cacheType cacheType) (Cache, error) {
	if cacheType == redis {
		return &RedisCache{
			data: map[string]string{},
		}, nil
	}

	if cacheType == mem {
		return &MemCache{
			data: map[string]string{},
		}, nil
	}

	return nil, errors.New("error cache type")
}

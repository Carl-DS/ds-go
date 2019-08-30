package memory

import (
	"github.com/go-redis/redis"
)

var StorageRedisClient redis.Client

/**
 * 初始化实时存储服务器
 */
func InitializationStorageRedis() {
	StorageRedisClient = *redis.NewClient(&redis.Options{
		Addr:     "172.17.50.225:6379",
		Password: "_Zy0N85fz_chri",
		DB:       12,
	})
}

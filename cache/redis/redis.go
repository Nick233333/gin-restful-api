package redis

import (
	"gin-restful-api/pkg/setting"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func CreateRedis() *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:       setting.RedisSetting.Addr,
		Password:   setting.RedisSetting.Password,
		DB:         setting.RedisSetting.DB,
		MaxRetries: setting.RedisSetting.MaxRetries,
	})
	_, error := RedisClient.Ping().Result()
	if error != nil {
		panic("redis connection error")
	}
	return RedisClient

}

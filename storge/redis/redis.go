package redis

import "github.com/redis/go-redis/v9"

var redisClient *redis.Client

func Init() {
	redisClientTmp := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:32768",
		Username: "default",
		Password: "redispw",
	})
	redisClient = redisClientTmp
}

func GetRedisClient() *redis.Client {
	return redisClient
}

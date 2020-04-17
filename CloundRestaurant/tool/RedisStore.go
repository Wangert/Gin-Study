package tool

import (
	"github.com/go-redis/redis"
)

type RedisStore struct {
	client *redis.Client
}

var Redis RedisStore

//初始化redis参数
func InitRedis() *RedisStore {

	redisConfig := GetConfig().Redis
	//创建一个redis对象
	client := redis.NewClient(&redis.Options{
		Addr:	   redisConfig.Addr + ":" + redisConfig.Port,
		Password:  redisConfig.Password,
		DB: 	   redisConfig.Db,
	})

	Redis = RedisStore{client}

	return &Redis
}

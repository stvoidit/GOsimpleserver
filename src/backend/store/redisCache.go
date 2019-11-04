package store

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

// myRedis - ...
type myRedis struct {
	*redis.Client
}

// Redis - ...
var Redis myRedis

func init() {
	Redis = myRedis{redisConn()}
}

func redisConn() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: Config.password,
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}

// SetJSON - for cache json response
func (r *myRedis) SetJSON(key string, val interface{}) {
	b, _ := json.Marshal(val)
	Redis.Set(key, b, 15*time.Minute)
}

// GetCache - ...
func (r *myRedis) GetCache(key string) ([]byte, error) {
	val, err := Redis.Get(key).Bytes()
	if err != nil {
		return nil, err
	}
	return val, nil
}

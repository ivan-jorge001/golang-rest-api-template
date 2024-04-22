package cache

import (
	"aitrainer-api/config"
	"context"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client
var Ctx = context.Background()

func InitRedis(cfg *config.Config) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // Redis server address
		Password: "",           // Password, leave empty if none
		DB:       0,            // Default DB
	})
}

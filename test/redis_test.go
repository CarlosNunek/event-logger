package test

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"eventlogger/config"
)

func TestRedisPublish(t *testing.T) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: config.GetEnv("REDIS_HOST", "localhost") + ":" + config.GetEnv("REDIS_PORT", "6379"),
	})

	err := rdb.Publish(ctx, "login", "mensaje de test").Err()
	if err != nil {
		t.Errorf("Error publicando en Redis: %v", err)
	}
}

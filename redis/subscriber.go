package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"eventlogger/config"
	"eventlogger/influx"
)

func SubscribeAndLog(ctx context.Context) error {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s",
			config.GetEnv("REDIS_HOST", "localhost"),
			config.GetEnv("REDIS_PORT", "6379")),
	})

	sub := rdb.Subscribe(ctx, "usuarios", "login")
	ch := sub.Channel()

	for msg := range ch {
		fmt.Printf("📥 Evento recibido [%s]: %s\n", msg.Channel, msg.Payload)
		err := influx.WriteEvent(msg.Channel, msg.Payload, time.Now())
		if err != nil {
			log.Printf("❌ Error al escribir en InfluxDB: %v\n", err)
		}
	}

	return nil
}

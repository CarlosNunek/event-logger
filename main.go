package main

import (
	"context"
	"eventlogger/config"
	"eventlogger/redis"
	"fmt"
	"log"
)

func main() {
	config.LoadEnv()
	ctx := context.Background()
	fmt.Println("📡 Escuchando eventos de Redis...")
	err := redis.SubscribeAndLog(ctx)
	if err != nil {
		log.Fatalf("Error al escuchar Redis: %v", err)
	}
}

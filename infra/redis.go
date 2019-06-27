package infra

import (
	"github.com/bolg-developers/BoLG-Server/internal/redisutil"
	"github.com/bolg-developers/BoLG-Server/server/config"
	"github.com/go-redis/redis"
)

var (
	clients       = make([]*redis.Client, config.Env().MaxRoomCount)
	ClientManager = redisutil.ClientManager{Unused: clients}
)

func init() {
	for i := 0; i < config.Env().MaxRoomCount; i++ {
		clients[i] = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       i,
		})
	}
}

package app

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
)

type RedisEventBus struct {
	Client *redis.Client
	Channel string
}

func (b *RedisEventBus) Publish(event interface{}) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return b.Client.Publish(context.Background(), b.Channel, data).Err()
}

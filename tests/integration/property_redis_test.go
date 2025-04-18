package integration

import (
	"context"
	"encoding/json"
	"testing"
	"time"
	"overhelloworld/internal/app"
	"overhelloworld/internal/domain"
	"github.com/go-redis/redis/v8"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
)

func TestRedisEventBus_Property(t *testing.T) {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		t.Skip("REDIS_ADDR not set; skipping Redis property test")
	}
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 10
	props := gopter.NewProperties(parameters)

	props.Property("Published events are received by subscriber", prop.ForAll(
		func(msg string) bool {
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			client := redis.NewClient(&redis.Options{Addr: redisAddr})
			bus := &app.RedisEventBus{Client: client, Channel: "hello_events"}
			ch := make(chan string, 1)
			pubsub := client.Subscribe(ctx, "hello_events")
			defer func() { _ = pubsub.Unsubscribe(ctx, "hello_events") }()
			go func() {
				for m := range pubsub.Channel() {
					var evt domain.HelloSaid
					if err := json.Unmarshal([]byte(m.Payload), &evt); err == nil {
						ch <- evt.Message
					}
				}
			}()
			if err := bus.Publish(domain.HelloSaid{Message: msg}); err != nil {
				return false
			}
			select {
			case got := <-ch:
				return got == msg
			case <-ctx.Done():
				return false
			}
		},
		gen.RegexMatch("[a-zA-Z0-9 .,!?:;'\"-]{1,100}"),
	))

	props.TestingRun(t)
}

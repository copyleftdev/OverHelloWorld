package main

import (
    "log"
    "net/http"
    "os"

    "github.com/go-redis/redis/v8"
    "overhelloworld/internal/api"
    "overhelloworld/internal/app"
    "overhelloworld/internal/domain"
    "overhelloworld/internal/readmodel"
)

func main() {
    // Observability hooks
    app.InitMetrics()
    go func() {
        log.Println("Prometheus metrics at :8080/metrics")
        http.Handle("/metrics", app.MetricsHandler())
    }()
    app.Trace("Starting OverHelloWorld server")
    app.Metric("server_start")
    app.Log("Bootstrapping components")

    // Set up plugin
    plugins := []app.Plugin{&app.ASCIIPlugin{}, &app.TTSPlugin{}, &app.LEDPlugin{}}

    // Set up read model
    readModel := &readmodel.HelloReadModel{}

    // Set up event bus (prefer Redis if REDIS_ADDR is set)
    var eventBus app.EventBus
    redisAddr := os.Getenv("REDIS_ADDR")
    if redisAddr != "" {
        client := redis.NewClient(&redis.Options{Addr: redisAddr})
        eventBus = &app.RedisEventBus{Client: client, Channel: "hello_events"}
        app.Log("Using RedisEventBus")
    } else {
        eventBus = &app.InMemoryEventBus{}
        app.Log("Using InMemoryEventBus")
    }

    // Set up file event store for event sourcing
    eventStore := &app.FileEventStore{Path: "events.jsonl"}
    // Replay events on startup
    if err := eventStore.Replay(func(evt map[string]interface{}) {
        if msg, ok := evt["Message"].(string); ok {
            e := domain.HelloSaid{
                ID: evt["ID"].(string),
                Message: msg,
                SaidAt: domain.ParseTime(evt["SaidAt"]),
            }
            readModel.Add(e)
            for _, p := range plugins {
                p.OnHelloSaid(e.Message)
            }
        }
    }); err != nil {
        log.Printf("Replay error: %v", err)
    }

    // Wrap event bus to update read model, plugins, and persist events
    eventBus = &eventBusWithHooks{
        EventBus:  eventBus,
        ReadModel: readModel,
        Plugins:   plugins,
        EventStore: eventStore,
    }

    commandHandler := &app.HelloCommandHandler{EventBus: eventBus}
    helloAPI := &api.HelloAPI{CommandHandler: commandHandler}
    helloQueryAPI := &api.HelloQueryAPI{ReadModel: readModel}

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if _, err := w.Write([]byte("Hello, OverEngineered World!")); err != nil {
            log.Printf("Write error: %v", err)
        }
    })
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            helloAPI.SayHelloHandler(w, r)
        } else if r.Method == http.MethodGet {
            helloQueryAPI.GetHellosHandler(w, r)
        } else {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}

// eventBusWithHooks wraps EventBus to update read model, call plugins, and persist events

type eventBusWithHooks struct {
    app.EventBus
    ReadModel *readmodel.HelloReadModel
    Plugins   []app.Plugin
    EventStore *app.FileEventStore
}

func (b *eventBusWithHooks) Publish(event interface{}) error {
    if e, ok := event.(domain.HelloSaid); ok {
        b.ReadModel.Add(e)
        for _, p := range b.Plugins {
            p.OnHelloSaid(e.Message)
        }
        if err := b.EventStore.Append(e); err != nil {
            log.Printf("Append error: %v", err)
        }
        app.Trace("HelloSaid event processed")
        app.Metric("hello_said")
    }
    return b.EventBus.Publish(event)
}

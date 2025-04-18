package integration

import (
	"os"
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
	"encoding/json"
	"overhelloworld/internal/app"
	"overhelloworld/internal/domain"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestHelloEventPersistence_Property(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 50
	props := gopter.NewProperties(parameters)

	props.Property("Any string message is persisted and replayed", prop.ForAll(
		func(msg string) bool {
			// Clean up file store before each run
			os.Remove("test_events.jsonl")
			store := &app.FileEventStore{Path: "test_events.jsonl"}
			event := domain.HelloSaid{
				ID:      "test-id",
				Message: msg,
				SaidAt:  domain.ParseTime("2025-04-17T20:00:00Z"),
			}
			if err := store.Append(event); err != nil {
				t.Fatalf("failed to append event: %v", err)
			}

			replayed := false
			if err := store.Replay(func(evt map[string]interface{}) {
				if m, ok := evt["Message"].(string); ok && m == msg {
					replayed = true
				}
			}); err != nil {
				return false
			}
			return replayed
		},
		gen.AnyString(),
	))

	props.TestingRun(t)
}

// --- Plugin Side Effect Property Test ---
type mockPlugin struct {
	Received []string
}
func (p *mockPlugin) OnHelloSaid(msg string) {
	p.Received = append(p.Received, msg)
}

func TestPluginReceivesAllMessages_Property(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 30
	props := gopter.NewProperties(parameters)

	props.Property("Plugin receives any string message", prop.ForAll(
		func(msg string) bool {
			plugin := &mockPlugin{}
			bus := &app.InMemoryEventBus{}
			hooks := &eventBusWithPlugin{
				EventBus: bus,
				Plugin:   plugin,
			}
			handler := &app.HelloCommandHandler{EventBus: hooks}
			if err := handler.Handle(app.SayHelloCommand{Message: msg}); err != nil {
				return false
			}
			return len(plugin.Received) == 1 && plugin.Received[0] == msg
		},
		gen.AnyString(),
	))

	props.TestingRun(t)
}

type eventBusWithPlugin struct {
	app.EventBus
	Plugin *mockPlugin
}

func (b *eventBusWithPlugin) Publish(event interface{}) error {
	if e, ok := event.(domain.HelloSaid); ok {
		b.Plugin.OnHelloSaid(e.Message)
	}
	return b.EventBus.Publish(event)
}

// --- API Flow Property Test ---
func TestAPIFlow_Property(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	props := gopter.NewProperties(parameters)

	props.Property("POST then GET returns the message", prop.ForAll(
		func(msg string) bool {
			os.Remove("test_api_events.jsonl")
			bus := &app.InMemoryEventBus{}
			handler := &app.HelloCommandHandler{EventBus: bus}
			readModel := &mockReadModel{}
			apiHandler := &mockAPI{CommandHandler: handler, ReadModel: readModel}

			mux := http.NewServeMux()
			mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodPost {
					apiHandler.SayHelloHandler(w, r)
				} else if r.Method == http.MethodGet {
					apiHandler.GetHellosHandler(w, r)
				}
			})
			ts := httptest.NewServer(mux)
			defer ts.Close()

			// POST
			body := map[string]string{"message": msg}
			b, _ := json.Marshal(body)
			resp, err := http.Post(ts.URL+"/hello", "application/json", bytes.NewReader(b))
			if err != nil || resp.StatusCode != http.StatusAccepted {
				return false
			}

			// GET
			resp, err = http.Get(ts.URL+"/hello")
			if err != nil {
				return false
			}
			var hellos []domain.Hello
			if err := json.NewDecoder(resp.Body).Decode(&hellos); err != nil {
				return false
			}
			resp.Body.Close()
			for _, h := range hellos {
				if h.Message == msg {
					return true
				}
			}
			return false
		},
		gen.RegexMatch("[a-zA-Z0-9 .,!?:;'\"-]{1,100}"), // Only printable ASCII strings
	))

	props.TestingRun(t)
}

type mockReadModel struct{
	msgs []string
}
func (m *mockReadModel) Add(event domain.HelloSaid) {
	m.msgs = append(m.msgs, event.Message)
}
func (m *mockReadModel) All() []domain.Hello {
	h := make([]domain.Hello, len(m.msgs))
	for i, msg := range m.msgs {
		h[i] = domain.Hello{Message: msg}
	}
	return h
}

type mockAPI struct {
	CommandHandler *app.HelloCommandHandler
	ReadModel      *mockReadModel
}
func (api *mockAPI) SayHelloHandler(w http.ResponseWriter, r *http.Request) {
	var req struct{ Message string `json:"message"` }
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cmd := app.SayHelloCommand{Message: req.Message}
	if err := api.CommandHandler.Handle(cmd); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	api.ReadModel.Add(domain.HelloSaid{Message: req.Message})
	w.WriteHeader(http.StatusAccepted)
	if _, err := w.Write([]byte("Hello event accepted!")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (api *mockAPI) GetHellosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(api.ReadModel.All()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

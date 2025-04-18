package integration

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"overhelloworld/internal/api"
	"overhelloworld/internal/app"
	"overhelloworld/internal/domain"
	"overhelloworld/internal/readmodel"
)

func TestHelloFlow(t *testing.T) {
	readModel := &readmodel.HelloReadModel{}
	bus := &app.InMemoryEventBus{}
	busWithHooks := &eventBusWithHooks{
		EventBus:  bus,
		ReadModel: readModel,
		Plugins:   []app.Plugin{},
	}
	handler := &app.HelloCommandHandler{EventBus: busWithHooks}
	apiHandler := &api.HelloAPI{CommandHandler: handler}
	queryAPI := &api.HelloQueryAPI{ReadModel: readModel}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			apiHandler.SayHelloHandler(w, r)
		} else if r.Method == http.MethodGet {
			queryAPI.GetHellosHandler(w, r)
		}
	})
	ts := httptest.NewServer(mux)
	defer ts.Close()

	// POST /hello
	body := map[string]string{"message": "Integration Hello"}
	b, _ := json.Marshal(body)
	resp, err := http.Post(ts.URL+"/hello", "application/json", bytes.NewReader(b))
	if err != nil || resp.StatusCode != http.StatusAccepted {
		t.Fatalf("POST failed: %v, status: %v", err, resp.StatusCode)
	}

	// GET /hello
	resp, err = http.Get(ts.URL + "/hello")
	if err != nil {
		t.Fatalf("GET failed: %v", err)
	}
	data, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if !bytes.Contains(data, []byte("Integration Hello")) {
		t.Fatalf("Expected response to contain 'Integration Hello', got: %s", string(data))
	}
}

type eventBusWithHooks struct {
	app.EventBus
	ReadModel *readmodel.HelloReadModel
	Plugins   []app.Plugin
}

func (b *eventBusWithHooks) Publish(event interface{}) error {
	if e, ok := event.(domain.HelloSaid); ok {
		b.ReadModel.Add(e)
		for _, p := range b.Plugins {
			p.OnHelloSaid(e.Message)
		}
	}
	return b.EventBus.Publish(event)
}

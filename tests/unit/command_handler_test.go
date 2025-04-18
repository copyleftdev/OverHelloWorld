package unit

import (
	"testing"
	"overhelloworld/internal/app"
)

func TestHelloCommandHandler_PublishesEvent(t *testing.T) {
	bus := &app.InMemoryEventBus{}
	handler := &app.HelloCommandHandler{EventBus: bus}
	cmd := app.SayHelloCommand{Message: "Test Hello"}
	err := handler.Handle(cmd)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(bus.Events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(bus.Events))
	}
}

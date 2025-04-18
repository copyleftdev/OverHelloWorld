package app

import (
	"github.com/google/uuid"
	"overhelloworld/internal/domain"
	"time"
)

type SayHelloCommand struct {
	Message string
}

type CommandHandler interface {
	Handle(cmd interface{}) error
}

type HelloCommandHandler struct {
	EventBus EventBus
}

func (h *HelloCommandHandler) Handle(cmd interface{}) error {
	c, ok := cmd.(SayHelloCommand)
	if !ok {
		return nil
	}
	event := domain.HelloSaid{
		ID:      uuid.NewString(),
		Message: c.Message,
		SaidAt:  time.Now(),
	}
	return h.EventBus.Publish(event)
}

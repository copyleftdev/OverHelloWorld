package app

type EventBus interface {
	Publish(event interface{}) error
}

type InMemoryEventBus struct {
	Events []interface{}
}

func (b *InMemoryEventBus) Publish(event interface{}) error {
	b.Events = append(b.Events, event)
	return nil
}

package readmodel

import (
	"sync"
	"overhelloworld/internal/domain"
)

type HelloReadModel struct {
	mu sync.RWMutex
	Hellos []domain.Hello
}

func (rm *HelloReadModel) Add(event domain.HelloSaid) {
	rm.mu.Lock()
	rm.Hellos = append(rm.Hellos, domain.Hello(event))
	rm.mu.Unlock()
}

func (rm *HelloReadModel) All() []domain.Hello {
	rm.mu.RLock()
	defer rm.mu.RUnlock()
	return append([]domain.Hello(nil), rm.Hellos...)
}

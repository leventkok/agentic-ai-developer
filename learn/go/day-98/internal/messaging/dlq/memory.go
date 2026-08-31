package dlq

import (
	"sync"

	"learn/go/day-98/internal/messaging"
)

type Memory struct {
	mu       sync.Mutex
	messages []messaging.Event
}

func NewMemory() *Memory {
	return &Memory{}
}

func (m *Memory) Push(evt messaging.Event) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.messages = append(m.messages, evt)
}

func (m *Memory) Len() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.messages)
}

func (m *Memory) Messages() []messaging.Event {
	m.mu.Lock()
	defer m.mu.Unlock()
	out := append([]messaging.Event(nil), m.messages...)
	return out
}

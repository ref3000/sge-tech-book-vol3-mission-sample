package main

import (
	"context"
)

type Manager struct {
	events map[EventType][]Event
}

func NewManager() *Manager {
	return &Manager{
		events: make(map[EventType][]Event),
	}
}

type typeManagerKey struct{}

func SetManager(ctx context.Context, manager *Manager) context.Context {
	return context.WithValue(ctx, typeManagerKey{}, manager)
}

func GetManager(ctx context.Context) *Manager {
	manager, ok := ctx.Value(typeManagerKey{}).(*Manager)
	if !ok {
		return nil
	}
	return manager
}

func (m *Manager) Push(event Event) {
	t := event.EventType()
	m.events[t] = append(m.events[t], event)
}

func (m *Manager) Events(t EventType) []Event {
	return m.events[t]
}

func Events[T Event](ctx context.Context) []T {
	m := GetManager(ctx)
	var obj T
	events := m.Events(obj.EventType())
	s := make([]T, 0, len(events))
	for _, e := range events {
		s = append(s, e.(T))
	}
	return s
}

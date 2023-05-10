package main

import (
	"context"
)

type Manager struct {
	events []interface{}
}

func NewManager() *Manager {
	return &Manager{}
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

func (m *Manager) Push(event interface{}) {
	m.events = append(m.events, event)
}

func (m *Manager) Events() []interface{} {
	return m.events
}

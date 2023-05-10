package main

type EventType int

const (
	None EventType = iota
	LevelUp
	QuestClear
)

type Event interface {
	EventType() EventType
}

type LevelUpEvent struct {
	Level int
}

func (e LevelUpEvent) EventType() EventType {
	return LevelUp
}

type QuestClearEvent struct {
	QuestID   int
	ClearRank string
}

func (e QuestClearEvent) EventType() EventType {
	return QuestClear
}

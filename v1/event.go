package main

type LevelUpEvent struct {
	Level int
}

type QuestClearEvent struct {
	QuestID   int
	ClearRank string
}

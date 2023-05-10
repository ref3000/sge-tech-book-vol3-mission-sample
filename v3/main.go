package main

import (
	"context"
	"fmt"
)

func main() {
	manager := NewManager()
	ctx := SetManager(context.Background(), manager)
	doBusinessLogic(ctx)
	progressMission(ctx)
}

func doBusinessLogic(ctx context.Context) {
	manager := GetManager(ctx)
	manager.Push(QuestClearEvent{QuestID: 123, ClearRank: "A"})
	manager.Push(LevelUpEvent{Level: 10})
	manager.Push(QuestClearEvent{QuestID: 321, ClearRank: "S"})
}

func progressMission(ctx context.Context) {
	for _, v := range Events[LevelUpEvent](ctx) {
		fmt.Printf("%T, %d\n", v, v.Level)
	}
	for _, v := range Events[QuestClearEvent](ctx) {
		fmt.Printf("%T, %d, %s\n", v, v.QuestID, v.ClearRank)
	}
}

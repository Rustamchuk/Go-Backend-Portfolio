package model

type User struct {
	Id      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Balance int    `json:"balance" db:"balance"`
}

type Quest struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Cost int    `json:"cost" db:"cost"`
}

type CompletedQuest struct {
	Id      int `json:"id" db:"id"`
	UserID  int `json:"user_id" db:"user_id"`
	QuestID int `json:"quest_id" db:"quest_id"`
}

type Achievements struct {
	Balance int
	Quests  []Quest
}

package service

import (
	"VK-Quest/internal/model"
	"VK-Quest/internal/repository"
	"context"
)

type QuestServiceImplementation struct {
	repo repository.QuestRepo
}

func NewQuestServiceImplementation(repo repository.QuestRepo) *QuestServiceImplementation {
	return &QuestServiceImplementation{repo: repo}
}

func (q *QuestServiceImplementation) CreateQuest(ctx context.Context, quest *model.Quest) error {
	return q.repo.CreateQuest(ctx, quest)
}

func (q *QuestServiceImplementation) GetQuest(ctx context.Context, id int) (*model.Quest, error) {
	return q.repo.GetQuest(ctx, id)
}

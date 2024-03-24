package service

import (
	"VK-Quest/internal/model"
	"VK-Quest/internal/repository"
	"context"
)

type CompleteQuestServiceImplementation struct {
	repo repository.CompletedQuestRepo
}

func NewCompleteQuestServiceImplementation(repo repository.CompletedQuestRepo) *CompleteQuestServiceImplementation {
	return &CompleteQuestServiceImplementation{repo: repo}
}

func (c *CompleteQuestServiceImplementation) SetQuestComplete(ctx context.Context, completeQuest *model.CompletedQuest) error {
	return c.repo.AddCompleteQuest(ctx, completeQuest)
}

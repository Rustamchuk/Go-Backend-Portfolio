package service

import (
	"VK-Quest/internal/model"
	"VK-Quest/internal/repository"
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, user *model.User) error
	IncreaseUserBalance(ctx context.Context, id, value int) error
	GetUser(ctx context.Context, id int) (*model.User, error)
	GetAchievements(ctx context.Context, id int) (*model.Achievements, error)
}

type QuestService interface {
	CreateQuest(ctx context.Context, quest *model.Quest) error
	GetQuest(ctx context.Context, id int) (*model.Quest, error)
}

type CompletedQuestService interface {
	SetQuestComplete(ctx context.Context, completeQuest *model.CompletedQuest) error
}

type Service struct {
	UserService
	QuestService
	CompletedQuestService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UserService:           NewUserServiceImplementation(repos.UserRepo),
		QuestService:          NewQuestServiceImplementation(repos.QuestRepo),
		CompletedQuestService: NewCompleteQuestServiceImplementation(repos.CompletedQuestRepo),
	}
}

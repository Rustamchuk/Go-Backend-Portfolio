package repository

import (
	"VK-Quest/internal/model"
	"VK-Quest/internal/repository/postgres"
	"context"
	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *model.User) error
	IncreaseUserBalance(ctx context.Context, id, value int) error
	GetUser(ctx context.Context, id int) (*model.User, error)
	GetAchievements(ctx context.Context, id int) (*model.Achievements, error)
}

type QuestRepo interface {
	CreateQuest(ctx context.Context, quest *model.Quest) error
	GetQuest(ctx context.Context, id int) (*model.Quest, error)
}

type CompletedQuestRepo interface {
	AddCompleteQuest(ctx context.Context, completeQuest *model.CompletedQuest) error
}

type Repository struct {
	UserRepo
	QuestRepo
	CompletedQuestRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepo:           postgres.NewUserPostgres(db),
		QuestRepo:          postgres.NewQuestPostgres(db),
		CompletedQuestRepo: postgres.NewCompleteQuestPostgres(db),
	}
}

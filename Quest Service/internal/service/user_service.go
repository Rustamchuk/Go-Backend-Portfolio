package service

import (
	"VK-Quest/internal/model"
	"VK-Quest/internal/repository"
	"context"
)

type UserServiceImplementation struct {
	repo repository.UserRepo
}

func NewUserServiceImplementation(repo repository.UserRepo) *UserServiceImplementation {
	return &UserServiceImplementation{repo: repo}
}

func (u *UserServiceImplementation) IncreaseUserBalance(ctx context.Context, id, value int) error {
	return u.repo.IncreaseUserBalance(ctx, id, value)
}

func (u *UserServiceImplementation) CreateUser(ctx context.Context, user *model.User) error {
	return u.repo.CreateUser(ctx, user)
}

func (u *UserServiceImplementation) GetUser(ctx context.Context, id int) (*model.User, error) {
	return u.repo.GetUser(ctx, id)
}

func (u *UserServiceImplementation) GetAchievements(ctx context.Context, id int) (*model.Achievements, error) {
	return u.repo.GetAchievements(ctx, id)
}

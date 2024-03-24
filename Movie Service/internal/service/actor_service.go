package service

import (
	"VK-Test_Ex/internal/model"
	"VK-Test_Ex/internal/repository"
	"context"
)

type ActorServiceImplementation struct {
	repo repository.ActorRepo
}

func NewActorServiceImplementation(repo repository.ActorRepo) *ActorServiceImplementation {
	return &ActorServiceImplementation{repo: repo}
}

func (a *ActorServiceImplementation) CreateActor(ctx context.Context, actor *model.Actor) error {
	return a.repo.CreateActor(ctx, actor)
}

func (a *ActorServiceImplementation) UpdateActor(ctx context.Context, actor *model.Actor) error {
	return a.repo.UpdateActor(ctx, actor)
}

func (a *ActorServiceImplementation) DeleteActor(ctx context.Context, id int64) error {
	return a.repo.DeleteActor(ctx, id)
}

func (a *ActorServiceImplementation) GetActors(ctx context.Context) ([]model.Actor, error) {
	return a.repo.GetActors(ctx)
}

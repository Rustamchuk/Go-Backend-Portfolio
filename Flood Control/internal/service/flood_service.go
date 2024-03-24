package service

import (
	"context"
	"task/internal/model"
	"task/internal/repository"
	"time"
)

type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}

type FloodControlService struct {
	maxResponseCnt int
	intervalSize   int
	database       repository.FloodControlRepo
}

func NewFloodControlService(maxResponseCnt, intervalSize int, database repository.FloodControlRepo) FloodControlService {
	return FloodControlService{
		maxResponseCnt: maxResponseCnt,
		intervalSize:   intervalSize,
		database:       database,
	}
}

func (f *FloodControlService) Check(ctx context.Context, userID int64) (bool, error) {
	currentTime := time.Now()
	intervalStart := currentTime.Add(-time.Duration(f.intervalSize) * time.Second)

	err := f.database.AddResponse(ctx, model.RequestStat{UserId: userID, Time: currentTime})
	if err != nil {
		return false, err
	}

	count, err := f.database.GetStatWithInterval(ctx, model.RequestStat{UserId: userID, Time: intervalStart})
	if err != nil {
		return false, err
	}
	return count <= f.maxResponseCnt, nil
}

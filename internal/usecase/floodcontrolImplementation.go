package usecase

import (
	"FloodControl/internal/model"
	"FloodControl/internal/repository"
	"context"
	"math"
	"strconv"
	"time"
)

type FloodControlImplementation struct {
	Repository   *repository.Repository
	FloodControl model.FloodControlParameters
}

func NewQuestUseCaseImplementation(repository *repository.Repository, floodControl model.FloodControlParameters) *FloodControlImplementation {
	return &FloodControlImplementation{
		Repository:   repository,
		FloodControl: floodControl,
	}
}

func (fl *FloodControlImplementation) Check(ctx context.Context, userID int64) (bool, error) {
	currentTime := time.Now()
	id := strconv.FormatInt(userID, 10)
	err := fl.AddRequest(ctx, id, currentTime)
	if err != nil {
		return false, err
	}

	err = fl.DeletePreviousRequests(ctx, id, fl.FloodControl.Interval)
	if err != nil {
		return false, err
	}

	amount, err := fl.GetAmountOfRequests(ctx, id)
	if err != nil {
		return false, err
	}

	if amount >= fl.FloodControl.Limit {
		return false, nil
	}
	return true, nil

}

func (fl *FloodControlImplementation) GetAmountOfRequests(ctx context.Context, userID string) (int64, error) {

	amount, err := fl.Repository.GetAmountOfRequests(ctx, userID)
	if err != nil {
		return math.MaxInt64, err
	}

	return amount, nil
}

func (fl *FloodControlImplementation) AddRequest(ctx context.Context, userID string, moment time.Time) error {

	err := fl.Repository.AddRequest(ctx, userID, moment)
	if err != nil {
		return err
	}

	return nil
}

func (fl *FloodControlImplementation) DeletePreviousRequests(ctx context.Context, userID string, interval time.Duration) error {
	err := fl.Repository.DeletePreviousRequests(ctx, userID, interval)
	if err != nil {
		return err
	}

	return nil
}

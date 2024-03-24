package repository

import (
	"context"
	"time"
)

type RequestRepository interface {
	GetAmountOfRequests(ctx context.Context, userID string) (int64, error)
	AddRequest(ctx context.Context, userID string, currentTime time.Time) error
	DeletePreviousRequests(ctx context.Context, userID string, interval time.Duration) error
}

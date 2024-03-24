package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"math"
	"strconv"
	"time"
)

type RequestRedisRepository struct {
	client *redis.Client
}

func NewRequestRedisRepository(client *redis.Client) *RequestRedisRepository {
	return &RequestRedisRepository{client: client}
}

func (r *RequestRedisRepository) GetAmountOfRequests(ctx context.Context, userID string) (count int64, err error) {
	requests, err := r.client.LRange(ctx, userID, 0, -1).Result()
	if err != nil {
		return math.MaxInt64, nil
	}
	return int64(len(requests)), nil
}

func (r *RequestRedisRepository) AddRequest(ctx context.Context, userID string, currentTime time.Time) error {
	err := r.client.LPush(ctx, userID, currentTime.UnixMilli()).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RequestRedisRepository) DeletePreviousRequests(ctx context.Context, userID string, interval time.Duration) error {

	requests, err := r.client.LRange(ctx, userID, 0, -1).Result()
	if err != nil {
		return err
	}

	var index int64 = 0
	currentTimeFromRequests := requests[0]
	currentTimeInt, _ := strconv.ParseInt(currentTimeFromRequests, 10, 64)
	intervalStr := time.UnixMilli(interval.Milliseconds())
	formattedInterval := fmt.Sprintf("%d", intervalStr.UnixMilli())
	intervalInt, _ := strconv.ParseInt(formattedInterval, 10, 64)

	for _, t := range requests {
		tInt, _ := strconv.ParseInt(t, 10, 64)
		if currentTimeInt-tInt >= intervalInt {
			index++
		}
	}

	if index != 0 {
		err = r.client.LTrim(ctx, userID, index, -1).Err()
	}

	if err != nil {
		return err
	}
	return nil
}

package repository

import "github.com/redis/go-redis/v9"

type Repository struct {
	RequestRepository
}

func NewRepository(client *redis.Client) *Repository {
	return &Repository{
		RequestRepository: NewRequestRedisRepository(client),
	}
}

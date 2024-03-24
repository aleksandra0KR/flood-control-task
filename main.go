package main

import (
	"FloodControl/configs"
	"FloodControl/internal/model"
	"FloodControl/internal/repository"
	"FloodControl/internal/usecase"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func main() {

	config, err := configs.InitConfig("config.yml")
	if err != nil {
		log.Fatalf("initialization of configs have failed %s", err.Error())
	}

	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host + ":" + config.Redis.Port,
		Password: "",
		DB:       0,
		Protocol: 3,
	})

	foodControlParam := model.FloodControlParameters{Interval: config.Interval, Limit: config.Limit}
	repo := repository.NewRepository(client)
	var floodControlImplementation FloodControl = usecase.NewQuestUseCaseImplementation(repo, foodControlParam)
	ctx := context.Background()

	for i := 0; i < 3; i++ {

		result, err := floodControlImplementation.Check(ctx, config.UserID)
		if err != nil {
			fmt.Println("failed to check", err)
		}

		if result {
			fmt.Println("successful pass")
		} else {
			fmt.Println("doesn't pass")
		}

		time.Sleep(time.Second)
	}

	for i := 0; i < 3; i++ {

		result, err := floodControlImplementation.Check(ctx, config.UserID)
		if err != nil {
			fmt.Println("failed to check", err)
		}

		if result {
			fmt.Println("successful pass")
		} else {
			fmt.Println("doesn't pass")
		}

		time.Sleep(time.Second * 6)
	}
}

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}

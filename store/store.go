package store

import (
	"log"

	"github.com/go-redis/redis"
)

type Service struct {
	client *redis.Client
}

type Input struct {
	RedisURL string
}

func New(input *Input) *Service {
	if input == nil {
		log.Fatal("input is required")
	}

	client := redis.NewClient(&redis.Options{
		Addr: input.RedisURL
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	return &Service{
		client: client,
	}
}

func (s *Service) Set(key string, valut interface{}) error {
	exp := time.Duration(600 * time.Second) // 10 minutes
	return s.client.Set(key, value, exp).Err()
}

func (s *Service) Get(key string) (interface{}, error) {
	value, err := s.client.Get(key).Result()
	if err != nil {
		return nil, err
	}

	return value, nil
}
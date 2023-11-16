package storage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Storage struct {
	clientRedis *redis.Client
}

func (s *Storage) GetRedis(key string) (string, error) {
	val, err := s.clientRedis.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (s *Storage) SaveRedis(key string, value string) error {
	return s.clientRedis.Set(ctx, key, value, 168*time.Hour).Err()
}

func NewStorage(cr *redis.Client) *Storage {
	return &Storage{
		clientRedis: cr,
	}
}

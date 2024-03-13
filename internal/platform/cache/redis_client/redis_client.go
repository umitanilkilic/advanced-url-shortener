package redis_client

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
	"github.com/umitanilkilic/advanced-url-shortener/internal/model"
)

type RedisServerConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(ctx context.Context, cfg *RedisServerConfig) (*RedisClient, error) {

	rClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	pong, err := rClient.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("fail to connect to redis: %v", err)
	}
	fmt.Println(pong)

	return &RedisClient{client: rClient}, nil
}
func (c *RedisClient) Close() error {
	return c.client.Close()
}

func (c *RedisClient) SaveMapping(ctx context.Context, urlStruct *model.ShortURL) error {

	shortUrlJson, err := json.Marshal(urlStruct)
	if err != nil {
		return err
	}

	err = c.client.Set(ctx, strconv.Itoa(urlStruct.ID), shortUrlJson, 0).Err()

	if err != nil {
		return fmt.Errorf("fail cannot save: %v, ID: %v, originalUrl: %v", err, urlStruct.ID, urlStruct.Long)
	}
	return nil
}

func (c *RedisClient) RetrieveLongUrl(ctx context.Context, shortUrlID string) (model.ShortURL, error) {
	result, err := c.client.Get(ctx, shortUrlID).Result()

	if err != nil {
		return model.ShortURL{}, fmt.Errorf("fail to retrieve long url: %v", err)
	}

	var model model.ShortURL = model.ShortURL{}
	err = json.Unmarshal([]byte(result), &model)
	if err != nil {
		return model, fmt.Errorf("fail to unmarshal: %v", err)
	}

	return model, nil
}

package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/umitanilkilic/advanced-url-shortener/model"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(ctx context.Context, connectionString string) (*RedisClient, error) {
	opt, err := redis.ParseURL(connectionString)
	if err != nil {
		return nil, fmt.Errorf("fail to parse redis url: %v", err)
	}

	rClient := redis.NewClient(opt)

	_, err = rClient.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("fail to connect to redis: %v", err)
	}

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

	err = c.client.Set(ctx, (urlStruct.UrlID), shortUrlJson, 0).Err()

	if err != nil {
		return fmt.Errorf("fail cannot save: %v, ID: %v, originalUrl: %v", err, urlStruct.UrlID, urlStruct.Long)
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

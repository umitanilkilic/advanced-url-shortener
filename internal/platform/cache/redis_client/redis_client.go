package redis_client

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/umitanilkilic/advanced-url-shortener/internal/model"
)

type CacheService struct {
	redisClient *redis.Client
}

var cacheService = &CacheService{}
var ctx = context.Background()

func ConnectToRedis(redisAddress string, redisPort string, redisPassword string, dbNo int) (*CacheService, error) {

	rClient := redis.NewClient(&redis.Options{
		Addr:     redisAddress + ":" + redisPort,
		Password: redisPassword,
		DB:       dbNo,
	})

	pong, err := rClient.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("fail to connect to redis: %v", err)
	}
	fmt.Println(pong)

	cacheService.redisClient = rClient
	return cacheService, nil
}

func SaveMapping(urlStruct model.ShortURL) error {

	///err := cacheService.redisClient.Set(ctx, shortUrl, originalUrl, 0).Err()
	data := map[string]interface{}{
		"ID":         urlStruct.ID,
		"LongUrl":    urlStruct.Long,
		"CreateDate": urlStruct.CreatedAt,
	}
	err := cacheService.redisClient.HMSet(ctx, "sUrl", data).Err()

	if err != nil {
		return fmt.Errorf("fail cannot save: %v, ID: %v, originalUrl: %v", err, urlStruct.ID, urlStruct.Long)
	}
	return nil
}

func RetrieveLongUrl(shortUrl string) (string, error) {
	result, err := cacheService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		return "", fmt.Errorf("fail to retrieve long url: %v", err)
	}
	return result, nil
}

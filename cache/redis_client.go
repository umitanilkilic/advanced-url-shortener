package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
	"github.com/umitanilkilic/advanced-url-shortener/config"
	"github.com/umitanilkilic/advanced-url-shortener/model"
)

var Client *redis.Client

var ctx = context.Background()

func InitializeRedisClient() error {
	connectionString := (*config.Config)["REDIS_URL"]

	opt, err := redis.ParseURL(connectionString)
	if err != nil {
		return fmt.Errorf("fail to parse redis url: %v", err)
	}

	rClient := redis.NewClient(opt)

	_, err = rClient.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("fail to connect to redis: %v", err)
	}
	Client = rClient
	return nil
}

func Close() error {
	return Client.Close()
}

func SaveMapping(urlStruct *model.ShortURL) error {

	shortUrlJson, err := json.Marshal(urlStruct)
	if err != nil {
		return err
	}

	err = Client.Set(ctx, strconv.Itoa(int(urlStruct.UrlID)), shortUrlJson, 0).Err()

	if err != nil {
		return fmt.Errorf("fail cannot save: %v, ID: %v, originalUrl: %v", err, urlStruct.UrlID, urlStruct.Long)
	}
	return nil
}

func RetrieveLongUrl(shortUrlID string) (model.ShortURL, error) {
	result, err := Client.Get(ctx, shortUrlID).Result()

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

package helper

import (
	"context"
	"fmt"
	"log"

	"github.com/umitanilkilic/advanced-url-shortener/internal/model"
	"github.com/umitanilkilic/advanced-url-shortener/internal/platform/cache/redis_client"
	"github.com/umitanilkilic/advanced-url-shortener/internal/platform/database/pg_client"
)

var rd *redis_client.RedisClient

var pg *pg_client.PostgresClient

var defaultContext context.Context = context.Background()

/* type Store interface {
	SaveMapping(ctx context.Context, urlStruct *model.ShortURL) error
	RetrieveLongUrl(ctx context.Context, shortUrlID string) (model.ShortURL, error)
	Close() error
} */

func RunDatabases(rdConnStr string, pgConnStr string) error {
	var err error
	rd, err = redis_client.NewRedisClient(defaultContext, rdConnStr)
	if err != nil {
		return fmt.Errorf("error while creating redis client: %v", err)
	}
	fmt.Println("redis client created")

	pg, err = pg_client.NewPostgresClient(defaultContext, pgConnStr)
	if err != nil {
		return fmt.Errorf("error while creating postgres client: %v", err)
	}

	return nil
}

func RetrieveLongUrl(shortUrlID *string) (string, error) {
	url, err := rd.RetrieveLongUrl(defaultContext, *shortUrlID)
	if err != nil {
		log.Println("url not found on redis: ", err)
		url, err = pg.RetrieveLongUrl(defaultContext, *shortUrlID)
		if err != nil {
			return "", fmt.Errorf("url not found on postgres & redis: %v", err)
		}
		rd.SaveMapping(defaultContext, &url)
	}
	return url.Long, nil
}

func StoreUrl(url model.ShortURL) error {
	err := pg.SaveMapping(defaultContext, &url)
	if err != nil {
		return fmt.Errorf("error while saving to postgres: %v", err)
	}
	err = rd.SaveMapping(defaultContext, &url)
	if err != nil {
		return fmt.Errorf("error while saving to redis: %v", err)
	}
	return nil
}

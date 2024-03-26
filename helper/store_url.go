package helper

import (
	"context"
	"fmt"
	"log"

	"github.com/umitanilkilic/advanced-url-shortener/cache"
	"github.com/umitanilkilic/advanced-url-shortener/database"
	"gorm.io/gorm"
)

var rd *cache.RedisClient

var pg *gorm.DB

var defaultContext context.Context = context.Background()

/* type Store interface {
	SaveMapping(ctx context.Context, urlStruct *model.ShortURL) error
	RetrieveLongUrl(ctx context.Context, shortUrlID string) (model.ShortURL, error)
	Close() error
} */

func RunDatabases(rdConnStr string, pgConnStr string) error {
	var err error
	rd, err = cache.NewRedisClient(defaultContext, rdConnStr)
	if err != nil {
		return fmt.Errorf("error while creating redis client: %v", err)
	}
	fmt.Println("redis client created")

	pg = database.DB

	return nil
}

func RetrieveLongUrl(shortUrlID *string) (string, error) {
	url, err := rd.RetrieveLongUrl(defaultContext, *shortUrlID)
	if err != nil {
		log.Println("url not found on redis: ", err)
		url, err = database.RetrieveLongUrl(*shortUrlID)
		if err != nil {
			return "", fmt.Errorf("url not found on postgres & redis: %v", err)
		}
		rd.SaveMapping(defaultContext, &url)
	}
	return url.Long, nil
}

/* func StoreUrl(url model.ShortURL) error {
	err := pg.SaveMapping(defaultContext, &url)
	if err != nil {
		return fmt.Errorf("error while saving to postgres: %v", err)
	}
	err = rd.SaveMapping(defaultContext, &url)
	if err != nil {
		return fmt.Errorf("error while saving to redis: %v", err)
	}
	return nil
} */

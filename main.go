package main

import (
	"log"
	"os"
	"strings"

	"github.com/umitanilkilic/advanced-url-shortener/cmd/urlshortener"
	"github.com/umitanilkilic/advanced-url-shortener/internal/helper"
)

/* var rd *redis_client.RedisClient
var pg *pg_client.PostgresClient

var ctx = context.Background() */

func main() {
	cfg := make(map[string]string)
	for _, e := range os.Environ() {
		if i := strings.Index(e, "="); i >= 0 {
			cfg[e[:i]] = e[i+1:]
		}
	}

	err := helper.RunDatabases(cfg["REDIS_CONNECTION_STRING"], cfg["POSTGRES_CONNECTION_STRING"])
	if err != nil {
		log.Fatal(err)
	}

	err = urlshortener.RunUrlShortener(cfg)
	if err != nil {
		log.Fatal(err)
	}

	/* 	godotenv.Load()
	   	err := helper.RunDatabases(os.Getenv("REDIS_CONNECTION_STRING"), os.Getenv("POSTGRES_CONNECTION_STRING"))
	   	if err != nil {
	   		panic(err)
	   	}
	   	_ = urlshortener.RunUrlShortener(map[string]string{
	   		"MAX_RATE_LIMIT":        os.Getenv("MAX_RATE_LIMIT"),
	   		"RATE_LIMIT_EXPIRATION": os.Getenv("RATE_LIMIT_EXPIRATION"),
	   		"APP_ADDRESS":           os.Getenv("APP_ADDRESS"),
	   		"APP_PORT":              os.Getenv("APP_PORT"),
	   	}) */
}

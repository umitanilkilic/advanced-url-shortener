package main

import (
	"context"

	"github.com/umitanilkilic/advanced-url-shortener/cmd/urlshortener"
	"github.com/umitanilkilic/advanced-url-shortener/internal/config"
	"github.com/umitanilkilic/advanced-url-shortener/internal/helper"
)

func init() {

}

func main() {
	dbConfig, err := config.LoadDatabaseConfig()
	if err != nil {
		panic("Error loading database config")
	}

	appConfig, err := config.AppConfig()
	if err != nil {
		panic("Error loading app config")
	}

	// Define a context for the application
	ctx := context.TODO()

	helper.StartStoreServices(ctx, *dbConfig)

	urlshortener.RunUrlShortener(*appConfig)
}

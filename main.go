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
	cfg, _ := config.LoadDatabaseConfig()

	// Define a context for the application
	ctx := context.TODO()

	helper.StartStoreServices(ctx, *cfg)

	urlshortener.RunUrlShortener()
}

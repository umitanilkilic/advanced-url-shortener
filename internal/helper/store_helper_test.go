package helper_test

import (
	"context"
	"testing"

	"github.com/umitanilkilic/advanced-url-shortener/internal/config"
	"github.com/umitanilkilic/advanced-url-shortener/internal/helper"
	"github.com/umitanilkilic/advanced-url-shortener/internal/platform/cache/redis_client"
	"github.com/umitanilkilic/advanced-url-shortener/internal/platform/database/pg_client"
)

func TestStartStoreServices(t *testing.T) {
	// Initialize the necessary dependencies
	//rd := &redis_client.RedisClient{}
	//pg := &pg_client.PostgresClient{}
	defaultContext := context.Background()
	cfg := config.DatabaseConfig{}

	err := helper.StartStoreServices(defaultContext, cfg)
	if err != nil {
		t.Errorf("StartStoreServices returned an error: %v", err)
	}

	// Add more test cases as needed
}

func TestRetrieveLongUrl(t *testing.T) {
	// Initialize the necessary dependencies
	rd := &redis_client.RedisClient{}
	pg := &pg_client.PostgresClient{}
	defaultContext := context.Background()

	// Mock the short URL ID
	shortUrlID := "abc123"

	longUrl, err := helper.RetrieveLongUrl(defaultContext, shortUrlID, rd, pg)
	if err != nil {
		t.Errorf("RetrieveLongUrl returned an error: %v", err)
	}

	// Add assertions to validate the result
	if longUrl != "expected_long_url" {
		t.Errorf("RetrieveLongUrl returned incorrect long URL: got %s, want %s", longUrl, "expected_long_url")
	}

	// Add more test cases as needed
}

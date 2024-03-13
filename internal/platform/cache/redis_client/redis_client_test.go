package redis_client

import (
	"context"
	"testing"

	"github.com/umitanilkilic/advanced-url-shortener/internal/model"
)

var testData = model.ShortURL{
	ID:        125,
	Long:      "https://www.google.com",
	CreatedAt: "2021-07-01 00:00:00",
}

func TestConnection(t *testing.T) {
	ctx := context.Background()
	_, err := NewRedisClient(ctx, "localhost", "6379", "adam1234", 0)
	if err != nil {
		t.Errorf("ConnectToRedis failed: %v", err)
	}
}
func TestSaveMapping(t *testing.T) {
	ctx := context.Background()
	redis, err := NewRedisClient(ctx, "localhost", "6379", "adam1234", 0)
	if err != nil {
		t.Errorf("ConnectToRedis failed: %v", err)
	}
	err = redis.SaveMapping(ctx, &testData)
	if err != nil {
		t.Errorf("SaveMapping failed: %v", err)
	}
}
func TestRetrieveLongUrl(t *testing.T) {
	ctx := context.Background()
	redis, err := NewRedisClient(ctx, "localhost", "6379", "adam1234", 0)
	if err != nil {
		t.Errorf("ConnectToRedis failed: %v", err)
	}
	_, err = redis.RetrieveLongUrl(ctx, "125")
	if err != nil {
		t.Errorf("RetrieveLongUrl failed: %v", err)
	}
}

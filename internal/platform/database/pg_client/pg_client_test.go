package pg_client

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
	pg, err := NewPostgresClient(ctx, "localhost", "5432", "postgres", "adam1234", "shortened_urls", "disable")
	if err != nil {
		t.Errorf("ConnectToPostgres failed: %v", err)
	}
	err = pg.checkDatabaseConnection(ctx)
	if err != nil {
		t.Errorf("checkDatabaseConnection failed: %v", err)
	}

}
func TestSaveMapping(t *testing.T) {
	ctx := context.Background()
	pg, err := NewPostgresClient(ctx, "localhost", "5432", "postgres", "adam1234", "shortened_urls", "disable")
	if err != nil {
		t.Errorf("ConnectToPostgres failed: %v", err)
	}

	err = pg.SaveMapping(ctx, &testData)
	if err != nil {
		t.Errorf("SaveMapping failed: %v", err)
	}
}
func TestRetrieveLongUrl(t *testing.T) {
	ctx := context.Background()
	pg, err := NewPostgresClient(ctx, "localhost", "5432", "postgres", "adam1234", "shortened_urls", "disable")
	if err != nil {
		t.Errorf("ConnectToPostgres failed: %v", err)
	}
	_, err = pg.RetrieveLongUrl(ctx, "125")
	if err != nil {
		t.Errorf("RetrieveLongUrl failed: %v", err)
	}
}

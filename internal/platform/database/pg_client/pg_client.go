package pg_client

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/umitanilkilic/advanced-url-shortener/internal/model"
)

type PostgresClient struct {
	db *sqlx.DB
}

func NewPostgresClient(ctx context.Context, connectionString string) (*PostgresClient, error) {
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("could not connect to Postgres: %w", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	client := &PostgresClient{db: db}
	return client, nil
}

func (c *PostgresClient) Close() error {
	return c.db.Close()
}

func (p *PostgresClient) SaveMapping(ctx context.Context, urlStruct *model.ShortURL) error {

	insertStatement := "INSERT INTO shorturl (url_id ,long_url, created_at) VALUES ($1, $2, $3)"
	_, err := p.db.ExecContext(ctx, insertStatement, urlStruct.UrlID, urlStruct.Long, urlStruct.CreatedAt)

	return err
}

func (p *PostgresClient) RetrieveLongUrl(ctx context.Context, shortUrlID string) (model.ShortURL, error) {
	if err := p.checkDatabaseConnection(ctx); err != nil {
		return model.ShortURL{}, err
	}

	var model model.ShortURL
	queryStatement := "SELECT url_id, long_url, created_at FROM public.shorturl WHERE url_id = $1"
	err := p.db.Get(&model, queryStatement, shortUrlID)
	return model, err
}

func (c *PostgresClient) checkDatabaseConnection(ctx context.Context) error {
	if err := c.db.PingContext(ctx); err != nil {
		return fmt.Errorf("database connection lost: %w", err)
	}
	return nil
}

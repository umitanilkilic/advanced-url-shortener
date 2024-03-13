package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/umitanilkilic/advanced-url-shortener/internal/platform/cache/redis_client"
	"github.com/umitanilkilic/advanced-url-shortener/internal/platform/database/pg_client"
)

type DatabaseConfig struct {
	RedisServerConfig    redis_client.RedisServerConfig
	PostgresServerConfig pg_client.PostgresServerConfig
}

type FiberConfig struct {
	AppAddress          string
	AppPort             string
	RateLimitMax        int
	RateLimitExpiration int
}

func LoadDatabaseConfig() (*DatabaseConfig, error) {
	err := godotenv.Load("database.env")
	if err != nil {
		panic("Error loading database.env file")
	}

	redisServerConfig := redis_client.RedisServerConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}

	postgresServerConfig := pg_client.PostgresServerConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DB:       os.Getenv("POSTGRES_DB"),
		SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
		//Table:    os.Getenv("POSTGRES_TABLE"),
	}
	return &DatabaseConfig{RedisServerConfig: redisServerConfig, PostgresServerConfig: postgresServerConfig}, nil
}

func AppConfig() (*FiberConfig, error) {
	err := godotenv.Load("app.env")
	if err != nil {
		panic("Error loading ratelimiter.env file")
	}
	rateLimiterConfig := FiberConfig{AppAddress: os.Getenv("APP_ADDRESS"), AppPort: os.Getenv("APP_PORT")}

	max, err := strconv.Atoi(os.Getenv("RATE_LIMIT_MAX"))

	if err != nil {
		panic("Error Rate Limiter Max is not a number")
	}
	rateLimiterConfig.RateLimitMax = max

	expiration, err := strconv.Atoi(os.Getenv("RATE_LIMIT_EXPIRATION"))

	if err != nil {
		panic("Error Rate Limiter Expiration is not a number")
	}
	rateLimiterConfig.RateLimitExpiration = expiration

	return &rateLimiterConfig, nil
}

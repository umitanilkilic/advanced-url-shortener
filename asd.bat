@echo off
set REDIS_CONNECTION_STRING=redis://:adam1234@localhost:6379/0
set POSTGRES_CONNECTION_STRING=postgresql://postgres:adam1234@localhost:5432/shortened_urls?sslmode=disable
set MAX_RATE_LIMIT=10
set RATE_LIMIT_EXPIRATION=60
set APP_ADDRESS=localhost
set APP_PORT=8080
advanced-url-shortener.exe

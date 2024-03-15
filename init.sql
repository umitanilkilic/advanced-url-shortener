CREATE TABLE IF NOT EXISTS public.shorturl (
    id SERIAL PRIMARY KEY,
    url_id VARCHAR(255),
    long_url VARCHAR(255),
    created_at TIMESTAMP WITHOUT TIME ZONE
);

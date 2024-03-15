CREATE TABLE IF NOT EXISTS public.shorturl (
    url_id SERIAL PRIMARY KEY,
    long_url VARCHAR(255),
    created_at TIMESTAMP WITHOUT TIME ZONE
);

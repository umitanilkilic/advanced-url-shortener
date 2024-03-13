package model

type ShortURL struct {
	ID        int    `json:"id" db:"url_id"`
	Long      string `json:"long_url" db:"long_url"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

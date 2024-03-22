package model

type ShortURL struct {
	Active    bool   `json:"active" db:"active"`
	UrlID     string `json:"url_id" db:"url_id"`
	Name      string `json:"name" db:"name"`
	Long      string `json:"long_url" db:"long_url"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
	ExpiresAt string `json:"expires_at" db:"expires_at"`
}

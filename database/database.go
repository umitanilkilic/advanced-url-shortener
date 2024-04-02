package database

import (
	"errors"

	"github.com/umitanilkilic/advanced-url-shortener/config"
	"github.com/umitanilkilic/advanced-url-shortener/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabaseConnection() error {
	var err error
	connStr := (*config.Config)["DATABASE_URL"]
	DB, err = gorm.Open(postgres.Open(connStr))
	if err != nil {
		return errors.New("failed to connect database")
	}
	err = ExecuteDatabaseMigrations()
	if err != nil {
		return errors.New("failed to execute database migrations")
	}
	return nil
}

func ExecuteDatabaseMigrations() error {
	return DB.AutoMigrate(&model.User{}, &model.ShortURL{}, &model.UrlStats{})
}

func RetrieveLongUrl(shortUrlID string) (model.ShortURL, error) {
	var url model.ShortURL
	result := DB.Where("url_id = ?", shortUrlID).First(&url)
	if result.Error != nil {
		return model.ShortURL{}, result.Error
	}
	return url, nil
}

func SaveMapping(urlStruct *model.ShortURL) error {
	result := DB.Create(urlStruct)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

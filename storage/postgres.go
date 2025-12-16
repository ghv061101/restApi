package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
	URL      string // optional full connection URL (takes precedence if set)
}

func NewConnection(config *Config) (*gorm.DB, error) {
	// If a full DATABASE_URL/URL is provided, use that
	if config.URL != "" {
		db, err := gorm.Open(postgres.Open(config.URL), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return db, nil
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
		config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

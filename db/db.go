package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB_Settings struct {
	Host     string
	Port     uint
	User     string
	Password string
	Name     string
	SSL      string
}

func CreateDB(settings DB_Settings) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", settings.Host, settings.User, settings.Password, settings.Name, settings.Port, settings.SSL)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Shortener{})

	return db, nil
}

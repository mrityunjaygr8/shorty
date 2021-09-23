package app

import (
	"gorm.io/gorm"

	"github.com/mrityunjaygr8/shorty/db"
)

type App struct {
	DB *gorm.DB
}

type Config struct {
	DB_NAME string
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT uint
	DB_SSL  string
}

func Setup(config Config) App {
	app := App{}

	database, err := db.CreateDB(db.DB_Settings{Host: config.DB_HOST, Port: config.DB_PORT, User: config.DB_USER, Password: config.DB_PASS, Name: config.DB_NAME, SSL: config.DB_SSL})
	if err != nil {
		panic("Error connecting to database")
	}

	app.DB = database
	return app
}

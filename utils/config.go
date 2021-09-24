package utils

import (
	"fmt"

	"github.com/mrityunjaygr8/shorty/app"
	"github.com/spf13/viper"
)

func GetConfig() (app.Config, error) {
	config := app.Config{}
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		return app.Config{}, err
	}

	return config, nil
}

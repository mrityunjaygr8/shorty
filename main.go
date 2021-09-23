package main

import (
	"fmt"

	"github.com/mrityunjaygr8/shorty/app"
)

func main() {
	app := app.Setup(app.Config{DB_NAME: "postgres", DB_USER: "root", DB_PASS: "secret", DB_HOST: "localhost", DB_PORT: 5432, DB_SSL: "disable"})
	fmt.Println(app)

	token, err := app.Create("www.google.com")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(token)
}

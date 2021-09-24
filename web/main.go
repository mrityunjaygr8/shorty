package main

import (
	"fmt"

	"github.com/mrityunjaygr8/shorty/utils"
)

func main() {
	fmt.Println("YO")
	config, err := utils.GetConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println(config)

}

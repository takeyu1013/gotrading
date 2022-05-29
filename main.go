package main

import (
	"fmt"

	"github.com/takeyu1013/gotrading/config"
)

func main() {
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)
}

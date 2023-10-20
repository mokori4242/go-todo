package main

import (
	"fmt"
	"app/config"
)

func main() {
	fmt.Println(config.Config.AppPort)
	fmt.Println(config.Config.DBUser)
	fmt.Println(config.Config.DBPassword)
	fmt.Println(config.Config.DBHost)
	fmt.Println(config.Config.DBPort)
	fmt.Println(config.Config.LogFile)
}
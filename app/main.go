package main

import (
	"fmt"
	// "app/config"
	"app/models"
	// "log"
)

func main() {
	// fmt.Println(config.Config.AppPort)
	// fmt.Println(config.Config.DBUser)
	// fmt.Println(config.Config.DBPassword)
	// fmt.Println(config.Config.DBHost)
	// fmt.Println(config.Config.DBPort)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	// fmt.Println(models.Db)

	u := &models.User{
		Name: "test",
		Email: "test@example.com",
		Password: "password",
	}
	fmt.Println(u)

	u.CreateUser()
}
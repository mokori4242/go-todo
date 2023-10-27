package main

import (
	"fmt"
	// "app/config"
	"app/models"
	// "log"
	"context"
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

	// u := &models.User{
	// 	Name: "test",
	// 	Email: "test@example.com",
	// 	Password: "password",
	// }
	// fmt.Println(u)

	// u.CreateUser()
	ctx := context.Background()
	u, err := models.GetUser(ctx, 1)

	if err != nil {
			panic(err)
	}

	if u.Name == "tekitou" {
		fmt.Println("User exists")
	}

	fmt.Println(u)// 更新前のレコード

	u.Name = "a"
	u.Email = "a@a,com"

	u.UpdateUser(ctx)

	fmt.Println(u)// 更新後のレコード
}
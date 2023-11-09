package main

import (
	"fmt"
	"app/config"
	"app/models"
	// "log"
	"context"
)

func main() {
	ctx := context.Background()
	config.InitDB()
	models.CreateUsersTable(ctx)
	models.CreateTodosTable(ctx)
	// fmt.Println(config.Config.AppPort)
	// fmt.Println(config.Config.DBUser)
	// fmt.Println(config.Config.DBPassword)
	// fmt.Println(config.Config.DBHost)
	// fmt.Println(config.Config.DBPort)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	// fmt.Println(models.Db)

	u := &models.User{
		Name: "washoi",
		Email: "washoi@example.com",
		Password: "password",
	}
	fmt.Println(u)

	u.CreateUser(ctx)
	// u, err := models.GetUser(ctx, 2)

	// if err != nil {
	// 		panic(err)
	// }

	// if u.Name == "tekitou" {
	// 	fmt.Println("User exists")
	// }

	// fmt.Println(u)// 更新前のレコード

	// u.Name = "a"
	// u.Email = "a@a,com"

	// u.UpdateUser(ctx)

	// fmt.Println(u)// 更新後のレコード
	// u.DeleteUser(ctx)
	// u2, err := models.GetUser(ctx, 1)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(u2)// 更新前のレコード
}
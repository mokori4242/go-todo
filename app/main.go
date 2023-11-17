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

	t, err := models.GetTodo(ctx, 2)

	if err != nil {
			panic(err)
	}


	fmt.Println(t)

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
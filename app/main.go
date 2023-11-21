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

	cu := models.User{
		Name: "tekitou",
		Email: "a@example.com",
		Password: "password",
	}

	cu.CreateUser(ctx)

	u, err := models.GetUser(ctx, 1)

	if err != nil {
		panic(err)
	}

	u.CreateTodo(ctx, "test")
	u.CreateTodo(ctx, "test2")
	u.CreateTodo(ctx, "test3")

	t, err := models.GetTodos()

	if err != nil {
			panic(err)
	}

	for _, v := range t {
	fmt.Println(v.Content)
	}

	// fmt.Println(t)

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
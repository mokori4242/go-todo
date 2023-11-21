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

	cu, err := models.GetUser(ctx, 2)

	cu.CreateTodo(ctx, "tekitou1")
	cu.CreateTodo(ctx, "tekitou2")
	cu.CreateTodo(ctx, "tekitou3")

	tby, err := cu.GetTodosByUser(ctx)

	if err != nil {
		panic(err)
	}

	for _, v := range tby {
		fmt.Println(v)
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
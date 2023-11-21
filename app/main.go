package main

import (
	"fmt"
	"app/config"
	"app/models"
	"log"
	"context"
)

func main() {
	ctx := context.Background()
	config.InitDB()
	models.CreateUsersTable(ctx)
	models.CreateTodosTable(ctx)

	// cu, err := models.GetUser(ctx, 2)

	// cu.CreateTodo(ctx, "tekitou1")
	// cu.CreateTodo(ctx, "tekitou2")
	// cu.CreateTodo(ctx, "tekitou3")

	// tby, err := cu.GetTodosByUser(ctx)

	// if err != nil {
	// 	panic(err)
	// }

	// for _, v := range tby {
	// 	fmt.Println(v)
	// }

	t, err := models.GetTodo(ctx, 1)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(t.Content)

	err = t.UpdateTodoContent(ctx, "更新だよ")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(t.Content)
}
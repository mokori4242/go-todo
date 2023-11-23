package main

import (
	// "fmt"
	"app/config"
	"app/models"
	"app/controllers"
	// "log"
	"context"
)

func main() {
	ctx := context.Background()
	config.InitDB()
	models.CreateUsersTable(ctx)
	models.CreateTodosTable(ctx)
	controllers.StartMainServer()
}
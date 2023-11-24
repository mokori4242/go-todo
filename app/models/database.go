package models

import (
	"context"
)

func init() {
	ctx := context.Background()
	CreateUsersTable(ctx)
	CreateTodosTable(ctx)
	CreateSessionsTable(ctx)
}
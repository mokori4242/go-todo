package models

import (
	"time"
	"log"
	"context"
)

type User struct {
	ID int
	UUID string
	Name string
	Email string
	Password string
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `INSERT INTO users (
		uuid,
		name,
		email,
		password,
		created_at) VALUES ($1, $2, $3, $4, $5)`

	_, err = Db.Exec(cmd, createUUID(), u.Name, u.Email, Encrypt(u.Password), time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUser(ctx context.Context, id int) (user User, err error) {
	user = User{}
	cmd := `select * from users where id = $1`
	err = Db.QueryRowContext(ctx, cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) UpdateUser(ctx context.Context) (err error) {

	cmd := `update users set name = $1, email = $2 where id = $3`
	_, err = Db.ExecContext(ctx, cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
package models

import (
	"time"
	"log"
	"context"
	"app/config"
	"fmt"
	"github.com/pborman/uuid"
	"crypto/sha1"
)

const tableNameU = "users"

type User struct {
	ID int
	UUID string
	Name string
	Email string
	Password string
	CreatedAt time.Time
}

type Session struct {
	ID int
	UUID string
	Email string
	UserID int
	CreatedAt time.Time
}

func CreateUsersTable(ctx context.Context) {
	cmd := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id SERIAL PRIMARY KEY,
		uuid UUID NOT NULL UNIQUE,
		name TEXT,
		email TEXT,
		password TEXT,
		created_at TIMESTAMP)`, tableNameU)

	if _, err := config.Db.ExecContext(ctx, cmd); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Successfully created user tables.")
}

func (u *User) CreateUser(ctx context.Context) (err error) {
	cmd := `INSERT INTO users (
		uuid,
		name,
		email,
		password,
		created_at) VALUES ($1, $2, $3, $4, $5)`

	_, err = config.Db.ExecContext(ctx, cmd, createUUID(), u.Name, u.Email, Encrypt(u.Password), time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUser(ctx context.Context, id int) (user User, err error) {
	user = User{}
	cmd := `select * from users where id = $1`
	err = config.Db.QueryRowContext(ctx, cmd, id).Scan(
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
	_, err = config.Db.ExecContext(ctx, cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser(ctx context.Context) (err error) {
	cmd := `delete from users where id = $1`
	_, err = config.Db.ExecContext(ctx, cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUserByEmail(ctx context.Context, email string) (user User, err error) {
	user = User{}
	cmd := `select * from users where email = $1`
	err = config.Db.QueryRowContext(ctx, cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) CreateSession(ctx context.Context) (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions (
		uuid,
		email,
		user_id,
		created_at) values ($1, $2, $3, $4)`

	_, err = config.Db.ExecContext(ctx, cmd1, createUUID(), u.Email, u.ID, time.Now())

	if err != nil {
		log.Fatalln(err)
	}

	cmd2 := `select * from sessions where user_id = $1 and email = $2`
	err = config.Db.QueryRowContext(ctx, cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt,
	)
	return session, err
}

func (s *Session) CheckSession(ctx context.Context) (valid bool, err error) {
	cmd := `select * from sessions where uuid = $1`
	err = config.Db.QueryRowContext(ctx, cmd, s.UUID).Scan(
		&s.ID,
		&s.UUID,
		&s.Email,
		&s.UserID,
		&s.CreatedAt,
	)
	if err != nil {
		valid = false
		return valid, err
	}
	valid = true
	return valid, err
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj = uuid.NewRandom()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
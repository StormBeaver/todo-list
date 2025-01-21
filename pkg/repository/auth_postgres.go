package repository

import (
	"fmt"

	"github.com/StormBeaver/todo-app"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING users_id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Passowrd)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
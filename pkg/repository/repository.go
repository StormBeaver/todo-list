package repository

import (
	"github.com/StormBeaver/todo-app"
	"github.com/jmoiron/sqlx"
)

type Authorisation interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface{}

type TodoItem interface{}

type Repository struct {
	Authorisation
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Authorisation: NewAuthPostgres(db)}
}

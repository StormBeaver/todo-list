package service

import (
	"github.com/StormBeaver/todo-app"
	"github.com/StormBeaver/todo-app/pkg/repository"
)

type Authorisation interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface{}

type TodoItem interface{}

type Service struct {
	Authorisation
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Authorisation: NewAuthService(repos.Authorisation)}
}

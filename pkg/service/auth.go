package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/StormBeaver/todo-app"
	"github.com/StormBeaver/todo-app/pkg/repository"
)

const salt = "ws46xertyurfcyvi"

type AuthService struct {
	repo repository.Authorisation
}

func NewAuthService(repo repository.Authorisation) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Passowrd = generatePasswordHash(user.Passowrd)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

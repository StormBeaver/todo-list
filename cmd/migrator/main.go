package main

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	m, err := migrate.New(
		"file://c:/Users/Meow/go/src/github.com/StormBeaver/todo-app/pkg/migrations",
		"postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No migrations to apply")
			return
		}
		panic(err)
	}

	fmt.Println("All migrations applied successfully")
}

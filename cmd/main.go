package main

import (
	"log"

	"github.com/StormBeaver/todo-app"
	"github.com/StormBeaver/todo-app/pkg/handler"
)

func main() {
	handlers := handler.Handler{}

	srvr := new(todo.Server)
	if err := srvr.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

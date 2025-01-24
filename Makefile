docker_postgres:
	docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres

migrations:
	go run ./cmd/migrator

run: docker_postgres migrations
	echo DB_PASSWORD=qwerty > .env
	go run ./cmd/main.go


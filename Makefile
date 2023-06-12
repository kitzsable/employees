build:
	go build -o employees ./cmd/app/main.go

migration_up: 
	migrate -path ./database -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

migration_down: 
	migrate -path ./database -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down
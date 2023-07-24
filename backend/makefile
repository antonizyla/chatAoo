dev:
	CompileDaemon -build="go build cmd/api/main.go" -command="./main" -color=true
resetDatabase:
	sudo docker-compose down && sudo docker-compose up -d && go run ./cmd/migrate/main.go
test:
	go test ./...

dev:
	CompileDaemon -build="go build cmd/api/main.go" -command="./main" -color=true
emptyDB:
	sudo docker-compose down && sudo docker-compose up -d 
test:
	go test ./...
initaliseDB:
	go run ./cmd/migrate/main.go

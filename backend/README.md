## To install Compile Daemon to watch and recompile upon save use: 

```bash
go install -mod=mod github.com/githubnemo/CompileDaemon
```

## Requires Swag to generate api documentation

```bash 
go install github.com/swaggo/swag/cmd/swag@latest
```

## To Run 

```bash 
make
```
or 
```bash 
CompileDaemon -build="go build cmd/api/main.go" -command="./main" -color=true
```

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var createTables = `
    
    CREATE TABLE IF NOT EXISTS chat (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        description TEXT NOT NULL,
        name TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT NOW()
    );

    CREATE TABLE IF NOT EXISTS users (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        username TEXT NOT NULL
    );

    CREATE TABLE IF NOT EXISTS users_chat (
        chat_id UUID REFERENCES chat(id),
        user_id UUID REFERENCES users(id),
        PRIMARY KEY (chat_id, user_id)
    );

`

func handleError(err error) {
	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		log.Printf("[error] %s:%d %v", filename, line, err)
	}
}

var db *pgxpool.Pool

func init() {

	er := godotenv.Load()
	handleError(er)

	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	dbpool, err := pgxpool.New(context.Background(), connection)
	if err != nil {
		fmt.Println("Error connecting to database")
		log.Fatal(err)
	}
	db = dbpool

	_, err = dbpool.Exec(context.Background(), createTables)
	if err != nil {
		fmt.Println("Error creating tables")
		handleError(err)
	} else {
		fmt.Println("Tables created/updated to reflect schema")
	}

}

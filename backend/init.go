package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"runtime"
)

var createTables = `
    
    CREATE TABLE IF NOT EXISTS chat (
        id UUID PRIMARY KEY,
        name TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT NOW()
    );

`

var db *sql.DB

func handleError(err error) {
	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		log.Printf("[error] %s:%d %v", filename, line, err)
	}
}

func init() {

	er := godotenv.Load()
	handleError(er)

	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/London", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	DB, err := sql.Open("postgres", connection)
	handleError(err)

	err = DB.Ping()
	if err != nil {
		handleError(err)
	} else {
		fmt.Println("Connected to database Successfully")
	}

	_, err = DB.Exec(createTables)
	handleError(err)

	db = DB

}

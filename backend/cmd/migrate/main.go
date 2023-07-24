package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	db, err := initialiseDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// order of database tables to be created
	tables := []string{"createChatTable", "createUsersTable", "createUsersChatTable", "createMessagesTable"}

	for _, table := range tables {
		// read the file and execute the query
		fileContents, err := os.ReadFile(fmt.Sprintf("migrations/%s.sql", table))
		if err != nil {
			log.Fatalf("Unable to read file: %v", err)
		}

		_, err = db.Exec(context.Background(), string(fileContents))
		if err != nil {
			log.Fatalf("Unable to execute query: %v", err)
		}

		fmt.Printf("Successfully executed query: %s\n", table)
	}

}

func initialiseDatabase() (pool *pgxpool.Pool, err error) {

	er := godotenv.Load()
	if er != nil {
		return nil, er
	}

	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	dbpool, err := pgxpool.New(context.Background(), connection)
	if err != nil {
		fmt.Println("Error connecting to database")
		log.Fatal(err)
	}

	return dbpool, nil
}

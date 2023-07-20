package main

import (
	"backend/api/router"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// create global database pool
var DB *pgxpool.Pool

func main() {

	// handle database
	database, err := initialiseDatabase()
	if err != nil {
		log.Fatal(err)
	}
	DB = database
	defer DB.Close()
	
    r := router.New(DB)

	s := http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	// handle http server
	log.Printf("Starting http server on port %s", s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Failed to Start server")
	} else {
		log.Println("Server Started Successfully")
	}
	defer s.Close()
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

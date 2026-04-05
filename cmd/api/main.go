package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type application struct {
	DB *sql.DB
}

func main() {
	connStr := "user=studyuser password=secret dbname=studyapp sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		DB: db,
	}

	mux := http.NewServeMux()
	app.routes(mux)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

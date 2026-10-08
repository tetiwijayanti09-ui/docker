package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := db.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		response := map[string]string{
			"message":  "pong",
			"database": "disconnected",
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := map[string]string{
		"message":  "pong",
		"database": "connected",
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPassword +
		"@tcp(" + dbHost + ":" + dbPort + ")/" + dbName

	var err error

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	http.HandleFunc("/", pingHandler)

	log.Println("Hello World application listening on :5000")

	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}
}
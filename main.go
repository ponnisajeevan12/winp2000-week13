package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Response struct {
	CurrentTime string `json:"current_time"`
}

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get database credentials from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// DSN (Data Source Name) for MySQL connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	http.HandleFunc("/current-time", func(w http.ResponseWriter, r *http.Request) {
		location, err := time.LoadLocation("America/Toronto")
		if err != nil {
			http.Error(w, "Loading time zone - failed", http.StatusInternalServerError)
			return
		}

		// Get current time in Toronto timezone
		currentTime := time.Now().In(location)

		// Define the time format (e.g., "2006-01-02 15:04:05")
		timeFormat := "2006-01-02 15:04:05"
		formattedTime := currentTime.Format(timeFormat)

		// Insert the formatted time into the MySQL database
		_, err = db.Exec("INSERT INTO time_log (timestamp) VALUES (?)", formattedTime)
		if err != nil {
			http.Error(w, "Logging time to database - Failed", http.StatusInternalServerError)
			return
		}

		// Create the response object with formatted time
		response := Response{CurrentTime: formattedTime}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

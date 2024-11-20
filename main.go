package main

import (
	"net/http"
	"retail-pulse/api"
)

func main() {
	// Define routes
	http.HandleFunc("/api/submit/", api.SubmitJob)
	http.HandleFunc("/api/status", api.GetJobStatus)

	// Start the server
	port := ":8080"
	println("Server running on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		println("Error starting server:", err.Error())
	}
}

// package main

// import (
// 	"database/sql"

// 	"fmt"
// 	"log"
// 	"net/http"

// 	"retail-pulse/api"

// 	_ "github.com/go-sql-driver/mysql" // MySQL driver
// )

// var db *sql.DB

// func init() {
// 	// Set up the MySQL connection string (replace with your actual credentials)
// 	dsn := "root:Beena1123$@tcp(127.0.0.1:3306)/Store"
// 	var err error
// 	db, err = sql.Open("mysql", dsn)
// 	if err != nil {
// 		log.Fatal("Error opening database:", err)
// 	}

// 	// Ping the database to check if the connection is successful
// 	if err := db.Ping(); err != nil {
// 		log.Fatal("Error pinging database:", err)
// 	}

// 	fmt.Println("Successfully connected to the database!")
// }

// func main() {
// 	// Define routes
// 	http.HandleFunc("/api/submit/", api.SubmitJob)
// 	http.HandleFunc("/api/status", api.GetJobStatus)

// 	// New route to upload CSV data
// 	http.HandleFunc("/api/uploadcsv", uploadCSV)

// 	// Start the server
// 	port := ":8080"
// 	println("Server running on port", port)
// 	if err := http.ListenAndServe(port, nil); err != nil {
// 		println("Error starting server:", err.Error())
// 	}
// }

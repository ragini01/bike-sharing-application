package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"BIKE-SHARING-SERVICE/internal/config"
	"BIKE-SHARING-SERVICE/internal/handlers"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// Connect to database
	db, err := db.Connect(cfg.DBURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Close database connection when the application exits
	defer db.Close()

	// Initialize router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/login", handlers.Login(db, GetUserSession)).Methods("POST")
	router.HandleFunc("/bicycles", handlers.ListBicycles(db)).Methods("GET")
	router.HandleFunc("/bicycles/rent", authMiddleware(handlers.RentBicycle(db))).Methods("POST")
	router.HandleFunc("/bicycles/return", authMiddleware(handlers.ReturnBicycle(db))).Methods("POST")

	// Start HTTP server
	log.Fatal(http.ListenAndServe(":8080", router))

}

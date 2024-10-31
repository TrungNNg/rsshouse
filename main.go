package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/TrungNNg/rsshouse/api"
	"github.com/TrungNNg/rsshouse/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	cfg := api.ApiConfig{
		DB:     database.New(db),
		Secret: os.Getenv("SECRET"),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", cfg.HomeHandler)
	mux.HandleFunc("POST /api/signup", cfg.SignUp)
	mux.HandleFunc("POST /api/login", cfg.Login)
	mux.HandleFunc("POST /api/logout", cfg.Logout)
	mux.HandleFunc("GET /api/reset", cfg.ResetUsers)
	mux.HandleFunc("POST /api/refresh", cfg.Refresh)
	mux.HandleFunc("POST /api/revoke", cfg.Revoke)

	// authenticated enpoints
	mux.HandleFunc("GET /api/test", cfg.Test)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Println("Starting server on :" + port)
	log.Fatal(server.ListenAndServe())
}

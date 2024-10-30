package main

import (
	"log"
	"net/http"
	"os"

	"github.com/TrungNNg/rsshouse/api"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	cfg := api.ApiConfig{}

	mux := http.NewServeMux()
	mux.HandleFunc("/", cfg.HomeHandler)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Println("Starting server on :" + port)
	log.Fatal(server.ListenAndServe())
}

package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/TrungNNg/rsshouse/internal/api"
	"github.com/TrungNNg/rsshouse/internal/database"
	"github.com/TrungNNg/rsshouse/internal/jobs"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mmcdole/gofeed"
)

func testGoFeed() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	fp := gofeed.NewParser()
	fp.UserAgent = "rsshouse"
	feed, _ := fp.ParseURLWithContext("http://feeds.twit.tv/twit.xml", ctx)

	post := feed.Items[1]
	println(post.Title)
	println(post.Description)
	println("Content", post.Content)
	println("LINK:", post.Link)
	for _, l := range post.Links {
		println(l)
	}
	println("Update time: ", post.UpdatedParsed)
	println("Publisised time: ", post.PublishedParsed)
	for _, aut := range post.Authors {
		println(aut.Name)
	}
	println("ID HERE: ", post.GUID)
	println("IMAGE TITLE: ", post.Image.Title)
	println("IMAGE URL: ", post.Image.URL)
}

func main() {

	//testGoFeed()

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

	fp := gofeed.NewParser()
	fp.UserAgent = "rsshouse"

	cfg := api.ApiConfig{
		DB:                          database.New(db),
		Secret:                      os.Getenv("SECRET"),
		Parser:                      fp,
		FetchFeedInterval:           time.Minute * 30,
		CleanupRefreshTokenInterval: time.Hour * 12,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", cfg.HomeHandler)
	mux.HandleFunc("POST /api/signup", cfg.SignUp)
	mux.HandleFunc("POST /api/login", cfg.Login)
	mux.HandleFunc("GET /api/reset", cfg.ResetUsers)
	mux.HandleFunc("POST /api/refresh", cfg.Refresh)
	mux.HandleFunc("POST /api/revoke", cfg.Revoke) // use this for logout

	// authenticated endpoints
	mux.HandleFunc("GET /api/test", cfg.Test)
	mux.HandleFunc("POST /api/feeds", cfg.AddFeed) // need test
	mux.HandleFunc("POST /api/subscribe", cfg.SubcribeFeed)
	mux.HandleFunc("POST /api/unsubscribe", cfg.UnsubcribeFeed)
	mux.HandleFunc("GET /api/posts", cfg.GetPosts)
	mux.HandleFunc("POST /api/save", cfg.SavePost)     // save post for user, need test
	mux.HandleFunc("POST /api/unsave", cfg.UnsavePost) // unsave post for user, need test

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// a background task that clean up expired or revoked refresh token
	go jobs.CleanUpRefreshToken(&cfg)

	// start the server
	log.Println("Starting server on :" + port)
	log.Fatal(server.ListenAndServe())
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bryantang1107/Rss_Miner/http_handler"
	"github.com/bryantang1107/Rss_Miner/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq" // need to import db driver for sqlc
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("UNABLE TO LOAD ENVIRONMENT VARIABLES")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT NOT FOUND")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB CONNECTION STRING NOT FOUND")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to DB")
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	db := database.New(conn)
	apiCfg := http_handler.ApiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute)

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", http_handler.HandlerReadiness)
	v1Router.Get("/error", http_handler.HandleErr)
	v1Router.Post("/user", apiCfg.HandlerCreateUser)
	v1Router.Get("/user", apiCfg.MiddlewareAuth(apiCfg.HandlerGetUser))
	v1Router.Get("/feed", apiCfg.HandlerGetFeed)
	v1Router.Post("/feed", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeed))
	v1Router.Get("/posts", apiCfg.MiddlewareAuth(apiCfg.HandlerGetPostsForUser))
	v1Router.Get("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerGetFeedFollow))
	v1Router.Post("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeedFollow))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.MiddlewareAuth(apiCfg.HandlerDeleteFeedFollow))

	router.Mount("/v1", v1Router) // attach subrouter to main router

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Printf("Server Starting On Port %s", port)
	fmt.Println()
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Test")

	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT NOT FOUND")
	}

	router := chi.NewRouter()

	server := http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	fmt.Println(port)
}

package main

import (
	"log"
	"net/http"
	"pencarian_elektronik/handlers"
	"pencarian_elektronik/repositories"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found. Using system environment variables.")
	}

	repositories.InitDB()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.SearchPageHandler)
	http.HandleFunc("/search", handlers.SearchHandler)

	certFile := "certs/server.crt"
	keyFile := "certs/server.key"
	port := ":8443"

	log.Printf("Server running at https://localhost%\n", port)
	log.Fatal(http.ListenAndServeTLS(port, certFile, keyFile, nil))
}

package main

import (
	"log"
	"net/http"

	"github.com/xinglixing/chat-app/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("Start channel listener...")
	go handlers.ListenToWSChannel()

	log.Println("Starting web server on port 8080")
	_ = http.ListenAndServe(":8080", mux)
}

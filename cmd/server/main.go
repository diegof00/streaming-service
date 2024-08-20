package main

import (
	"log"
	"net/http"

	"github.com/diegof00/streaming-service/internal/handlers"
)

func main() {
	http.HandleFunc("/ws", handlers.HandleWebSocket)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

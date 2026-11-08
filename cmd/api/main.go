package main

import (
	"log"
	"net/http"

	"github.com/MuHucTep/REST-API-task-flow/internal/handler"
)

func main() {
	http.HandleFunc("/health", handler.Health)
	http.HandleFunc("/", handler.Root)

	addr := ":8000"

	log.Printf("Server is starting on %s\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

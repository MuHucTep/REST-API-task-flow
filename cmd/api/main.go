package main

import (
	"log"
	"net/http"

	"github.com/MuHucTep/REST-API-task-flow/internal/config"
	"github.com/MuHucTep/REST-API-task-flow/internal/handler"
	"github.com/MuHucTep/REST-API-task-flow/internal/middleware"
	"github.com/MuHucTep/REST-API-task-flow/internal/repository"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.Load()

	db, err := repository.NewDB(
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPswd,
		cfg.DBName,
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	r := chi.NewRouter()

	r.Use(middleware.Recovery)
	r.Use(middleware.Logger)

	r.Get("/health", handler.Health)
	r.Get("/", handler.Root)

	log.Printf("Server is starting on %s\n", cfg.HTTPAddr)

	if err := http.ListenAndServe(cfg.HTTPAddr, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

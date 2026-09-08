package main

import (
	"context"
	"log"
	"time"

	"ebook-store/backend/internal/config"
	"ebook-store/backend/internal/ebook"
	"ebook-store/backend/internal/httpapi"
	"ebook-store/backend/internal/platform/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load configuration: %v", err)
	}

	connectionContext, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	databasePool, err := database.NewPostgresPool(
		connectionContext,
		cfg.DatabaseURL,
	)
	if err != nil {
		log.Fatalf("connect to database: %v", err)
	}
	defer databasePool.Close()

	log.Println("database connection established")

	ebookRepository := ebook.NewPostgresRepository(databasePool)
	ebookService := ebook.NewService(ebookRepository)
	ebookHandler := ebook.NewHandler(ebookService)

	router := httpapi.NewRouter(ebookHandler)

	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

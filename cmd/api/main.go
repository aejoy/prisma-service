package main

import (
	"fmt"
	"github.com/aejoy/prisma-service/internal/config"
	"github.com/aejoy/prisma-service/internal/handlers/http"
	"github.com/aejoy/prisma-service/internal/repositories/db"
	"github.com/aejoy/prisma-service/internal/repositories/storage"
	"github.com/aejoy/prisma-service/internal/services/photos"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	db, err := db.NewPostgres(cfg.Shards)
	if err != nil {
		panic(err)
	}

	storage, err := storage.NewStorage(cfg.Storage.Domain, cfg.Storage.Endpoint,
		cfg.Storage.AccessKeyID, cfg.Storage.SecretAccessKey,
		cfg.Storage.Bucket, cfg.Storage.Region)
	if err != nil {
		panic(err)
	}

	service := photos.NewPhotoService(db, storage)
	handlers := http.NewHTTPHandlers(service)

	router := fiber.New(fiber.Config{BodyLimit: 30 * 1024 * 1024}) // max=30mB

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{cfg.Service.AllowOrigin},
		AllowMethods: []string{"GET", "POST", "PATCH"},
	}))

	router.Use(logger.New(logger.Config{Format: "${pid} ${status} - ${method} ${path}\n"}))

	group := router.Group(cfg.Service.APIPrefix)
	group.Post("/photos/upload", handlers.Upload)
	group.Get("/photos", handlers.Photos)

	log.Fatalln(router.Listen(fmt.Sprintf(":%d", cfg.Service.PORT)))
}

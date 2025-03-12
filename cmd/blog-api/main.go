package main

import (
	"blog-api/internal/configs"
	"blog-api/internal/database"
	"blog-api/internal/handlers"
	"blog-api/internal/logger"
	"blog-api/internal/repository"
	"blog-api/internal/routers"
	"blog-api/internal/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

func init() {
	logger.LoadLogger()
}

func main() {

	logger.Log.Info().Msg("blog-api application starting")

	config := configs.LoadConfig()

	gorm, err := database.ConnectDB(&config.DBUrl)
	if err != nil {
		return
	}

	repo, err := repository.NewBlogRepo(gorm)
	if err != nil {
		return
	}

	service, err := services.NewBlogService(&repo)
	if err != nil {
		return
	}

	handler, err := handlers.NewBlogHandler(&service)
	if err != nil {
		return
	}

	app := fiber.New()
	routers.SetupRoutes(app, handler)

	logger.Log.Info().Msg("sever is listening")
	port := config.Port
	if port == "" {
		port = ":8080"
	}
	log.Fatal(app.Listen(config.Port))
}

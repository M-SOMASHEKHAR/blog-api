package main

import (
	config "blog-api/internal/configs"
	"blog-api/internal/database"
	"blog-api/internal/handlers"
	"blog-api/internal/repository"
	routes "blog-api/internal/routers"
	"blog-api/internal/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config := config.LoadConfig()

	gorm, err := database.ConnectDB(config.DBUrl)
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
	routes.SetupRoutes(app, handler)

	log.Fatal(app.Listen(config.Port))
}

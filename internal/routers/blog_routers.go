package routes

import (
	"blog-api/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, handler handlers.BlogHandler) {
	api := app.Group("/api/blog-post")
	api.Post("/", handler.CreateBlog)
	api.Get("/", handler.GetAllBlogs)
	api.Get("/:id", handler.GetBlogByID)
	api.Patch("/:id", handler.UpdateBlog)
	api.Delete("/:id", handler.DeleteBlog)
}

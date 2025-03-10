package handlers

import (
	"blog-api/internal/services"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service services.BlogService
}

type BlogHandler interface {
	CreateBlog(c *fiber.Ctx)
	GetAllBlogs(c *fiber.Ctx)
	GetBlogByID(c *fiber.Ctx)
	UpdateBlog(c *fiber.Ctx)
	DeleteBlog(c *fiber.Ctx)
}

func NewBlogHandler(service *services.BlogService) (BlogHandler, error) {
	if service == nil {
		return nil, errors.New("")
	}
	return &handler{
		service: *service,
	}, nil
}

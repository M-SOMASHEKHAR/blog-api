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
	CreateBlog(c *fiber.Ctx) error
	GetAllBlogs(c *fiber.Ctx) error
	GetBlogByID(c *fiber.Ctx) error
	UpdateBlog(c *fiber.Ctx) error
	DeleteBlog(c *fiber.Ctx) error
}

func NewBlogHandler(service *services.BlogService) (BlogHandler, error) {
	if service == nil {
		return nil, errors.New("")
	}
	return &handler{
		service: *service,
	}, nil
}

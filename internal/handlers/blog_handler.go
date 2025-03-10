package handlers

import (
	"blog-api/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) CreateBlog(c *fiber.Ctx) error {
	blog := new(models.Blog)
	if err := c.BodyParser(blog); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.service.CreateBlogService(blog)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create blog"})
	}

	return c.Status(201).JSON(blog)
}

func (h *handler) GetAllBlogs(c *fiber.Ctx) error {
	blogs, err := h.service.GetAllBlogsService()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch blogs"})
	}
	return c.JSON(blogs)
}

func (h *handler) GetBlogByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog, err := h.service.GetBlogByIDService(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Blog not found"})
	}
	return c.JSON(blog)
}

func (h *handler) UpdateBlog(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog, err := h.service.GetBlogByIDService(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Blog not found"})
	}

	if err := c.BodyParser(&blog); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	h.service.UpdateBlogService(&blog)
	return c.JSON(blog)
}

func (h *handler) DeleteBlog(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	err := h.service.DeleteBlogService(uint(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete blog"})
	}
	return c.SendStatus(204)
}

package handlers

import (
	"blog-api/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) CreateBlog(c *fiber.Ctx) {
	blog := new(models.Blog)
	if err := c.BodyParser(blog); err != nil {
		c.Status(400).JSON(fiber.Map{"error": err.Error()})
		return
	}

	err := h.service.CreateBlogService(blog)
	if err != nil {
		c.Status(500).JSON(fiber.Map{"error": "Failed to create blog"})
		return
	}

	c.Status(201).JSON(blog)
	return
}

func (h *handler) GetAllBlogs(c *fiber.Ctx) {
	blogs, err := h.service.GetAllBlogsService()
	if err != nil {
		c.Status(500).JSON(fiber.Map{"error": "Failed to fetch blogs"})
		return
	}
	c.JSON(blogs)
	return
}

func (h *handler) GetBlogByID(c *fiber.Ctx) {
	id, _ := strconv.Atoi(c.Params("id"))
	blog, err := h.service.GetBlogByIDService(uint(id))
	if err != nil {
		c.Status(404).JSON(fiber.Map{"error": "Blog not found"})
		return
	}
	c.JSON(blog)
	return
}

func (h *handler) UpdateBlog(c *fiber.Ctx) {
	id, _ := strconv.Atoi(c.Params("id"))
	blog, err := h.service.GetBlogByIDService(uint(id))
	if err != nil {
		c.Status(404).JSON(fiber.Map{"error": "Blog not found"})
		return
	}

	if err := c.BodyParser(&blog); err != nil {
		c.Status(400).JSON(fiber.Map{"error": err.Error()})
		return
	}

	h.service.UpdateBlogService(&blog)
	c.JSON(blog)
	return
}

func (h *handler) DeleteBlog(c *fiber.Ctx) {
	id, _ := strconv.Atoi(c.Params("id"))
	err := h.service.DeleteBlogService(uint(id))
	if err != nil {
		c.Status(500).JSON(fiber.Map{"error": "Failed to delete blog"})
		return
	}
	c.SendStatus(204)
	return
}

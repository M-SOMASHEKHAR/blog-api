package handlers

import (
	"blog-api/internal/logger"
	"blog-api/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) CreateBlog(c *fiber.Ctx) error {
	logger.Log.Info().Msg("CreateBlog API started")

	blog := new(models.Blog)
	if err := c.BodyParser(blog); err != nil {
		logger.Log.Error().Err(err).Msg("Failed to parse request body")
		logger.Log.Info().Msg("CreateBlog API ended")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":   "Invalid request payload",
			"error": err.Error(),
		})
	}

	err := h.service.CreateBlogService(blog)
	if err != nil {
		logger.Log.Info().Msg("CreateBlog API ended")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg":   "Failed to create blog",
			"error": err.Error(),
		})
	}

	logger.Log.Info().Msg("CreateBlog API ended")
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"msg":  "blog creation successfull",
		"blog": blog,
	})
}

func (h *handler) GetAllBlogs(c *fiber.Ctx) error {
	logger.Log.Info().Msg("GetAllBlogs API started")

	blogs, err := h.service.GetAllBlogsService()
	if err != nil {
		logger.Log.Info().Msg("GetAllBlogs API ended")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg":   "Failed to fetch blogs",
			"error": err.Error(),
		})
	}

	logger.Log.Info().Msg("GetAllBlogs API ended")
	return c.JSON(fiber.Map{
		"msg":   "blogs fetched successfully",
		"blogs": blogs,
	})
}

func (h *handler) GetBlogByID(c *fiber.Ctx) error {
	logger.Log.Info().Msg("GetBlogByID API started")

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid blog ID")
		logger.Log.Info().Msg("GetBlogByID API ended")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid blog ID"})
	}

	blogID := uint(id)
	blog, err := h.service.GetBlogByIDService(&blogID)
	if err != nil {
		logger.Log.Info().Msg("GetBlogByID API ended")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg":   "Blog not found",
			"error": err.Error(),
		})
	}

	logger.Log.Info().Msg("GetBlogByID API ended")
	return c.JSON(fiber.Map{
		"msg":  "blog fetched successfully",
		"blog": blog,
	})
}

func (h *handler) UpdateBlog(c *fiber.Ctx) error {
	logger.Log.Info().Msg("UpdateBlog API started")

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid blog ID")
		logger.Log.Info().Msg("UpdateBlog API ended")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid blog ID"})
	}

	blogID := uint(id)
	if err := h.service.UpdateBlogService(&blogID); err != nil {
		logger.Log.Info().Msg("UpdateBlog API ended")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg":   "Failed to update blog",
			"error": err.Error(),
		})
	}

	logger.Log.Info().Msg("UpdateBlog API ended")
	return c.JSON(fiber.Map{
		"msg":     "blog updated successfully",
		"blog id": blogID,
	})
}

func (h *handler) DeleteBlog(c *fiber.Ctx) error {
	logger.Log.Info().Msg("DeleteBlog API started")

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid blog ID")
		logger.Log.Info().Msg("DeleteBlog API ended")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid blog ID"})
	}

	blogID := uint(id)
	err = h.service.DeleteBlogService(&blogID)
	if err != nil {
		logger.Log.Info().Msg("DeleteBlog API ended")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg":   "Failed to delete blog",
			"error": err.Error(),
		})
	}

	logger.Log.Info().Msg("DeleteBlog API ended")
	return c.JSON(fiber.Map{
		"msg":     "blog deleted successfully",
		"blog id": blogID,
	})
}

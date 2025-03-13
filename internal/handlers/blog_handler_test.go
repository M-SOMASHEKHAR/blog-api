package handlers

import (
	"blog-api/internal/mocks"
	"blog-api/internal/models"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_handler_CreateBlog(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockBlogService(ctrl)
	handler := &handler{service: mockService}

	app := fiber.New()
	app.Post("/blogs", handler.CreateBlog)

	t.Run("success", func(t *testing.T) {
		mockService.EXPECT().CreateBlogService(gomock.Any()).Return(nil).Times(1)

		req := httptest.NewRequest(http.MethodPost, "/blogs", strings.NewReader(`{"title":"Test Blog","description":"Test Description","body":"Test Body"}`))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	})

	t.Run("failure", func(t *testing.T) {
		mockService.EXPECT().CreateBlogService(gomock.Any()).Return(errors.New("failed to create blog")).Times(1)

		req := httptest.NewRequest(http.MethodPost, "/blogs", strings.NewReader(`{"title":"Test Blog","description":"Test Description","body":"Test Body"}`))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})
}

func Test_handler_GetAllBlogs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockBlogService(ctrl)
	handler := &handler{service: mockService}

	app := fiber.New()
	app.Get("/blogs", handler.GetAllBlogs)

	t.Run("success", func(t *testing.T) {
		mockService.EXPECT().GetAllBlogsService().Return(&[]models.Blog{}, nil).Times(1)

		req := httptest.NewRequest(http.MethodGet, "/blogs", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("failure", func(t *testing.T) {
		mockService.EXPECT().GetAllBlogsService().Return(nil, errors.New("failed to fetch blogs")).Times(1)

		req := httptest.NewRequest(http.MethodGet, "/blogs", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})
}

func Test_handler_GetBlogByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockBlogService(ctrl)
	handler := &handler{service: mockService}

	app := fiber.New()
	app.Get("/blogs/:id", handler.GetBlogByID)

	t.Run("success", func(t *testing.T) {
		mockService.EXPECT().GetBlogByIDService(gomock.Any()).Return(&models.Blog{}, nil).Times(1)

		req := httptest.NewRequest(http.MethodGet, "/blogs/1", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("failure", func(t *testing.T) {
		mockService.EXPECT().GetBlogByIDService(gomock.Any()).Return(nil, errors.New("blog not found")).Times(1)

		req := httptest.NewRequest(http.MethodGet, "/blogs/1", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})
}

func Test_handler_UpdateBlog(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockBlogService(ctrl)
	handler := &handler{service: mockService}

	app := fiber.New()
	app.Patch("/blogs/:id", handler.UpdateBlog)

	t.Run("success", func(t *testing.T) {
		mockService.EXPECT().UpdateBlogService(gomock.Any()).Return(nil).Times(1)

		req := httptest.NewRequest(http.MethodPatch, "/blogs/1", strings.NewReader(`{"title":"Updated Blog"}`))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("failure", func(t *testing.T) {
		mockService.EXPECT().UpdateBlogService(gomock.Any()).Return(errors.New("failed to update blog")).Times(1)

		req := httptest.NewRequest(http.MethodPatch, "/blogs/1", strings.NewReader(`{"title":"Updated Blog"}`))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})
}

func Test_handler_DeleteBlog(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockBlogService(ctrl)
	handler := &handler{service: mockService}

	app := fiber.New()
	app.Delete("/blogs/:id", handler.DeleteBlog)

	t.Run("success", func(t *testing.T) {
		mockService.EXPECT().DeleteBlogService(gomock.Any()).Return(nil).Times(1)

		req := httptest.NewRequest(http.MethodDelete, "/blogs/1", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("failure", func(t *testing.T) {
		mockService.EXPECT().DeleteBlogService(gomock.Any()).Return(errors.New("failed to delete blog")).Times(1)

		req := httptest.NewRequest(http.MethodDelete, "/blogs/1", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})
}

package services

import (
	"blog-api/internal/mocks"
	"blog-api/internal/models"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_service_CreateBlogService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBlogRepo(ctrl)
	service := &service{repo: mockRepo}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().CreateBlog(gomock.Any()).Return(nil).Times(1)

		blog := &models.Blog{
			Title:       "Test Blog",
			Description: "Test Description",
			Body:        "Test Body",
		}
		err := service.CreateBlogService(blog)

		assert.NoError(t, err)
		assert.NotZero(t, blog.CreatedAt)
		assert.NotZero(t, blog.UpdatedAt)
	})

	t.Run("failure", func(t *testing.T) {
		mockRepo.EXPECT().CreateBlog(gomock.Any()).Return(errors.New("failed to create blog")).Times(1)

		blog := &models.Blog{
			Title:       "Test Blog",
			Description: "Test Description",
			Body:        "Test Body",
		}
		err := service.CreateBlogService(blog)

		assert.Error(t, err)
	})
}

func Test_service_GetAllBlogsService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBlogRepo(ctrl)
	service := &service{repo: mockRepo}

	t.Run("failure", func(t *testing.T) {
		mockRepo.EXPECT().GetAllBlogs().Return(nil, errors.New("failed to fetch blogs")).Times(1)

		blogs, err := service.GetAllBlogsService()

		assert.Error(t, err)
		assert.Nil(t, blogs)
	})

	t.Run("no blogs found", func(t *testing.T) {
		mockRepo.EXPECT().GetAllBlogs().Return(&[]models.Blog{}, nil).Times(1)

		blogs, err := service.GetAllBlogsService()

		assert.Error(t, err)
		assert.Nil(t, blogs)
	})
}

func Test_service_GetBlogByIDService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBlogRepo(ctrl)
	service := &service{repo: mockRepo}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().GetBlogByID(gomock.Any()).Return(&models.Blog{}, nil).Times(1)

		id := uint(1)
		blog, err := service.GetBlogByIDService(&id)

		assert.NoError(t, err)
		assert.NotNil(t, blog)
	})

	t.Run("failure", func(t *testing.T) {
		mockRepo.EXPECT().GetBlogByID(gomock.Any()).Return(nil, errors.New("blog not found")).Times(1)

		id := uint(1)
		blog, err := service.GetBlogByIDService(&id)

		assert.Error(t, err)
		assert.Nil(t, blog)
	})
}

func Test_service_UpdateBlogService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBlogRepo(ctrl)
	service := &service{repo: mockRepo}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().GetBlogByID(gomock.Any()).Return(&models.Blog{}, nil).Times(1)
		mockRepo.EXPECT().UpdateBlog(gomock.Any()).Return(nil).Times(1)

		id := uint(1)
		err := service.UpdateBlogService(&id)

		assert.NoError(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		mockRepo.EXPECT().GetBlogByID(gomock.Any()).Return(nil, errors.New("blog not found")).Times(1)

		id := uint(1)
		err := service.UpdateBlogService(&id)

		assert.Error(t, err)
	})
}

func Test_service_DeleteBlogService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBlogRepo(ctrl)
	service := &service{repo: mockRepo}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().GetBlogByID(gomock.Any()).Return(&models.Blog{}, nil).Times(1)
		mockRepo.EXPECT().DeleteBlog(gomock.Any()).Return(nil).Times(1)

		id := uint(1)
		err := service.DeleteBlogService(&id)

		assert.NoError(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		mockRepo.EXPECT().GetBlogByID(gomock.Any()).Return(nil, errors.New("blog not found")).Times(1)

		id := uint(1)
		err := service.DeleteBlogService(&id)

		assert.Error(t, err)
	})
}

package services

import (
	"blog-api/internal/logger"
	"blog-api/internal/models"
	"errors"
	"fmt"
	"time"
)

func (s *service) CreateBlogService(blog *models.Blog) error {
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()
	return s.repo.CreateBlog(blog)
}

func (s *service) GetAllBlogsService() (*[]models.Blog, error) {
	blogs, err := s.repo.GetAllBlogs()
	if err != nil {
		return nil, err
	}

	if len(*blogs) < 1 {
		logger.Log.Error().Msg(fmt.Sprintf("Failed to get blogs: %v", errors.New("no blogs found")))
		return nil, errors.New("no blogs found")
	}

	logger.Log.Info().Int("count", len(*blogs)).Msg("Fetched all blogs successfully")

	return blogs, nil
}

func (s *service) GetBlogByIDService(id *uint) (*models.Blog, error) {
	return s.repo.GetBlogByID(id)
}

func (s *service) UpdateBlogService(id *uint) error {
	blog, err := s.repo.GetBlogByID(id)
	if err != nil {
		return err
	}
	blog.UpdatedAt = time.Now()
	return s.repo.UpdateBlog(blog)
}

func (s *service) DeleteBlogService(id *uint) error {
	_, err := s.repo.GetBlogByID(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteBlog(id)
}

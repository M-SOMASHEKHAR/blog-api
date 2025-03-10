package services

import (
	"blog-api/internal/models"
	"errors"
	"time"
)

func (s *service) CreateBlogService(blog *models.Blog) error {
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()
	return s.repo.CreateBlog(blog)
}

func (s *service) GetAllBlogsService() ([]models.Blog, error) {
	return s.repo.GetAllBlogs()
}

func (s *service) GetBlogByIDService(id uint) (models.Blog, error) {
	blog, err := s.repo.GetBlogByID(id)
	if err != nil {
		return models.Blog{}, errors.New("blog not found")
	}
	return blog, nil
}

func (s *service) UpdateBlogService(blog *models.Blog) error {
	blog.UpdatedAt = time.Now()
	return s.repo.UpdateBlog(blog)
}

func (s *service) DeleteBlogService(id uint) error {
	return s.repo.DeleteBlog(id)
}

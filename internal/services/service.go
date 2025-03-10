package services

import (
	"blog-api/internal/models"
	"blog-api/internal/repository"
	"errors"
)

type service struct {
	repo repository.BlogRepo
}

type BlogService interface {
	CreateBlogService(blog *models.Blog) error
	GetAllBlogsService() ([]models.Blog, error)
	GetBlogByIDService(id uint) (models.Blog, error)
	UpdateBlogService(blog *models.Blog) error
	DeleteBlogService(id uint) error
}

func NewBlogService(repo *repository.BlogRepo) (BlogService, error) {
	if repo == nil {
		return nil, errors.New("")
	}
	return &service{
		repo: *repo,
	}, nil
}

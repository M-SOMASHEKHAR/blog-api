package repository

import (
	"blog-api/internal/models"
	"errors"

	"gorm.io/gorm"
)

type database struct {
	db *gorm.DB
}

type BlogRepo interface {
	CreateBlog(blog *models.Blog) error
	GetAllBlogs() ([]models.Blog, error)
	GetBlogByID(id uint) (models.Blog, error)
	UpdateBlog(blog *models.Blog) error
	DeleteBlog(id uint) error
}

func NewBlogRepo(db *gorm.DB) (BlogRepo, error) {
	if db == nil {
		return nil, errors.New("database connection is nil")
	}
	return &database{db: db}, nil
}

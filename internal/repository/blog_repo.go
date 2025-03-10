package repository

import (
	"blog-api/internal/models"
)

func (r *database) CreateBlog(blog *models.Blog) error {
	return r.db.Create(blog).Error
}

func (r *database) GetAllBlogs() ([]models.Blog, error) {
	var blogs []models.Blog
	err := r.db.Find(&blogs).Error
	return blogs, err
}

func (r *database) GetBlogByID(id uint) (models.Blog, error) {
	var blog models.Blog
	err := r.db.First(&blog, id).Error
	return blog, err
}

func (r *database) UpdateBlog(blog *models.Blog) error {
	return r.db.Save(blog).Error
}

func (r *database) DeleteBlog(id uint) error {
	return r.db.Delete(&models.Blog{}, id).Error
}

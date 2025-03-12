package repository

import (
	"blog-api/internal/logger"
	"blog-api/internal/models"
	"fmt"
)

func (r *database) CreateBlog(blog *models.Blog) error {
	err := r.db.Create(blog).Error
	if err != nil {
		logger.Log.Error().Str("title", blog.Title).Msg(fmt.Sprintf("Failed to create blog: %v", err))
		return err
	}

	logger.Log.Info().Uint("blog id", blog.ID).Str("title", blog.Title).Msg("Blog created successfully")

	return nil
}

func (r *database) GetAllBlogs() (*[]models.Blog, error) {
	var blogs []models.Blog
	err := r.db.Find(&blogs).Error
	if err != nil {
		logger.Log.Error().Msg(fmt.Sprintf("Failed to fetch all blogs: %v", err))
		return nil, err
	}

	return &blogs, nil
}

func (r *database) GetBlogByID(id *uint) (*models.Blog, error) {
	var blog models.Blog
	err := r.db.First(&blog, *id).Error
	if err != nil {
		logger.Log.Warn().Uint("blog id", *id).Msg(fmt.Sprintf("Blog not found: %v", err))
		return nil, err
	}

	logger.Log.Info().Uint("blog id", blog.ID).Str("title", blog.Title).Msg("Blog fetched successfully")

	return &blog, nil
}

func (r *database) UpdateBlog(blog *models.Blog) error {
	err := r.db.Save(blog).Error
	if err != nil {
		logger.Log.Error().Uint("blog id", blog.ID).Msg(fmt.Sprintf("Failed to update blog: %v", err))
		return err
	}

	logger.Log.Info().Uint("blog id", blog.ID).Msg("Blog updated successfully")

	return nil
}

func (r *database) DeleteBlog(id *uint) error {
	err := r.db.Delete(&models.Blog{}, id).Error
	if err != nil {
		logger.Log.Error().Uint("blog id", *id).Msg(fmt.Sprintf("Failed to delete blog: %v", err))
		return err
	}

	logger.Log.Info().Uint("blog id", *id).Msg("Blog deleted successfully")

	return nil
}

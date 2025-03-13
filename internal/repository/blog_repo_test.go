package repository

import (
	"blog-api/internal/models"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	return gormDB, mock, nil
}

func Test_database_CreateBlog(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	assert.NoError(t, err)
	defer gormDB.DB()

	repo := &database{db: gormDB}

	t.Run("failure", func(t *testing.T) {
		blog := &models.Blog{
			Title:       "Test Blog",
			Description: "Test Description",
			Body:        "Test Body",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		mock.ExpectBegin()
		mock.ExpectExec(`INSERT INTO "blogs"`).
			WithArgs(blog.Title, blog.Description, blog.Body, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(errors.New("failed to create blog"))
		mock.ExpectRollback()

		err := repo.CreateBlog(blog)
		assert.Error(t, err)
	})
}

func Test_database_GetAllBlogs(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	assert.NoError(t, err)
	defer gormDB.DB()

	repo := &database{db: gormDB}

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "title", "description", "body", "created_at", "updated_at"}).
			AddRow(1, "Test Blog", "Test Description", "Test Body", time.Now(), time.Now())

		mock.ExpectQuery(`SELECT \* FROM "blogs"`).WillReturnRows(rows)

		blogs, err := repo.GetAllBlogs()
		assert.NoError(t, err)
		assert.NotNil(t, blogs)
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectQuery(`SELECT \* FROM "blogs"`).WillReturnError(errors.New("failed to fetch blogs"))

		blogs, err := repo.GetAllBlogs()
		assert.Error(t, err)
		assert.Nil(t, blogs)
	})
}

func Test_database_GetBlogByID(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	assert.NoError(t, err)
	defer gormDB.DB()

	repo := &database{db: gormDB}

	t.Run("failure", func(t *testing.T) {
		mock.ExpectQuery(`SELECT \* FROM "blogs" WHERE "blogs"."id" = \$1`).WithArgs(1).WillReturnError(errors.New("blog not found"))

		id := uint(1)
		blog, err := repo.GetBlogByID(&id)
		assert.Error(t, err)
		assert.Nil(t, blog)
	})
}

func Test_database_UpdateBlog(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	assert.NoError(t, err)
	defer gormDB.DB()

	repo := &database{db: gormDB}

	t.Run("failure", func(t *testing.T) {
		blog := &models.Blog{
			ID:          1,
			Title:       "Updated Blog",
			Description: "Updated Description",
			Body:        "Updated Body",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "blogs" SET`).
			WithArgs(blog.Title, blog.Description, blog.Body, sqlmock.AnyArg(), blog.ID).
			WillReturnError(errors.New("failed to update blog"))
		mock.ExpectRollback()

		err := repo.UpdateBlog(blog)
		assert.Error(t, err)
	})
}

func Test_database_DeleteBlog(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	assert.NoError(t, err)
	defer gormDB.DB()

	repo := &database{db: gormDB}

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`DELETE FROM "blogs" WHERE "blogs"."id" = \$1`).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		id := uint(1)
		err := repo.DeleteBlog(&id)
		assert.NoError(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`DELETE FROM "blogs" WHERE "blogs"."id" = \$1`).WithArgs(1).WillReturnError(errors.New("failed to delete blog"))
		mock.ExpectRollback()

		id := uint(1)
		err := repo.DeleteBlog(&id)
		assert.Error(t, err)
	})
}

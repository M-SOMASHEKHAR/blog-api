package database

import (
	"blog-api/internal/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dbURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	sql, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = sql.Ping()
	if err != nil {
		return nil, err
	}
	// Auto Migrate Models
	err = db.AutoMigrate(&models.Blog{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Database connected successfully!")
	return db, nil
}

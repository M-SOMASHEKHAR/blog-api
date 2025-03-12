package database

import (
	"blog-api/internal/logger"
	"blog-api/internal/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {

	logger.Log.Info().Msg("establishing database connection")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	//postgres database source name
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("failed to establish database connection")
	}

	sql, err := db.DB()
	if err != nil {
		logger.Log.Error().Err(err).Msg("failed to get the database instance")
		return nil, err
	}

	err = sql.Ping()
	if err != nil {
		logger.Log.Error().Err(err).Msg("databse connection is not live")
		return nil, err
	}
	// Auto Migrate Models
	err = db.AutoMigrate(&models.Blog{})
	if err != nil {
		logger.Log.Error().Err(err).Msg("failed to create db tables")
		return nil, err
	}

	logger.Log.Info().Msg("database connection successfull")
	return db, nil
}

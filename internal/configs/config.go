package config

import (
	logger "blog-api/internal/logging"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl string
	Port  string
}

func LoadConfig() Config {

	logger.Log.Info().Msg("loading env file")

	envPath, err := filepath.Abs("../../")
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("failed to get current working directory")
	}
	envPath = filepath.Join(envPath, "internal", "configs", "application.env")
	err = godotenv.Load(envPath)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("failed to load env fail")
	}

	err = godotenv.Load(envPath)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("failed to load env fail")
	}

	logger.Log.Info().Msg("env loded successfully")
	return Config{
		DBUrl: os.Getenv("DATABASE_URL"),
		Port:  os.Getenv("PORT"),
	}
}

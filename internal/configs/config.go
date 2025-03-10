package config

import (
	logger "blog-api/internal/logging"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl string
	Port  string
}

func LoadConfig() Config {

	logger.Log.Info().Msg("loading env file")
	err := godotenv.Load()
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("failed to load env fail")
	}

	logger.Log.Info().Msg("env loded successfully")
	return Config{
		DBUrl: os.Getenv("DATABASE_URL"),
		Port:  os.Getenv("PORT"),
	}
}

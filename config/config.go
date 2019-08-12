package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	ENV      = os.Getenv("ENVIRONMENT")
	APP_PORT = os.Getenv("APP_PORT")
	CSV_PATH = os.Getenv("CSV_PATH")
)

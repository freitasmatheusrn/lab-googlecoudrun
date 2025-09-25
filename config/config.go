package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	WeatherKey string
	WeatherUrl string
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Aviso: não foi possível carregar o arquivo .env: %v", err)
	}

	return &Config{
		WeatherKey: os.Getenv("WEATHER_KEY"),
		WeatherUrl: os.Getenv("WEATHER_URL"),
	}
}

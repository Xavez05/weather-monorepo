package main

import "os"

type Config struct {
	OpenWeatherAPIKey string
	Port              string
}

func LoadConfig() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		OpenWeatherAPIKey: os.Getenv("OPENWEATHER_API_KEY"), // ya no es requerida
		Port:              port,
	}, nil
}

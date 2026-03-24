package main

import (
	"log"

	"github.com/Xavez05/weather-monorepo/weather-app/internal/config"
	"github.com/Xavez05/weather-monorepo/weather-app/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("error de configuración: %v", err)
	}

	if err := server.Start(cfg); err != nil {
		log.Fatalf("error en el servidor: %v", err)
	}
}

package server

import (
	"fmt"
	"log"
	"net/http"

	restclient "github.com/Xavez05/weather-monorepo/apiclient/rest"
	soapclient "github.com/Xavez05/weather-monorepo/apiclient/soap"
	"github.com/Xavez05/weather-monorepo/weather-app/internal/config"
	"github.com/Xavez05/weather-monorepo/weather-app/internal/handlers"
	"github.com/Xavez05/weather-monorepo/weather-app/internal/routes"
	"github.com/Xavez05/weather-monorepo/weather-app/internal/services"
)

// Start construye todas las dependencias y arranca el servidor
// Principio: acá vive el "wiring" — ensamblado de dependencias
func Start(cfg *config.Config) error {
	// Clientes de la librería interna
	rest := restclient.NewClient(cfg.OpenWeatherAPIKey)
	soap := soapclient.NewClient()

	// Servicio
	svc := services.NewWeatherService(rest, soap)

	// Handlers
	weatherHandler := handlers.NewWeatherHandler(svc)
	homeHandler := handlers.NewHomeHandler("internal/templates/home.html")

	// Rutas
	mux := http.NewServeMux()
	routes.Register(mux, homeHandler, weatherHandler)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Servidor corriendo en http://localhost%s", addr)
	return http.ListenAndServe(addr, mux)
}

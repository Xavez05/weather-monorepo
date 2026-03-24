package routes

import (
	"github.com/Xavez05/weather-monorepo/weather-app/internal/handlers"
	"net/http"
)

func Register(
	mux *http.ServeMux,
	home *handlers.HomeHandler,
	weather *handlers.WeatherHandler,
) {
	mux.HandleFunc("/", home.Home)
	mux.HandleFunc("/api/weather/rest", weather.GetWeatherREST)
	mux.HandleFunc("/api/weather/soap", weather.GetWeatherSOAP)
}

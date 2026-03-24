package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Xavez05/weather-monorepo/apiclient"
)

var svc *WeatherService

func StartServer(cfg *Config) error {
	rest := apiclient.NewClientRest("")
	soap := apiclient.NewClientSoap()
	svc = NewWeatherService(rest, soap)

	mux := http.NewServeMux()
	Register(mux)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Servidor corriendo en http://localhost%s", addr)
	return http.ListenAndServe(addr, mux)
}

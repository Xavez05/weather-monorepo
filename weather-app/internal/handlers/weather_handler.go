package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Xavez05/weather-monorepo/weather-app/internal/services"
)

type WeatherHandler struct {
	svc *services.WeatherService
}

func NewWeatherHandler(svc *services.WeatherService) *WeatherHandler {
	return &WeatherHandler{svc: svc}
}

func (h *WeatherHandler) GetWeatherREST(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		respondError(w, "city es requerido", http.StatusBadRequest)
		return
	}

	result, err := h.svc.GetByCity(city)
	if err != nil {
		respondError(w, err.Error(), http.StatusBadGateway)
		return
	}

	respondJSON(w, result, http.StatusOK)
}

func (h *WeatherHandler) GetWeatherSOAP(w http.ResponseWriter, r *http.Request) {
	country := r.URL.Query().Get("country")
	if country == "" {
		respondError(w, "country es requerido (ej: GT, US, MX)", http.StatusBadRequest)
		return
	}

	result, err := h.svc.GetByCountry(country)
	if err != nil {
		respondError(w, err.Error(), http.StatusBadGateway)
		return
	}

	respondJSON(w, result, http.StatusOK)
}

// helpers privados — Single Responsibility: respuestas HTTP centralizadas
func respondJSON(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, message string, status int) {
	respondJSON(w, map[string]string{"error": message}, status)
}

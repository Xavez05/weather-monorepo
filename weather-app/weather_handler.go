package main

import (
	"encoding/json"
	"net/http"
)

func GetWeatherRESTHandler(w http.ResponseWriter, r *http.Request) {
	var req CityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.City == "" {
		respondError(w, "body inválido, se esperaba {\"city\": \"...\"}", http.StatusBadRequest)
		return
	}

	result, err := svc.GetByCity(req.City)
	if err != nil {
		respondError(w, err.Error(), http.StatusBadGateway)
		return
	}

	respondJSON(w, result, http.StatusOK)
}

func GetWeatherSOAPHandler(w http.ResponseWriter, r *http.Request) {
	var req CountryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Country == "" {
		respondError(w, "body inválido, se esperaba {\"country\": \"...\"}", http.StatusBadRequest)
		return
	}

	result, err := svc.GetByCountry(req.Country)
	if err != nil {
		respondError(w, err.Error(), http.StatusBadGateway)
		return
	}

	respondJSON(w, result, http.StatusOK)
}

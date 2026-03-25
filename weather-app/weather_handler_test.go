package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Xavez05/weather-monorepo/apiclient"
	"github.com/stretchr/testify/assert"
)

func setupSvc(mock *mockFetcher) {
	svc = NewWeatherService(mock, mock)
}

func TestGetWeatherRESTHandler_Success(t *testing.T) {
	setupSvc(&mockFetcher{
		restResponse: &apiclient.WeatherResponse{
			City:        "Guatemala",
			Temperature: 22.5,
			Source:      "REST",
		},
	})

	body := bytes.NewBufferString(`{"city": "Guatemala"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/weather/rest", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	GetWeatherRESTHandler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var result apiclient.WeatherResponse
	json.NewDecoder(w.Body).Decode(&result)
	assert.Equal(t, "Guatemala", result.City)
	assert.Equal(t, 22.5, result.Temperature)
}

func TestGetWeatherRESTHandler_EmptyCity(t *testing.T) {
	setupSvc(&mockFetcher{})

	body := bytes.NewBufferString(`{"city": ""}`)
	req := httptest.NewRequest(http.MethodPost, "/api/weather/rest", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	GetWeatherRESTHandler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetWeatherRESTHandler_ServiceError(t *testing.T) {
	setupSvc(&mockFetcher{
		restErr: errors.New("ciudad no encontrada"),
	})

	body := bytes.NewBufferString(`{"city": "CiudadInexistente"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/weather/rest", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	GetWeatherRESTHandler(w, req)

	assert.Equal(t, http.StatusBadGateway, w.Code)
}

func TestGetWeatherSOAPHandler_Success(t *testing.T) {
	setupSvc(&mockFetcher{
		soapResponse: &apiclient.WeatherResponse{
			City:   "Guatemala City",
			Source: "SOAP",
		},
		restResponse: &apiclient.WeatherResponse{
			City:        "Guatemala City",
			Temperature: 19.8,
			Source:      "REST",
		},
	})

	body := bytes.NewBufferString(`{"country": "GT"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/weather/soap", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	GetWeatherSOAPHandler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetWeatherSOAPHandler_EmptyCountry(t *testing.T) {
	setupSvc(&mockFetcher{})

	body := bytes.NewBufferString(`{"country": ""}`)
	req := httptest.NewRequest(http.MethodPost, "/api/weather/soap", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	GetWeatherSOAPHandler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetWeatherSOAPHandler_ServiceError(t *testing.T) {
	setupSvc(&mockFetcher{
		soapErr: errors.New("código de país no encontrado"),
	})

	body := bytes.NewBufferString(`{"country": "XX"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/weather/soap", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	GetWeatherSOAPHandler(w, req)

	assert.Equal(t, http.StatusBadGateway, w.Code)
}

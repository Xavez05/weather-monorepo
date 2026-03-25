package main

import (
	"errors"
	"testing"

	"github.com/Xavez05/weather-monorepo/apiclient"
	"github.com/stretchr/testify/assert"
)

type mockFetcher struct {
	restResponse *apiclient.WeatherResponse
	soapResponse *apiclient.WeatherResponse
	restErr      error
	soapErr      error
}

func (m *mockFetcher) GetWeatherRest(_ string) (*apiclient.WeatherResponse, error) {
	return m.restResponse, m.restErr
}

func (m *mockFetcher) GetWeatherSoap(_ string) (*apiclient.WeatherResponse, error) {
	return m.soapResponse, m.soapErr
}

func TestGetByCity_Success(t *testing.T) {
	mock := &mockFetcher{
		restResponse: &apiclient.WeatherResponse{
			City:        "Guatemala",
			Country:     "GT",
			Temperature: 22.5,
			Humidity:    75,
			Description: "parcialmente nublado",
			Source:      "REST",
		},
	}

	svc := NewWeatherService(mock, mock)
	result, err := svc.GetByCity("Guatemala")

	assert.NoError(t, err)
	assert.Equal(t, "Guatemala", result.City)
	assert.Equal(t, 22.5, result.Temperature)
	assert.Equal(t, "REST", result.Source)
}

func TestGetByCity_Error(t *testing.T) {
	mock := &mockFetcher{
		restErr: errors.New("ciudad no encontrada"),
	}

	svc := NewWeatherService(mock, mock)
	result, err := svc.GetByCity("CiudadInexistente")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "ciudad no encontrada")
}

func TestGetByCountry_Success(t *testing.T) {
	mock := &mockFetcher{
		soapResponse: &apiclient.WeatherResponse{
			City:   "Guatemala City",
			Source: "SOAP",
		},
		restResponse: &apiclient.WeatherResponse{
			City:        "Guatemala City",
			Country:     "GT",
			Temperature: 19.8,
			Humidity:    68,
			Description: "cielo despejado",
			Source:      "REST",
		},
	}

	svc := NewWeatherService(mock, mock)
	result, err := svc.GetByCountry("GT")

	assert.NoError(t, err)
	assert.Equal(t, "GT", result.Country)
	assert.Equal(t, 19.8, result.Temperature)
	assert.Equal(t, "SOAP + REST", result.Source)
	assert.Contains(t, result.Description, "Guatemala City")
}

func TestGetByCountry_SOAPError(t *testing.T) {
	mock := &mockFetcher{
		soapErr: errors.New("código de país no encontrado"),
	}

	svc := NewWeatherService(mock, mock)
	result, err := svc.GetByCountry("XX")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "código de país no encontrado")
}

func TestGetByCountry_RESTError(t *testing.T) {
	mock := &mockFetcher{
		soapResponse: &apiclient.WeatherResponse{
			City:   "Guatemala City",
			Source: "SOAP",
		},
		restErr: errors.New("ciudad no encontrada"),
	}

	svc := NewWeatherService(mock, mock)
	result, err := svc.GetByCountry("GT")

	assert.Error(t, err)
	assert.Nil(t, result)
}

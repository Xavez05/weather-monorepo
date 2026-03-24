package apiclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	geoURL     = "https://geocoding-api.open-meteo.com/v1/search"
	weatherURL = "https://api.open-meteo.com/v1/forecast"
)

// NewClient ya no necesita APIKey — Open-Meteo es sin autenticación
func NewClientRest(_ string) *Client {
	return &Client{
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) GetWeatherRest(city string) (*WeatherResponse, error) {
	lat, lon, country, err := c.geocode(city)
	if err != nil {
		return nil, err
	}

	return c.fetchWeather(city, country, lat, lon)
}

func (c *Client) geocode(city string) (lat, lon float64, country string, err error) {
	url := fmt.Sprintf("%s?name=%s&count=1&language=es&format=json", geoURL, city)

	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return 0, 0, "", NewAPIError("REST", err.Error(), 0)
	}
	defer resp.Body.Close()

	var result struct {
		Results []struct {
			Latitude    float64 `json:"latitude"`
			Longitude   float64 `json:"longitude"`
			CountryCode string  `json:"country_code"`
		} `json:"results"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, 0, "", NewAPIError("REST", "error parseando geocoding", 0)
	}

	if len(result.Results) == 0 {
		return 0, 0, "", NewAPIError("REST", "ciudad no encontrada", 404)
	}

	r := result.Results[0]
	return r.Latitude, r.Longitude, r.CountryCode, nil
}

func (c *Client) fetchWeather(city, country string, lat, lon float64) (*WeatherResponse, error) {
	url := fmt.Sprintf(
		"%s?latitude=%f&longitude=%f&current=temperature_2m,relative_humidity_2m,apparent_temperature,weather_code&timezone=auto",
		weatherURL, lat, lon,
	)

	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, NewAPIError("REST", err.Error(), 0)
	}
	defer resp.Body.Close()

	var raw struct {
		Current struct {
			Temperature float64 `json:"temperature_2m"`
			FeelsLike   float64 `json:"apparent_temperature"`
			Humidity    int     `json:"relative_humidity_2m"`
			WeatherCode int     `json:"weather_code"`
		} `json:"current"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, NewAPIError("REST", "error parseando clima", 0)
	}

	return &WeatherResponse{
		City:        city,
		Country:     country,
		Temperature: raw.Current.Temperature,
		FeelsLike:   raw.Current.FeelsLike,
		Humidity:    raw.Current.Humidity,
		Description: weatherCodeToDescription(raw.Current.WeatherCode),
		Source:      "REST",
	}, nil
}

// weatherCodeToDescription convierte el código WMO a texto legible
func weatherCodeToDescription(code int) string {
	switch {
	case code == 0:
		return "cielo despejado"
	case code <= 3:
		return "parcialmente nublado"
	case code <= 49:
		return "niebla"
	case code <= 59:
		return "llovizna"
	case code <= 69:
		return "lluvia"
	case code <= 79:
		return "nieve"
	case code <= 84:
		return "chubascos"
	case code <= 99:
		return "tormenta eléctrica"
	default:
		return "sin descripción"
	}
}

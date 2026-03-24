package apiclient

import "net/http"

type WeatherResponse struct {
	City        string  `json:"city"`
	Country     string  `json:"country"`
	Temperature float64 `json:"temperature"`
	FeelsLike   float64 `json:"feels_like"`
	Description string  `json:"description"`
	Humidity    int     `json:"humidity"`
	Source      string  `json:"source"`
}

type Client struct {
	HTTPClient *http.Client
}

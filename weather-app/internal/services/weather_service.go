package services

import (
	"github.com/Xavez05/weather-monorepo/apiclient/models"
)

type WeatherFetcher interface {
	GetWeather(query string) (*models.WeatherResponse, error)
}

type WeatherService struct {
	restFetcher WeatherFetcher
	soapFetcher WeatherFetcher
}

func NewWeatherService(rest, soap WeatherFetcher) *WeatherService {
	return &WeatherService{
		restFetcher: rest,
		soapFetcher: soap,
	}
}

func (s *WeatherService) GetByCity(city string) (*models.WeatherResponse, error) {
	return s.restFetcher.GetWeather(city)
}

func (s *WeatherService) GetByCountry(countryCode string) (*models.WeatherResponse, error) {

	countryInfo, err := s.soapFetcher.GetWeather(countryCode)
	if err != nil {
		return nil, err
	}

	weather, err := s.restFetcher.GetWeather(countryInfo.City)
	if err != nil {
		return nil, err
	}

	weather.Country = countryCode
	weather.Description = weather.Description + " en " + countryInfo.City
	weather.Source = "SOAP + REST"

	return weather, nil
}

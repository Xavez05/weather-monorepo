package main

import (
	"github.com/Xavez05/weather-monorepo/apiclient"
)

type WeatherFetcher interface {
	GetWeatherSoap(query string) (*apiclient.WeatherResponse, error)
	GetWeatherRest(query string) (*apiclient.WeatherResponse, error)
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

func (s *WeatherService) GetByCity(city string) (*apiclient.WeatherResponse, error) {
	return s.restFetcher.GetWeatherRest(city)
}

func (s *WeatherService) GetByCountry(countryCode string) (*apiclient.WeatherResponse, error) {

	countryInfo, err := s.soapFetcher.GetWeatherSoap(countryCode)
	if err != nil {
		return nil, err
	}

	weather, err := s.restFetcher.GetWeatherRest(countryInfo.City)
	if err != nil {
		return nil, err
	}

	weather.Country = countryCode
	weather.Description = weather.Description + " en " + countryInfo.City
	weather.Source = "SOAP + REST"

	return weather, nil
}

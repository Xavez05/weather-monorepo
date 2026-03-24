package soap

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/http"
	"time"

	apierrors "github.com/Xavez05/weather-monorepo/apiclient/errors"
	"github.com/Xavez05/weather-monorepo/apiclient/models"
)

const soapURL = "http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso"

type Client struct {
	HTTPClient *http.Client
}

func NewClient() *Client {
	return &Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) GetWeather(countryCode string) (*models.WeatherResponse, error) {
	body := `<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <CapitalCity xmlns="http://www.oorsprong.org/websamples.countryinfo">
      <sCountryISOCode>` + countryCode + `</sCountryISOCode>
    </CapitalCity>
  </soap:Body>
</soap:Envelope>`

	req, err := http.NewRequest("POST", soapURL, bytes.NewBufferString(body))
	if err != nil {
		return nil, apierrors.NewAPIError("SOAP", err.Error(), 0)
	}
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", "")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, apierrors.NewAPIError("SOAP", err.Error(), 0)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, apierrors.NewAPIError("SOAP", "error leyendo respuesta", 0)
	}

	var envelope CountryEnvelope
	if err := xml.Unmarshal(data, &envelope); err != nil {
		return nil, apierrors.NewAPIError("SOAP", "error parseando XML", 0)
	}

	capital := envelope.Body.Response.Capital
	if capital == "" {
		return nil, apierrors.NewAPIError("SOAP", "código de país no encontrado", 404)
	}

	return &models.WeatherResponse{
		City:        capital,
		Country:     countryCode,
		Temperature: 0,
		Description: "Capital obtenida vía SOAP",
		Source:      "SOAP",
	}, nil
}

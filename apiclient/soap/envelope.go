package soap

import "encoding/xml"

type CountryEnvelope struct {
	XMLName xml.Name    `xml:"Envelope"`
	Body    CountryBody `xml:"Body"`
}

type CountryBody struct {
	Response CapitalCityResponse `xml:"CapitalCityResponse"`
}

type CapitalCityResponse struct {
	Capital string `xml:"CapitalCityResult"`
}

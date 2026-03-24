package main

import "net/http"

func Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/weather/rest", GetWeatherRESTHandler)
	mux.HandleFunc("POST /api/weather/soap", GetWeatherSOAPHandler)
}

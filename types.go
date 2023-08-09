package main

type apiConfigData struct {
	OpenWeatherMapApiKey string
}

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/weather/", WeatherHandler)

	http.ListenAndServe(":8080", nil)
}

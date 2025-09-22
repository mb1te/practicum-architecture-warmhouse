package main

import (
	"log"
	"net/http"

	"github.com/mb1te/practicum-architecture-warmhouse/internal/domain/temperature"
	api "github.com/mb1te/practicum-architecture-warmhouse/internal/http"
	"github.com/mb1te/practicum-architecture-warmhouse/internal/http/handlers"
)

func main() {
	service := temperature.NewRandomService()
	handler := handlers.NewTemperatureHandler(service)
	router := api.NewRouter(handler)

	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Fatal(err)
	}
}

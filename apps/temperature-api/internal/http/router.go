package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mb1te/practicum-architecture-warmhouse/internal/http/handlers"
)

func NewRouter(h *handlers.TemperatureHandler) http.Handler {
	r := chi.NewRouter()

	r.Get("/temperature", h.GetTemperatureByLocation)
	r.Get("/temperature/{sensorID}", h.GetTemperatureBySensorID)

	return r
}

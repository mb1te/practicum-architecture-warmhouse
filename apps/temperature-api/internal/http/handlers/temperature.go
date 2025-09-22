package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mb1te/practicum-architecture-warmhouse/internal/domain/temperature"
)

type TemperatureHandler struct {
	service temperature.Service
}

func NewTemperatureHandler(service temperature.Service) *TemperatureHandler {
	return &TemperatureHandler{service: service}
}

func (h *TemperatureHandler) GetTemperatureByLocation(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")

	resp, err := h.service.GetTemperature(r.Context(), location, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	makeResponse(w, http.StatusOK, resp)
}

func (h *TemperatureHandler) GetTemperatureBySensorID(w http.ResponseWriter, r *http.Request) {
	sensorID := chi.URLParam(r, "sensorID")

	resp, err := h.service.GetTemperature(r.Context(), "", sensorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	makeResponse(w, http.StatusOK, resp)
}

func makeResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

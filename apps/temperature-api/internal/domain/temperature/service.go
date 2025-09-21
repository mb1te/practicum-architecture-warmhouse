package temperature

import (
	"context"
	"math/rand/v2"
	"time"
)

type Service interface {
	GetTemperature(ctx context.Context, location string, serviceID string) (Response, error)
}

type RandomService struct{}

func NewRandomService() *RandomService {
	return &RandomService{}
}

func (s *RandomService) GetTemperature(ctx context.Context, location string, serviceID string) (Response, error) {
	if location == "" {
		location = sensorIDToLocation(serviceID)
	}

	if serviceID == "" {
		serviceID = locationToSensorID(location)
	}

	r := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano())))
	temperature := -100 + r.Float64()*200

	return Response{
		Value:     temperature,
		Timestamp: time.Now(),
		Location:  location,
		SensorID:  serviceID,
	}, nil
}

func sensorIDToLocation(sensorID string) string {
	switch sensorID {
	case "1":
		return "Living Room"
	case "2":
		return "Bedroom"
	case "3":
		return "Kitchen"
	default:
		return "Unknown"
	}
}

func locationToSensorID(location string) string {
	switch location {
	case "Living Room":
		return "1"
	case "Bedroom":
		return "2"
	case "Kitchen":
		return "3"
	default:
		return "0"
	}
}

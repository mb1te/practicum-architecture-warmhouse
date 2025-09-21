package temperature

import "time"

type Response struct {
	Value     float64   `json:"value"`
	Timestamp time.Time `json:"timestamp"`
	Location  string    `json:"location"`
	SensorID  string    `json:"sensor_id"`
}

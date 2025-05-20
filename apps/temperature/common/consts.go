package common

import "time"

const (
	InputParamNotFound = "input param: %s, not found. error: %s"
)

// TemperatureResponse represents the response from the temperature API
type TemperatureResponse struct {
	Name        string    `json:"name"`
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`
	SensorType  string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	LastUpdated time.Time `json:"last_updated"`
}

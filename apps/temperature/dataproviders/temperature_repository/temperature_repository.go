package temperature_repository

import (
	"context"
	"fmt"
	"math/rand"
	"temperature/common"
	"temperature/common/pg"
)

type ITemperature interface {
	// GetTemperatureBySensorId получение температуры
	GetTemperatureBySensorId(ctx context.Context, sensorId int) (common.TemperatureResponse, error)
	// GetTemperatureBySensors получение температур
	GetTemperatureBySensors(ctx context.Context) ([]common.TemperatureResponse, error)
}

type userRepository struct {
	pg *pg.DB
}

// NewTemperatureRepository получить новый экземпляр
func NewTemperatureRepository(pg *pg.DB) ITemperature {
	return &userRepository{pg: pg}
}

// GetTemperatureBySensorId получение температуры
func (c *userRepository) GetTemperatureBySensorId(ctx context.Context, sensorId int) (common.TemperatureResponse, error) {
	var res common.TemperatureResponse

	query := `
		select id as sensor_id, name, type, location, value, unit, status, last_updated, created_at 
		from sensors where id = $1 limit 1
	`

	rows, err := c.pg.Pool.Query(ctx, query, sensorId)
	if err != nil {
		return common.TemperatureResponse{}, fmt.Errorf("error querying sensors: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&res.SensorID,
			&res.Name,
			&res.SensorType,
			&res.Location,
			&res.Value,
			&res.Unit,
			&res.Status,
			&res.LastUpdated,
			&res.CreatedAt,
		)
		if err != nil {
			return common.TemperatureResponse{}, fmt.Errorf("error scanning sensor row: %w", err)
		}
	}

	if err := rows.Err(); err != nil {
		return common.TemperatureResponse{}, fmt.Errorf("error iterating sensor rows: %w", err)
	}

	res.Value = rand.Float64()

	return res, nil
}

// GetTemperatureBySensors получение температур
func (c *userRepository) GetTemperatureBySensors(ctx context.Context) ([]common.TemperatureResponse, error) {
	res := make([]common.TemperatureResponse, 0, 10)

	query := `
		select id as sensor_id, name, type, location, value, unit, status, last_updated, created_at
		from sensors
	`

	rows, err := c.pg.Pool.Query(ctx, query)
	if err != nil {
		return []common.TemperatureResponse{}, fmt.Errorf("error querying sensors: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tempr common.TemperatureResponse
		err := rows.Scan(
			&tempr.SensorID,
			&tempr.Name,
			&tempr.SensorType,
			&tempr.Location,
			&tempr.Value,
			&tempr.Unit,
			&tempr.Status,
			&tempr.LastUpdated,
			&tempr.CreatedAt,
		)
		if err != nil {
			return []common.TemperatureResponse{}, fmt.Errorf("error scanning sensor row: %w", err)
		}
		res = append(res, tempr)
	}

	if err := rows.Err(); err != nil {
		return []common.TemperatureResponse{}, fmt.Errorf("error iterating sensor rows: %w", err)
	}

	return res, nil
}

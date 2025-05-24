package temperature

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"temperature/common"
	"temperature/dataproviders/temperature_repository"
)

type ITemperatures interface {
	DoGetTemperatures(ctx echo.Context) (common.SensorTemperatureList, error)
}

type temperatures struct {
	temperatureRepository temperature_repository.ITemperature
}

func NewTemperatures(temperatureRepository temperature_repository.ITemperature) ITemperatures {
	return &temperatures{
		temperatureRepository: temperatureRepository,
	}
}

// DoGetTemperatures получить список сенсоров и температур
// @Summary получить температуру по идентификатору сенсора
// @Produce      json
// @Success      200 {object}  common.SensorTemperatureList
// @Router       /temperatures [get]
func (t *temperatures) DoGetTemperatures(ctx echo.Context) (common.SensorTemperatureList, error) {
	appCtx := ctx.Request().Context()

	tmprs, err := t.temperatureRepository.GetTemperatureBySensors(appCtx)
	if err != nil {
		slog.Error(err.Error())

		return common.SensorTemperatureList{}, common.ReturnInternalError(ctx, err, "")
	}

	return tmprs, nil
}

package temperature

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log/slog"
	"temperature/common"
	"temperature/dataproviders/temperature_repository"
)

type ITemperature interface {
	DoGetTemperatureById(ctx echo.Context) (common.TemperatureResponse, error)
}

type temperature struct {
	temperatureRepository temperature_repository.ITemperature
}

func NewTemperature(temperatureRepository temperature_repository.ITemperature) ITemperature {
	return &temperature{
		temperatureRepository: temperatureRepository,
	}
}

func (t *temperature) DoGetTemperatureById(ctx echo.Context) (common.TemperatureResponse, error) {
	appCtx := ctx.Request().Context()
	sensorId, err := common.GetPathParamByName(ctx, Id)
	if err != nil {
		fmt.Println(err)

		return common.TemperatureResponse{}, common.ReturnInternalError(ctx, err, "")
	}

	tmpr, err := t.temperatureRepository.GetTemperatureBySensorId(appCtx, sensorId)
	if err != nil {
		slog.Error(err.Error())

		return common.TemperatureResponse{}, common.ReturnInternalError(ctx, err, "")
	}

	return tmpr, nil
}

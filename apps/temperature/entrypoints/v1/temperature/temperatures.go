package temperature

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"temperature/common"
	"temperature/dataproviders/temperature_repository"
)

type ITemperatures interface {
	DoGetTemperatures(ctx echo.Context) ([]common.TemperatureResponse, error)
}

type temperatures struct {
	temperatureRepository temperature_repository.ITemperature
}

func NewTemperatures(temperatureRepository temperature_repository.ITemperature) ITemperatures {
	return &temperatures{
		temperatureRepository: temperatureRepository,
	}
}

func (t *temperatures) DoGetTemperatures(ctx echo.Context) ([]common.TemperatureResponse, error) {
	appCtx := ctx.Request().Context()

	tmprs, err := t.temperatureRepository.GetTemperatureBySensors(appCtx)
	if err != nil {
		slog.Error(err.Error())

		return []common.TemperatureResponse{}, common.ReturnInternalError(ctx, err, "")
	}

	return tmprs, nil
}

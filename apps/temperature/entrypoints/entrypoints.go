package entrypoints

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"temperature/common"
	dataproviders2 "temperature/dataproviders"
	"temperature/entrypoints/v1/temperature"
)

const (
	Healthcheck = "/system/healthcheck"
)

type IEndpoints interface {
	HealthCheck(ctx echo.Context) error
	GetTemperatureById(ctx echo.Context) error
	GetTemperatures(ctx echo.Context) error
}

type endpoints struct {
	Providers dataproviders2.CoreProviders
}

func NewEndpoints(providers dataproviders2.CoreProviders) IEndpoints {
	return &endpoints{
		Providers: providers,
	}
}

func (e *endpoints) HealthCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &common.Status{Status: "OK"})
}

func (e *endpoints) GetTemperatureById(ctx echo.Context) error {
	temperatureEntripoints := temperature.NewTemperature(e.Providers.GetUserRepository())
	tmpr, err := temperatureEntripoints.DoGetTemperatureById(ctx)
	if err != nil {
		slog.Error(err.Error())

		return common.ReturnInternalError(ctx, err, "")
	}

	return ctx.JSON(http.StatusOK, &tmpr)
}

func (e *endpoints) GetTemperatures(ctx echo.Context) error {
	temperatureEntripoints := temperature.NewTemperatures(e.Providers.GetUserRepository())
	tmpr, err := temperatureEntripoints.DoGetTemperatures(ctx)
	if err != nil {
		slog.Error(err.Error())

		return common.ReturnInternalError(ctx, err, "")
	}

	return ctx.JSON(http.StatusOK, &tmpr)
}

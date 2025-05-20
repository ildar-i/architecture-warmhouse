package main

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"net/http"
	"os"
	"temperature/common/pg"
	"temperature/dataproviders"
	"temperature/entrypoints"
	"temperature/entrypoints/v1/temperature"
)

const (
	appName    = "temperature"
	appVersion = "v1.0"
	appPort    = "8081"
)

func main() {
	//logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	//config := appconfig.NewConfig(appName, appVersion, appPort)
	//err := config.LoadConfig()
	//if err != nil {
	//	logger.Error(err.Error())
	//}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	dbURL := getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/smarthome")

	newPg := pg.NewPostgres(dbURL)
	if err := newPg.InitPg(); err != nil {
		slog.Error("pg init", "error", err)

		return
	}

	defer newPg.Db.Close()

	providers, err := dataproviders.NewCoreProviders(newPg.Db)
	if err != nil {
		slog.Error("failed to init providers", "error", err)

		return
	}

	endp := entrypoints.NewEndpoints(providers)

	// Routes
	e.GET(entrypoints.Healthcheck, endp.HealthCheck)
	e.GET(temperature.GetTemperatureById, endp.GetTemperatureById)
	e.GET(temperature.GetTemperatures, endp.GetTemperatures)

	// Start server
	if err := e.Start(fmt.Sprintf(":%s", appPort)); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

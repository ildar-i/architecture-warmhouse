package main

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	eSwagger "github.com/swaggo/echo-swagger"
	"log/slog"
	"net/http"
	"os"
	"temperature/common"
	"temperature/common/pg"
	"temperature/dataproviders"
	_ "temperature/docs"
	"temperature/entrypoints"
	"temperature/entrypoints/v1/temperature"
)

const (
	appPort = "8081"
)

// @title Warm House
// @version 1.0
// @description Warm House API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8081
// @BasePath /
func main() {
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
	e.GET(common.SwaggerRoute, eSwagger.WrapHandler)

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

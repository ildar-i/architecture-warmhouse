package common

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// ErrorHttp описание ошибки
type ErrorHttp struct {
	Code      int      `json:"code"`
	ErrorText string   `json:"errorText"`
	Details   []string `json:"details"`
}

// ReturnInternalError вернуть детализированную ошибку 500
func ReturnInternalError(ctx echo.Context, err error, detail string) error {
	return ctx.JSON(http.StatusInternalServerError, ErrorHttp{
		ErrorText: fmt.Sprintf("%s", err),
		Details:   []string{detail},
	})

}

// GetPathParamByName получить переменную пути по ее именованию
func GetPathParamByName(ctx echo.Context, key string) (int, error) {
	res := ctx.Param(key)
	if len(res) == 0 {
		return 0, fmt.Errorf(InputParamNotFound, key, "")
	}

	out, err := strconv.Atoi(res)
	if err != nil {
		return 0, err
	}

	return out, nil
}

type Status struct {
	Status string `json:"Status"`
}

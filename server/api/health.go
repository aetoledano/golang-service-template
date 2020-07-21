package api

import (
	"github.com/labstack/echo"
	"net/http"
)

func (_ *Api) HealthCheck(c echo.Context, pathParams map[string]string) error {
	return c.String(http.StatusOK, "OK")
}

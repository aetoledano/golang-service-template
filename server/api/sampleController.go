package api

import (
	"github.com/aetoledano/golang-service-template/models"
	"github.com/labstack/echo"
)

func (_ *Api) SampleHandler(c echo.Context, pathParams map[string]string) error {
	return ok(c, &models.ApiMessage{Data: "sample found"})
}

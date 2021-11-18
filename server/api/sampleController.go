package api

import (
	"github.com/aetoledano/golang-service-template/models"
	"github.com/labstack/echo"
)

func (_ *Api) SampleHandler(c echo.Context) error {
	return ok(c, &models.ApiMessage{Data: "sample found"})
}

func (_ *Api) PostHandler(c echo.Context) error {
	return ok(c, &models.ApiMessage{Data: c.QueryParams()})
}

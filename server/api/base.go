package api

import (
	"github.com/aetoledano/golang-service-template/models"
	"github.com/labstack/echo"
	"net/http"
)

type Api struct{}

func ok(c echo.Context, msg *models.ApiMessage) error {
	msg.Code = http.StatusOK
	return c.JSON(msg.Code, msg)
}

func internal(c echo.Context, msg *models.ApiMessage) error {
	msg.Code = http.StatusInternalServerError
	return c.JSON(msg.Code, msg)
}

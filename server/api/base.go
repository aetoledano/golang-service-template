package api

import (
	"github.com/labstack/echo"
	"github.com/aetoledano/golang-service-template/config"
	"github.com/aetoledano/golang-service-template/models"
	"net/http"
)

type Api struct{}

func response(c echo.Context, err *models.BusinessError, result interface{}) error {
	if err != nil {
		return statusError(c, err, result)
	}
	msg := new(models.ApiMessage)
	msg.Code = http.StatusOK
	msg.Tag = config.OK
	msg.Data = result
	return c.JSON(msg.Code, msg)
}

func statusNoContent(c echo.Context, err *models.BusinessError, result interface{}) error {
	if err != nil {
		return statusError(c, err, result)
	}
	msg := new(models.ApiMessage)
	msg.Code = http.StatusNoContent
	msg.Tag = config.NOCONTENT
	return c.JSON(msg.Code, msg)
}

func statusCreated(c echo.Context, err *models.BusinessError, result interface{}) error {
	if err != nil {
		return statusError(c, err, result)
	}
	msg := new(models.ApiMessage)
	msg.Code = http.StatusCreated
	msg.Tag = config.CREATED
	msg.Data = result
	return c.JSON(msg.Code, msg)
}

func statusError(c echo.Context, err *models.BusinessError, result interface{}) error {
	msg := new(models.ApiMessage)
	msg.Code = err.Code
	msg.Tag = err.Tag
	msg.Data = result
	return c.JSON(msg.Code, msg)
}

func badRequest(c echo.Context, result interface{}) error {
	msg := new(models.ApiMessage)
	msg.Code = http.StatusBadRequest
	msg.Tag = config.BADREQUEST
	msg.Data = result
	return c.JSON(msg.Code, msg)
}

func validationError(c echo.Context, result interface{}) error {
	msg := new(models.ApiMessage)
	msg.Code = http.StatusBadRequest
	msg.Tag = config.VALIDATIONERRORS
	msg.Data = result
	return c.JSON(msg.Code, msg)
}

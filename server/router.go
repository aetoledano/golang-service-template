package server

import (
	"github.com/aetoledano/golang-service-template/config"
	"github.com/aetoledano/golang-service-template/constants"
	"github.com/aetoledano/golang-service-template/models"
	controllers "github.com/aetoledano/golang-service-template/server/api"
	oa3 "github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
	"reflect"
	"strings"
)

var router *oa3.Router
var api = &controllers.Api{}

func configRoutes(e *echo.Echo) {

	e.File(constants.API_DOCS_PATH, constants.API_DOCS_FILE)

	e.GET(constants.HEALTHCHECK_PATH, healthCheck)

	configOpenApiRouter(e)
}

func configOpenApiRouter(e *echo.Echo) {
	router = oa3.NewRouter().WithSwaggerFromFile(constants.API_DOCS_FILE)

	//root group, middlewares with global scope should be registered here
	root := e.Group(config.App.Api.Root)

	root.Any("/*", openApiRoutingHandler)
}

func openApiRoutingHandler(c echo.Context) error {

	route, pathParams, e := router.FindRoute(c.Request().Method, c.Request().URL)
	if e != nil {
		var routeError = e.(*oa3.RouteError)

		errorMsg := "routing error: no route matches " + c.Request().Method + " " + c.Request().URL.Path
		log.Error(errorMsg, " ", "reason: "+routeError.Reason)

		msg := new(models.ApiMessage)
		msg.Code = http.StatusNotFound
		msg.Data = errorMsg
		return c.JSON(msg.Code, msg)
	}

	handlerName := strings.Title(route.Operation.OperationID)
	targetApiHandler := reflect.ValueOf(api).MethodByName(handlerName)
	if !targetApiHandler.IsValid() {
		msg := new(models.ApiMessage)
		msg.Code = http.StatusInternalServerError
		msg.Data = "no handler found for operationId: " + handlerName
		return c.JSON(msg.Code, msg)
	}

	apiHandlerFunc, typeCheckOk := targetApiHandler.Interface().(func(echo.Context, map[string]string) error)
	if !typeCheckOk {
		log.Error(
			"wrong handler signature for operationId: " +
				route.Operation.OperationID + ", not type: func(echo.Context, map[string]string) error")
		msg := new(models.ApiMessage)
		msg.Data = "wrong handler signature for operationId: " + route.Operation.OperationID
		msg.Code = http.StatusInternalServerError
		return c.JSON(msg.Code, msg)
	}

	return apiHandlerFunc(c, pathParams)
}

func healthCheck(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "OK")
}

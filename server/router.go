package server

import (
	"github.com/aetoledano/golang-service-template/config"
	"github.com/aetoledano/golang-service-template/models"
	controllers "github.com/aetoledano/golang-service-template/server/api"
	oa3 "github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
	"reflect"
)

var router *oa3.Router
var api = &controllers.Api{}

func configRouter(e *echo.Echo) {
	router = oa3.NewRouter().WithSwaggerFromFile("api-definition.yml")
	//root group, middleware with global scope should be registered here
	//root := e.Group(config.API_ROOT, RoleSecurityMiddleware)
	root := e.Group(config.API_ROOT)
	root.Any("/*", openApiRoutinghandler)
}

func openApiRoutinghandler(c echo.Context) error {

	route, pathParams, e := router.FindRoute(c.Request().Method, c.Request().URL)
	if e != nil {
		var routeError = e.(*oa3.RouteError)

		errorMsg := "routing error: no route matches " + c.Request().Method + " " + c.Request().URL.Path
		log.Error(errorMsg, " ","router: "+routeError.Reason)

		msg := new(models.ApiMessage)
		msg.Code = http.StatusNotFound
		msg.Tag = errorMsg
		return c.JSON(msg.Code, msg)
	}

	targetApiHandler := reflect.ValueOf(api).MethodByName(route.Operation.OperationID)
	if !targetApiHandler.IsValid() {
		msg := new(models.ApiMessage)
		msg.Code = http.StatusInternalServerError
		msg.Tag = "no handler found for operationId: " + route.Operation.OperationID
		return c.JSON(msg.Code, msg)
	}

	apiHandlerFunc, typeCheckOk := targetApiHandler.Interface().(func(echo.Context, map[string]string) error)
	if !typeCheckOk {
		log.Error(
			"wrong handler signature for operationId: " +
				route.Operation.OperationID + ", not type: func(echo.Context, map[string]string) error")
		msg := new(models.ApiMessage)
		msg.Tag = "wrong handler signature for operationId: " + route.Operation.OperationID
		msg.Code = http.StatusInternalServerError
		return c.JSON(msg.Code, msg)
	}

	return apiHandlerFunc(c, pathParams)
}

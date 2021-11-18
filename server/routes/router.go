package routes

import (
	"github.com/aetoledano/golang-service-template/common/ioutil"
	"github.com/aetoledano/golang-service-template/constants"
	controllers "github.com/aetoledano/golang-service-template/server/api"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
	"reflect"
	"strings"
)

type PathDefinition struct {
	Paths map[string]map[string]Operation `yaml:"paths"`
}

type Operation struct {
	Name string `yaml:"operationId"`
}

var api = &controllers.Api{}

func ConfigRoutes(e *echo.Echo) {

	e.File(constants.API_DOCS_PATH, constants.API_DOCS_FILE)
	e.GET(constants.HEALTHCHECK_PATH, healthCheck)

	var routes PathDefinition
	err := ioutil.ReadFileAsYml(constants.API_DOCS_FILE, &routes)
	if err != nil {
		log.Fatal(constants.SERVER_START_FAIL_ROUTES_MSG, err)
	}

	for path, pathMethods := range routes.Paths {
		for method, operation := range pathMethods {
			name := strings.Title(operation.Name)
			apiHandler := getHandler(name)
			register := getMethod(e, method)

			register(path, apiHandler)
		}
	}
}

func getMethod(e *echo.Echo, name string) func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route {
	value := reflect.ValueOf(e).MethodByName(strings.ToUpper(name))
	if !value.IsValid() {
		log.Fatal(
			constants.SERVER_START_FAIL_ROUTES_MSG,
			"no valid method found for "+name,
		)
	}

	methodFunction, typeCheckOk := value.Interface().(func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route)

	if !typeCheckOk {
		log.Fatal(
			constants.SERVER_START_FAIL_ROUTES_MSG,
			"type checking failed for "+value.String(),
		)
	}

	return methodFunction
}

func getHandler(name string) func(echo.Context) error {
	value := reflect.ValueOf(api).MethodByName(name)
	if !value.IsValid() {
		log.Fatal(
			constants.SERVER_START_FAIL_ROUTES_MSG,
			"no valid handler found for "+name,
		)
	}

	apiHandler, typeCheckOk := value.Interface().(func(echo.Context) error)
	if !typeCheckOk {
		log.Fatal(
			constants.SERVER_START_FAIL_ROUTES_MSG,
			"wrong handler signature for "+name+" "+value.String(),
		)
	}

	return apiHandler
}

func healthCheck(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "OK")
}

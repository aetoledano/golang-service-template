package middleware

import (
	"github.com/aetoledano/golang-service-template/common/ioutil"
	"github.com/aetoledano/golang-service-template/constants"
	"github.com/labstack/echo"
	echo_middleware "github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
)

func BuildCORSMiddleware() echo.MiddlewareFunc {
	allowedOrigins := make([]string, 0)
	err := ioutil.ReadFileAsYml(constants.CORS_FILE, &allowedOrigins)
	if err != nil {
		log.Fatal("could not load CORS origins file ", err)
	}

	return echo_middleware.CORSWithConfig(echo_middleware.CORSConfig{
		AllowOrigins: allowedOrigins,
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowOrigin,
		},
	})
}

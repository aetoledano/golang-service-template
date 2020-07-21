package middleware

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func HeadersLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Info("Request headers: ", c.Request().Header)
		return next(c)
	}
}


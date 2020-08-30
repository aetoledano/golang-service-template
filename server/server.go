package server

import (
	"context"
	"fmt"
	"github.com/aetoledano/golang-service-template/config"
	"github.com/aetoledano/golang-service-template/constants"
	"github.com/aetoledano/golang-service-template/server/middleware"
	"github.com/aetoledano/golang-service-template/services/db"
	"github.com/carlescere/scheduler"
	"github.com/labstack/echo"
	echoMiddleware "github.com/labstack/echo/middleware"
	echoLog "github.com/onrik/logrus/echo"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var EchoInstance *echo.Echo
var DbCheckJob *scheduler.Job

func Start() {

	EchoInstance = echo.New()

	//logging config
	EchoInstance.Logger = echoLog.NewLogger(log.StandardLogger(), constants.APP_NAME)
	EchoInstance.Use(echoMiddleware.Logger())
	EchoInstance.Use(echoMiddleware.Recover())
	//EchoInstance.Use(HeadersLoggerMiddleware)

	//cors config
	EchoInstance.Use(middleware.BuildCORSMiddleware())

	//routing config
	configRoutes(EchoInstance)

	serverPort := strconv.Itoa(config.App.Server.Port)
	go func() {
		if err := EchoInstance.Start(":" + serverPort); err != nil {
			errorName := fmt.Sprintf("%s", err)
			if errorName != "http: Server closed" {
				log.Fatal(constants.SERVER_START_FAIL_MSG, err)
			}
		}
	}()

	var err error
	DbCheckJob, err = scheduler.Every(10).Seconds().Run(DbConnectionCheck)
	if err != nil {
		log.Fatal(err)
	}

	WaitForInterrupts()
}

func DbConnectionCheck() {
	err := db.DbConnect()
	if err != nil {
		log.Error("database error: ", err)
	}
}

func WaitForInterrupts() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGKILL)
	s := <-c
	fmt.Printf("⇨ Gracefully stopping service, on signal: %s\n\n", s)

	shutdownServer()
}

func shutdownServer() {
	if err := EchoInstance.Shutdown(context.Background()); err != nil {
		log.Error(err)
	}

	DbCheckJob.Quit <- true
	if err := db.DbClose(); err != nil {
		log.Error(err)
	}
}

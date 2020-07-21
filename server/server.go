package server

import (
	"context"
	"fmt"
	"github.com/carlescere/scheduler"
	"github.com/labstack/echo"
	echo_middleware "github.com/labstack/echo/middleware"
	echolog "github.com/onrik/logrus/echo"
	log "github.com/sirupsen/logrus"
	"github.com/aetoledano/golang-service-template/config"
	"github.com/aetoledano/golang-service-template/server/middleware"
	"github.com/aetoledano/golang-service-template/services/db"
	"os"
	"os/signal"
	"syscall"
)

var EchoInstance *echo.Echo
var DbCheckJob *scheduler.Job

func Start() {

	EchoInstance = echo.New()

	//logging
	EchoInstance.Logger = echolog.NewLogger(log.StandardLogger(), "")
	EchoInstance.Use(echo_middleware.Logger())
	EchoInstance.Use(echo_middleware.Recover())
	//EchoInstance.Use(HeadersLoggerMiddleware)

	//cors config
	EchoInstance.Use(middleware.CraftCORSMiddleware())

	//routing
	configRouter(EchoInstance)

	go func() {
		if err := EchoInstance.Start(":" + config.SERVER_PORT); err != nil {
			errorName := fmt.Sprintf("%s", err)
			if errorName != "http: Server closed" {
				log.Fatal("⇨ Server could not be started: ", err)
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

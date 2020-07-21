package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/aetoledano/golang-service-template/config"
	"github.com/aetoledano/golang-service-template/server"
	"os"
)

func main() {

	//logger configuration
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)

	config.Init()

	server.Start()
}

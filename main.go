package main

import (
	"github.com/aetoledano/golang-service-template/config"
	"github.com/aetoledano/golang-service-template/server"
)

func main() {
	config.Load()

	server.Start()
}

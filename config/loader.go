package config

import (
	"github.com/aetoledano/golang-service-template/common/ioutil"
	"github.com/aetoledano/golang-service-template/constants"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

var App configuration

func Load() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)

	appEnvironment := os.Getenv(constants.APP_ENV_NAME)
	if appEnvironment == constants.EMPTY {
		appEnvironment = constants.DEFAULT_ENV
	}

	printProfile(constants.APP_NAME, appEnvironment)

	configFileName := strings.ReplaceAll(
		constants.CONFIG_FILE,
		constants.ENV_PLACEHOLDER,
		appEnvironment,
	)

	err := ioutil.ReadFileAsYml(configFileName, &App)
	if err != nil {
		log.Fatal(constants.SERVER_START_FAIL_MSG, err)
	}
}

func printProfile(name, env string) {
	print("\n⇨ Starting " + name + " with profile: " + env)
}

type configuration struct {
	Server struct {
		Port int
	}
	Api struct {
		Root string
	}
	Database struct {
		Dialect string
		Url     string
	}
}

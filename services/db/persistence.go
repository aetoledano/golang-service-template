package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"github.com/aetoledano/golang-service-template/config"
)

var DB *gorm.DB

func DbConnect() error {

	if DB != nil {
		err := DB.DB().Ping()
		if err != nil {
			return err
		}
		return nil
	}

	var err error
	DB, err = gorm.Open(config.DIALECT, config.CONNECTION_STRING())
	if err != nil {
		DB = nil
		return err
	}

	DB.LogMode(true)
	DB.SetLogger(log.StandardLogger())

	autoMigrateSchema()

	return nil
}

func autoMigrateSchema() {
	//schema migrations goes here
}

func DbClose() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

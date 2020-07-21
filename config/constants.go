package config

import (
	"os"
	"strings"
)

////
// CONSTANTS
////

////
// SERVICE WIDE
////
var APP_NAME = "golang-service-template"
var SERVER_PORT = "8080"
var API_ROOT = ""
var EMPTY = ""

////
// MAP KEYS
////
var API_ROOT_KEY = "API_ROOT"
var SERVER_PORT_KEY = "APP_PORT"

////
// TAGS
////
var OK = "OK"
var CREATED = "CREATED"
var BADREQUEST = "BAD_REQUEST"
var VALIDATIONERRORS = "VALIDATION_ERRORS"
var INTERNAL = "INTERNAL_SERVER_ERROR"
var NOTFOUND = "NOT_FOUND"
var NOCONTENT = "NO_CONTENT"
var FORBIDDEN = "FORBIDDEN"
var UNAUTHORIZED = "UNAUTHORIZED"

////
// DATABASE
////
var DB_CONN_ENV = "DATABASE_URL"
var DIALECT = "postgres"
var DBNAME = "defaultdb"
var DBHOST = "127.0.0.1"
var DBPORT = "5432"
var DBUSER = "postgres"
var DBPASS = "postgres"
var DBSSL = "sslmode=disable"

var CONNECTION_STRING = func() string {
	connString := strings.TrimSpace(os.Getenv(DB_CONN_ENV))
	connString = strings.ReplaceAll(connString, "'", EMPTY)
	connString = strings.ReplaceAll(connString, "\"", EMPTY)
	if connString != "" {
		return connString
	}
	return "host=" + DBHOST + " port=" + DBPORT + " user=" + DBUSER + " dbname=" + DBNAME + " password=" + DBPASS + " " + DBSSL
}

////
// SECURITY
////
var AUTH_HEADER = "Authorization"

////
// FILES
////
var CORS_FILE = "cors.policy.yml"
var CONFIG_FILE = "config.dev.yml"
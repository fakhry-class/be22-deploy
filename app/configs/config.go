package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	JWT_SECRET string
)

type AppConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOSTNAME string
	DB_PORT     int
	DB_NAME     string
}

//	func ReadEnv() *AppConfig {
//		var app = AppConfig{}
//		app.DB_USERNAME = os.Getenv("DBUSER")
//		app.DB_PASSWORD = os.Getenv("DBPASS")
//		app.DB_HOSTNAME = os.Getenv("DBHOST")
//		portConv, errConv := strconv.Atoi(os.Getenv("DBPORT"))
//		if errConv != nil {
//			panic("error convert dbport")
//		}
//		app.DB_PORT = portConv
//		app.DB_NAME = os.Getenv("DBNAME")
//		JWT_SECRET = os.Getenv("JWTSECRET")
//		return &app
//	}
func ReadEnv() (*AppConfig, bool) {
	var isExist = true
	var app = AppConfig{}
	if value, found := os.LookupEnv("DBUSER"); found {
		app.DB_USERNAME = value
	} else {
		isExist = false
	}
	if value, found := os.LookupEnv("DBPASS"); found {
		app.DB_PASSWORD = value
	} else {
		isExist = false
	}
	if value, found := os.LookupEnv("DBHOST"); found {
		app.DB_HOSTNAME = value
	} else {
		isExist = false
	}
	if value, found := os.LookupEnv("DBPORT"); found {
		portConv, errConv := strconv.Atoi(value)
		if errConv != nil {
			panic("error convert dbport")
		}
		app.DB_PORT = portConv
	} else {
		isExist = false
	}
	if value, found := os.LookupEnv("DBNAME"); found {
		app.DB_NAME = value
	} else {
		isExist = false
	}
	if value, found := os.LookupEnv("JWTSECRET"); found {
		JWT_SECRET = value
	} else {
		isExist = false
	}

	return &app, isExist
}

func InitConfig() *AppConfig {
	var appConv *AppConfig
	var isExist bool
	appConv, isExist = ReadEnv()
	if !isExist {
		errLoad := godotenv.Load(".env")
		if errLoad != nil {
			panic("error load env")
		}
		appConv, _ = ReadEnv()
	}
	return appConv
}

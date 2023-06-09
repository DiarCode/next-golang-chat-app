package config

var AppConfig *AppConfigType

type AppConfigType struct {
	APP_PORT string

	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
}

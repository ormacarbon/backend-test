package config

import "os"

var (
	DATABASE_URL   = ""
	SERVICE_EMAIL = ""
	PASSWORD_SERVICE_EMAIL = ""
)

func LoadEnv() {
	DATABASE_URL = os.Getenv("DATABASE_URL")
	SERVICE_EMAIL = os.Getenv("SERVICE_EMAIL")
	PASSWORD_SERVICE_EMAIL = os.Getenv("PASSWORD_SERVICE_EMAIL")
}

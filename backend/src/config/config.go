package config

import "os"

var (
	DATABASE_URL   = ""
	SERVICE_EMAIL = ""
)

func LoadEnv() {
	DATABASE_URL = os.Getenv("DATABASE_URL")
	SERVICE_EMAIL = os.Getenv("SERVICE_EMAIL")
}

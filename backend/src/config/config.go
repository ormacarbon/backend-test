package config

import "os"

var (
	DATABASE_URL = ""
)

func LoadEnv() {
	DATABASE_URL = os.Getenv("DATABASE_URL")
}

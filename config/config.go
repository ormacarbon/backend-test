package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if os.Getenv("RENDER") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Aviso: Nenhum arquivo .env encontrado. Usando variáveis de ambiente do sistema.")
		}
	}
}
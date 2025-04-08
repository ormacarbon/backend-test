package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Andreffelipe/carbon_offsets_api/config"
	_ "github.com/lib/pq"
)

func ConnectDB(config config.Config) *sql.DB {
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresUser,
		config.PostgresPass,
		config.PostgresDb,
	)
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
	return db
}

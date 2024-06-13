package main

import (
	"fmt"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/rostis232/golibrary/internal/pkg/app"
)

func main() {
	a, err := app.NewApp(pgConfig())
	if err != nil {
		log.Panic(err)
	}
	port := os.Getenv("PORT")
	log.Infof("PORT: %s", port)
	if err := a.Run(port); err != nil {
		log.Panic(err)
	}
}

func pgConfig() string {
	pgHost := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")
	pgUser := os.Getenv("PG_USER")
	pgPass := os.Getenv("PG_PASS")
	pgDBname := os.Getenv("PG_DB_NAME")

	log.Infof("PG_HOST: %s", pgHost)
	log.Infof("PG_PORT: %s", pgPort)
	log.Infof("PG_USER: %s", pgUser)
	log.Infof("PG_PASS: %s", pgPass)
	log.Infof("PG_DB_NAME: %s", pgDBname)
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5", pgHost, pgPort, pgUser, pgPass, pgDBname)
}

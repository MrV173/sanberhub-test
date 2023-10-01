package main

import (
	"sanberhub-test/database"
	postgre "sanberhub-test/pkg"
	"sanberhub-test/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	e := echo.New()

	postgre.DatabaseInit()
	database.RunMigration()
	routes.RouteInit(e.Group("/api/v1"))
	e.Logger.Fatal(e.Start(":5000"))
}

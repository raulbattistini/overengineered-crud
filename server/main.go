package main

import (
	"fmt"
	"log"
	"os"
	"server/configs"
	"server/context"
	_ "server/context"
	"server/db"
	"server/enums"
	"server/hepers"
	"server/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	os.Setenv("TZ", "Africa/Cairo")
	config, err := configs.LoadConfig()
	if err != nil {
		hepers.Log(err.Error(), &err, enums.Error)
	}

	_, err = db.Connect(&config.DbConfig)
	if err != nil {
		hepers.Log(err.Error(), &err, enums.Error)
	}

	context.InitContext()

	e := echo.New()

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	routes.SetupRoutes(e)

	log.Println("Hello, World!")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.ApiConfig.Port)))
}

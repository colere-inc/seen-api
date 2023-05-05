package main

import (
	"github.com/colere-inc/seen-api/app/adapter"
	"github.com/colere-inc/seen-api/app/common/config"
	"github.com/colere-inc/seen-api/app/infrastructure"
	"github.com/colere-inc/seen-api/app/infrastructure/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Init
	config.Init()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// DI
	firebase := infrastructure.NewFirebase()
	db := infrastructure.NewDB(firebase)
	freeeAccounting := infrastructure.NewFreeeAccounting()

	repository := model.NewPartnerRepository(db, freeeAccounting)
	controller := adapter.NewPartnerController(repository)

	adapter.NewRouter(e, *controller)

	// Start server
	e.Logger.Fatal(e.Start(":" + config.Port))
}

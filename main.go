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
	freeeInvoice := infrastructure.NewFreeeInvoice()

	partnerRepository := model.NewPartnerRepository(db, freeeAccounting)
	invoiceRepository := model.NewInvoiceRepository(freeeInvoice)
	partnerController := adapter.NewPartnerController(partnerRepository)
	invoiceController := adapter.NewInvoiceController(invoiceRepository)

	adapter.NewRouter(e, *partnerController, *invoiceController)

	// Start server
	e.Logger.Fatal(e.Start(":" + config.Port))
}

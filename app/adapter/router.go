package adapter

import (
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, controller PartnerController) {
	e.GET("/accounting/partners", controller.Get())
	e.GET("/accounting/partners/:partnerID", controller.GetByID())
	e.POST("/accounting/partners", hello)
}

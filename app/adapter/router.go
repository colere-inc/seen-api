package adapter

import (
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, controller PartnerController) {
	e.GET("/accounting/partners/:partnerID", controller.Get())
	e.POST("/accounting/partners", hello)
}

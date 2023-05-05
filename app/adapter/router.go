package adapter

import (
	"github.com/colere-inc/seen-api/app/infrastructure"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, fa *infrastructure.FreeeAccounting) {
	e.GET("/accounting/partners/:partnerID", get(fa))
	e.POST("/accounting/partners", hello)
}

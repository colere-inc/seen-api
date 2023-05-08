package adapter

import (
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, partnerController PartnerController, invoiceController InvoiceController) {
	e.GET("/accounting/partners", partnerController.Get())
	e.POST("/accounting/partners", partnerController.Add())

	e.POST("/invoice/invoices", invoiceController.AddInvoice())
}

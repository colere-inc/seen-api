package adapter

import (
	"github.com/colere-inc/seen-api/app/domain/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type InvoiceController struct {
	InvoiceRepository repository.InvoiceRepository
}

func NewInvoiceController(ir repository.InvoiceRepository) *InvoiceController {
	return &InvoiceController{InvoiceRepository: ir}
}

func (ic *InvoiceController) AddInvoice() echo.HandlerFunc {
	return func(c echo.Context) error {
		var spaceID string
		var paymentDate string
		echo.FormFieldBinder(c).
			MustString("space_id", &spaceID).
			String("payment_date", &paymentDate)
		invoice, err := ic.InvoiceRepository.Add(spaceID, paymentDate)
		if err != nil {
			panic(err)
		}
		return c.JSON(http.StatusCreated, invoice)
	}
}

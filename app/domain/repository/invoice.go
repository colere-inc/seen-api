package repository

import "github.com/colere-inc/seen-api/app/domain/model"

type InvoiceRepository interface {
	Add(spaceID string, paymentDate string) (*model.Invoice, error)
}

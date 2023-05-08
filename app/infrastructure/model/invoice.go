package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/colere-inc/seen-api/app/domain/model"
	"github.com/colere-inc/seen-api/app/domain/repository"
	"github.com/colere-inc/seen-api/app/infrastructure"
	"log"
	"net/http"
	"strconv"
	"time"
)

const invoicePath = "/invoices"

type InvoiceRepository struct {
	FreeeInvoice *infrastructure.FreeeInvoice
}

func NewInvoiceRepository(freeeInvoice *infrastructure.FreeeInvoice) repository.InvoiceRepository {
	return InvoiceRepository{FreeeInvoice: freeeInvoice}
}

func (ir InvoiceRepository) Add(spaceID string, paymentDate string) (*model.Invoice, error) {
	// request body
	if paymentDate == "" {
		paymentDate = getFormattedNDaysLater(7) // TODO: default で何日後に設定するか確認.
	}

	// TODO: spaceID から partnerID を求める
	partnerID := "61616449"
	partnerIntID, err := strconv.ParseInt(partnerID, 10, 64)
	if err != nil {
		panic(err)
	}

	body := invoicePostRequestBody{
		CompanyID:                 ir.FreeeInvoice.CompanyId,
		BillingDate:               getFormattedToday(),
		PaymentDate:               paymentDate,
		TaskEntryMethod:           model.TaxEntryMethodIn, // default を内税にしているが良いのか
		TaxFraction:               model.TaxFractionOmit,  // default が良いのか
		WithholdingTaxEntryMethod: model.TaxEntryMethodIn, // default を内税にしているが良いのか
		PartnerID:                 partnerIntID,
		PartnerTitle:              model.PartnerTitleGroup, // TODO
		Lines:                     []model.InvoiceLine{},
	}
	requestBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	// request
	res := ir.FreeeInvoice.Do(http.MethodPost, invoicePath, nil, bytes.NewBuffer(requestBody))
	if res.StatusCode != http.StatusCreated {
		log.Println("failed")
		panic(fmt.Sprintf("unexpected status: got %v, error: %s", res.StatusCode, string(res.ResBody)))
	}
	log.Println("success")

	// unmarshal
	var invoiceRes invoiceResponse
	err = json.Unmarshal(res.ResBody, &invoiceRes)
	if err != nil {
		panic(err)
	}

	// TODO
	// add to Firestore

	return &invoiceRes.Invoice, err
}

func getFormattedToday() string {
	now := time.Now()
	return now.Format("2006-01-02")
}

func getFormattedNDaysLater(nDays int) string {
	now := time.Now()
	return now.AddDate(0, 0, nDays).Format("2006-01-02")
}

type invoicePostRequestBody struct {
	CompanyID                 string              `json:"company_id"`
	BillingDate               string              `json:"billing_date"`
	PaymentDate               string              `json:"payment_date"`
	TaskEntryMethod           string              `json:"task_entry_method"`
	TaxFraction               string              `json:"tax_fraction"`
	WithholdingTaxEntryMethod string              `json:"withholding_tax_entry_method"`
	PartnerID                 int64               `json:"partner_id"`
	PartnerTitle              string              `json:"partner_title"`
	Lines                     []model.InvoiceLine `json:"lines"`
}

type invoiceResponse struct {
	Invoice model.Invoice `json:"invoice"`
}

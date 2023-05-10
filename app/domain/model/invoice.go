package model

import "encoding/json"

// Invoice https://developer.freee.co.jp/reference/iv/reference#operations-tag-Invoices
type Invoice struct {
	ID                           int64         `json:"id"`                              // 請求書ID
	CompanyID                    string        `json:"company_id"`                      // 事業所ID
	InvoiceNumber                string        `json:"invoice_number"`                  // 請求書番号
	BranchNo                     int           `json:"branch_no"`                       // 枝番
	Subject                      string        `json:"subject"`                         // 件名
	TemplateId                   int64         `json:"template_id"`                     // 帳票テンプレートID
	TemplateName                 string        `json:"template_name"`                   // 帳票テンプレート名
	BillingDate                  string        `json:"billing_date"`                    // 請求日
	PaymentDate                  string        `json:"payment_date"`                    // 入金期日 (optional)
	InvoiceNote                  string        `json:"invoice_note"`                    // 備考
	Memo                         string        `json:"memo"`                            // 社内メモ
	SendingStatus                string        `json:"sending_status"`                  // 送付ステータス
	PaymentStatus                string        `json:"payment_status"`                  // 入金ステータス
	CancelStatus                 string        `json:"cancel_status"`                   // 取消済みかのステータス
	DealStatus                   string        `json:"deal_status"`                     // 取引ステータス
	DealID                       int64         `json:"deal_id"`                         // 取引ID (optional)
	TaxEntryMethod               string        `json:"tax_entry_method"`                // 消費税の内税・外税区分 (optional)
	TaxFraction                  string        `json:"tax_fraction"`                    // 消費税端数の計算方法 (optional)
	WithholdingTaxEntryMethod    string        `json:"withholding_tax_entry_method"`    // 源泉徴収の計算方法 (optional)
	TotalAmount                  float64       `json:"total_amount"`                    // 合計金額
	CreatedAt                    string        `json:"created_at"`                      // 作成日時
	AmountWithholdingTax         json.Number   `json:"amount_withholding_tax"`          // 源泉徴収税 (optional)
	AmountIncludingTax           float64       `json:"amount_including_tax"`            // 税込金額
	AmountExcludingTax           json.Number   `json:"amount_excluding_tax"`            // 小計 (税別)
	AmountTax                    json.Number   `json:"amount_tax"`                      // 消費税額 (税率ごとの項目は一旦割愛する)
	PartnerID                    int64         `json:"partner_id"`                      // 取引先ID
	BillingPartnerID             int64         `json:"billing_partner_id"`              // 請求先ID (optional, 取引先IDより優先される)
	PartnerName                  string        `json:"partner_name"`                    // 取引先名 (optional)
	PartnerTitle                 string        `json:"partner_title"`                   // 敬称 (optional)
	PartnerAddressZipcode        string        `json:"partner_address_zipcode"`         // 郵便番号 (optional)
	PartnerAddressPrefectureCode int           `json:"partner_address_prefecture_code"` // 都道府県コード (optional)
	PartnerAddressStreetName1    string        `json:"partner_address_street_name_1"`   // 取引先 市区町村・番地 e.g. 東京都ＸＸ区ＹＹ１−１−１ (optional)
	PartnerAddressStreetName2    string        `json:"partner_address_street_name_2"`   // 取引先 建物名・部屋番号など (optional)
	PartnerContactDepartment     string        `json:"partner_contact_department"`      // 取引先部署 (optional)
	PartnerContactName           string        `json:"partner_contact_name"`            // 取引先担当者名 (optional)
	CompanyContactName           string        `json:"company_contact_name"`            // 自社担当者名 (optional)
	Template                     template      `json:"template"`                        // 帳票テンプレート情報 (optional)
	Lines                        []InvoiceLine `json:"lines"`                           // 請求書の明細行
}

const (
	SendingStatusSent   = "sent"   // 送付済み
	SendingStatusUnsent = "unsent" // 送付待ち
)

const (
	PaymentStatusSettled   = "settled"   // 入金済み
	PaymentStatusUnsettled = "unsettled" // 入金待ち
)

const (
	CancelStatusCanceled   = "canceled"   // 取消済み
	CancelStatusUncanceled = "uncanceled" // 取消済みではない
)

const (
	DealStatusRegistered   = "registered"   // 登録済み
	DealStatusUnregistered = "unregistered" // 登録待ち
)

const (
	TaxEntryMethodIn  = "in"  // 税込表示 (内税)
	TaxEntryMethodOut = "out" // 税別表示 (外税)
)

const (
	TaxFractionOmit    = "omit"     // 切り捨て
	TaxFractionRoundUp = "round_up" // 切り上げ
	TaxFractionRound   = "round"    // 四捨五入
)

const (
	PartnerTitleGroup  = "御中"
	PartnerTitlePerson = "様"
	PartnerTitleBlank  = "（空白）"
)

type template struct {
	Title                     string `json:"title"`                     // 請求書タイトル
	InvoiceRegistrationNumber string `json:"invoiceRegistrationNumber"` // インボイス制度適格請求書発行事業者登録番号
	CompanyName               string `json:"companyName"`               // 自社名
	CompanyDescription        string `json:"company_description"`       // 自社情報
	BankAccountToTransfer     string `json:"bank_account_to_transfer"`  // 振込先
	Message                   string `json:"message"`                   // メッセージ
}

type InvoiceLine struct {
	ID          int64  `json:"id"`          // 明細行ID
	LineType    string `json:"type"`        // 明細の種類
	Description string `json:"description"` // 嫡用(品名) e.g. 切手代
	// SalesDate      string `json:"sales_date"`       // 取引日 (optional)
	// Unit           string `json:"unit"`             // 明細の単位数 e.g. 個 (optional)
	Quantity  int64  `json:"quantity"`   // 明細の数量 (optional)
	UnitPrice string `json:"unit_price"` // 明細の単価 (optional)
	TaxRate   int    `json:"tax_rate"`   // 税率 Enum: [0, 8, 10] (optional)
	// ReducedTaxRate bool   `json:"reduced_tax_rate"` // 軽減税率対象 (default false, true は tax_rate = 8 のときのみ指定可能)
	Withholding bool `json:"withholding"` // 源泉徴収対象
}

const (
	LineTypeItem = "item" // 品目行
	LineTypeText = "text" // テキスト行
)

const (
	TaxRate0  = 0
	TaxRate8  = 8
	TaxRate10 = 10
)

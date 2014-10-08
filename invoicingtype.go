package paypal

import "time"

// https://developer.paypal.com/webapps/developer/docs/api/#common-invoicing-objects

var (
	InvoiceStatusDraft             InvoiceStatus = "DRAFT"
	InvoiceStatusSent              InvoiceStatus = "SENT"
	InvoiceStatusPaid              InvoiceStatus = "PAID"
	InvoiceStatusMarkedAsPaid      InvoiceStatus = "MARKED_AS_PAID"
	InvoiceStatusCancelled         InvoiceStatus = "CANCELLED"
	InvoiceStatusRefunded          InvoiceStatus = "REFUNDED"
	InvoiceStatusPartiallyRefunded InvoiceStatus = "PARTIALLY_REFUNDED"
	InvoiceStatusMarkedAsRefunded  InvoiceStatus = "MARKED_AS_REFUNDED"

	BillingInfoLanguageDADK BillingInfoLanguage = "da_DK"
	BillingInfoLanguageDEDE BillingInfoLanguage = "de_DE"
	BillingInfoLanguageENAU BillingInfoLanguage = "en_AU"
	BillingInfoLanguageENGB BillingInfoLanguage = "en_GB"
	BillingInfoLanguageENUS BillingInfoLanguage = "en_US"
	BillingInfoLanguageESES BillingInfoLanguage = "es_ES"
	BillingInfoLanguageESXC BillingInfoLanguage = "es_XC"
	BillingInfoLanguageFRCA BillingInfoLanguage = "fr_CA"
	BillingInfoLanguageFRFR BillingInfoLanguage = "fr_FR"
	BillingInfoLanguageFRXC BillingInfoLanguage = "fr_XC"
	BillingInfoLanguageHEIL BillingInfoLanguage = "he_IL"
	BillingInfoLanguageIDID BillingInfoLanguage = "id_ID"
	BillingInfoLanguageITIT BillingInfoLanguage = "it_IT"
	BillingInfoLanguageJAJP BillingInfoLanguage = "ja_JP"
	BillingInfoLanguageNLNL BillingInfoLanguage = "nl_NL"
	BillingInfoLanguageNONO BillingInfoLanguage = "no_NO"
	BillingInfoLanguagePLPL BillingInfoLanguage = "pl_PL"
	BillingInfoLanguagePTBR BillingInfoLanguage = "pt_BR"
	BillingInfoLanguagePTPT BillingInfoLanguage = "pt_PT"
	BillingInfoLanguageRURU BillingInfoLanguage = "ru_RU"
	BillingInfoLanguageSVSE BillingInfoLanguage = "sv_SE"
	BillingInfoLanguageTHTH BillingInfoLanguage = "th_TH"
	BillingInfoLanguageTRTR BillingInfoLanguage = "tr_TR"
	BillingInfoLanguageZHCN BillingInfoLanguage = "zh_CN"
	BillingInfoLanguageZHHK BillingInfoLanguage = "zh_HK"
	BillingInfoLanguageZHTW BillingInfoLanguage = "zh_TW"
	BillingInfoLanguageZHXC BillingInfoLanguage = "zh_XC"

	PaymentTermTypeDueOnReceipt PaymentTermType = "DUE_ON_RECEIPT"
	PaymentTermTypeNet10        PaymentTermType = "NET_10"
	PaymentTermTypeNet15        PaymentTermType = "NET_15"
	PaymentTermTypeNet30        PaymentTermType = "NET_30"
	PaymentTermTypeNet45        PaymentTermType = "NET_45"

	PaymentDetailTypePaypal   PaymentDetailType = "PAYPAL"
	PaymentDetailTypeExternal PaymentDetailType = "EXTERNAL"

	PaymentDetailTransactionTypeSale          PaymentDetailTransactionType = "SALE"
	PaymentDetailTransactionTypeAuthorization PaymentDetailTransactionType = "AUTHORIZATION"
	PaymentDetailTransactionTypeCapture       PaymentDetailTransactionType = "CAPTURE"

	PaymentDetailMethodBankTransfer PaymentDetailMethod = "BANK_TRANSFER"
	PaymentDetailMethodCash         PaymentDetailMethod = "CASH"
	PaymentDetailMethodCheck        PaymentDetailMethod = "CHECK"
	PaymentDetailMethodCreditCard   PaymentDetailMethod = "CREDIT_CARD"
	PaymentDetailMethodDebitCard    PaymentDetailMethod = "DEBIT_CARD"
	PaymentDetailMethodPaypal       PaymentDetailMethod = "PAYPAL"
	PaymentDetailMethodWireTransfer PaymentDetailMethod = "WIRE_TRANSFER"
	PaymentDetailMethodOther        PaymentDetailMethod = "OTHER"

	RefundDetailTypePaypal   RefundDetailType = "PAYPAL"
	RefundDetailTypeExternal RefundDetailType = "EXTERNAL"
)

type (
	InvoiceStatus                string
	BillingInfoLanguage          string
	PaymentTermType              string
	PaymentDetailType            string
	PaymentDetailTransactionType string
	PaymentDetailMethod          string
	RefundDetailType             string

	// Invoice maps to invoice object
	Invoice struct {
		ID                         string          `json:"id"`
		Number                     string          `json:"number,omitempty"`
		URI                        string          `json:"uri"`
		Status                     InvoiceStatus   `json:"status"`
		MerchantInfo               *MerchantInfo   `json:"merchant_info"`
		BillingInfo                []BillingInfo   `json:"billing_info"`
		ShippingInfo               *ShippingInfo   `json:"shipping_info"`
		Items                      []InvoiceItem   `json:"items"`
		InvoiceDate                *time.Time      `json:"invoice_date"`
		PaymentTerm                *PaymentTerm    `json:"payment_term,omitempty"`
		Discount                   *Cost           `json:"discount,omitempty"`
		ShippingCost               *ShippingCost   `json:"shipping_cost,omitempty"`
		Custom                     *CustomAmount   `json:"custom,omitempty"`
		TaxCalculatedAfterDiscount bool            `json:"tax_calculated_after_discount,omitempty"`
		TaxInclusive               bool            `json:"tax_inclusive"`
		Terms                      string          `json:"terms,omitempty"`
		Note                       string          `json:"note,omitempty"`
		MerchantMemo               string          `json:"merchant_memo,omitempty"`
		LogoURL                    string          `json:"logo_url,omitempty"`
		TotalAmount                *Currency       `json:"total_amount"`
		PaymentDetails             []PaymentDetail `json:"payment_details"`
		RefundDetails              []RefundDetail  `json:"refund_details"`
		Metadata                   *Metadata       `json:"metadata"`
	}

	// InvoiceItem maps to invoice_item object
	InvoiceItem struct {
		Name        string     `json:"name"`
		Description string     `json:"description,omitempty"`
		Quantity    float64    `json:"quantity"`
		UnitPrice   *Currency  `json:"unit_price"`
		Tax         *Tax       `json:"tax,omitempty"`
		Date        *time.Time `json:"date,omitempty"`
		Discount    *Cost      `json:"discount,omitempty"`
	}

	// MerchantInfo maps to merchant_info object
	MerchantInfo struct {
		Email          string   `json:"email"`
		FirstName      string   `json:"first_name,omitempty"`
		LastName       string   `json:"last_name,omitempty"`
		Address        *Address `json:"address,omitempty"`
		BusinessName   string   `json:"business_name,omitempty"`
		Phone          *Phone   `json:"phone,omitempty"`
		Fax            *Phone   `json:"fax,omitempty"`
		Website        string   `json:"website,omitempty"`
		TaxID          string   `json:"tax_id,omitempty"`
		AdditionalInfo string   `json:"additional_info,omitempty"`
	}

	// BillingInfo maps to billing_info object
	BillingInfo struct {
		Email          string              `json:"email"`
		FirstName      string              `json:"first_name,omitempty"`
		LastName       string              `json:"last_name,omitempty"`
		BusinessName   string              `json:"business_name,omitempty"`
		Address        *Address            `json:"address,omitempty"`
		Language       BillingInfoLanguage `json:"language,omitempty"`
		AdditionalInfo string              `json:"additional_info,omitempty"`
	}

	// ShippingInfo maps to shipping_info object
	ShippingInfo struct {
		FirstName    string   `json:"first_name,omitempty"`
		LastName     string   `json:"last_name,omitempty"`
		BusinessName string   `json:"business_name,omitempty"`
		Address      *Address `json:"address,omitempty"`
	}

	// PaymentTerm maps to payment_term object
	PaymentTerm struct {
		TermType PaymentTermType `json:"term_type"`
		DueDate  *time.Time      `json:"due_date"`
	}

	// Cost maps to cost object
	Cost struct {
		Percent int       `json:"percent"`
		Amount  *Currency `json:"amount"`
	}

	// ShippingCost maps to shipping_cost object
	ShippingCost struct {
		Amount *Currency `json:"amount"`
		Tax    *Tax      `json:"tax"`
	}

	// Tax maps to tax object
	Tax struct {
		ID      string    `json:"id"`
		Name    string    `json:"name"`
		Percent int       `json:"percent"`
		Amount  *Currency `json:"amount"`
	}

	// CustomAmount maps to custom_amount object
	CustomAmount struct {
		Label  string    `json:"label"`
		Amount *Currency `json:"amount"`
	}

	// PaymentDetail maps to payment_detail object
	PaymentDetail struct {
		Type            PaymentDetailType            `json:"type"`
		TransactionID   string                       `json:"transaction_id"`
		TransactionType PaymentDetailTransactionType `json:"transaction_type"`
		Date            *time.Time                   `json:"date"`
		Method          PaymentDetailMethod          `json:"method"`
		Note            string                       `json:"note,omitempty"`
	}

	// RefundDetail maps to refund_detail object
	RefundDetail struct {
		Type RefundDetailType `json:"type"`
		Date *time.Time       `json:"date"`
		Note string           `json:"note,omitempty"`
	}

	// Metadata maps to metadata object
	Metadata struct {
		CreatedDate     *time.Time `json:"created_date"`
		CreatedBy       string     `json:"created_by"`
		CancelledDate   *time.Time `json:"cancelled_date"`
		CancelledBy     string     `json:"cancelled_by"`
		LastUpdatedDate *time.Time `json:"last_updated_date"`
		LastUpdatedBy   string     `json:"last_updated_by"`
		FirstSentDate   *time.Time `json:"first_sent_date"`
		LastSentDate    *time.Time `json:"last_sent_date"`
		LastSentBy      *time.Time `json:"last_sent_by"`
	}

	// Search maps to search object. Invoice search parameters
	Search struct {
		Email                 string        `json:"email,omitempty"`
		RecipientFirstName    string        `json:"recipient_first_name,omitempty"`
		RecipientLastName     string        `json:"recipient_last_name,omitempty"`
		RecipientBusinessName string        `json:"recipient_business_name,omitempty"`
		Number                string        `json:"number,omitempty"`
		Status                InvoiceStatus `json:"status,omitempty"`
		LowerTotalAmount      *Currency     `json:"lower_total_amount,omitempty"`
		UpperTotalAmount      *Currency     `json:"upper_total_amount,omitempty"`
		StartInvoiceDate      *time.Time    `json:"start_invoice_date,omitempty"`
		EndInvoiceDate        *time.Time    `json:"end_invoice_date,omitempty"`
		StartDueDate          *time.Time    `json:"start_due_date,omitempty"`
		EndDueDate            *time.Time    `json:"end_due_date,omitempty"`
		StartPaymentDate      *time.Time    `json:"start_payment_date,omitempty"`
		EndPaymentDate        *time.Time    `json:"end_payment_date,omitempty"`
		StartCreationDate     *time.Time    `json:"start_creation_date,omitempty"`
		EndCreationDate       *time.Time    `json:"end_creation_date,omitempty"`
		Page                  int           `json:"page,omitempty"`
		PageSize              int           `json:"page_size,omitempty"`
		TotalCountRequired    bool          `json:"total_count_required,omitempty"`
	}

	// Notification maps to notification object. Email/SMS notification
	Notification struct {
		Subject        string `json:"subject,omitempty"`
		Note           string `json:"note,omitempty"`
		SendToMerchant bool   `json:"send_to_merchant"`
		SendToPayer    bool   `json:"send_to_payer"`
	}
)

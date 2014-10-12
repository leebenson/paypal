package paypal

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
	PaymentTermTypeNoDueDate    PaymentTermType = "NO_DUE_DATE"
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
		ID                         string          `json:"id,omitempty"`
		Number                     string          `json:"number,omitempty"`
		URI                        string          `json:"uri,omitempty"`
		Status                     InvoiceStatus   `json:"status,omitempty"`
		MerchantInfo               *MerchantInfo   `json:"merchant_info"`
		BillingInfo                []BillingInfo   `json:"billing_info"`
		ShippingInfo               *ShippingInfo   `json:"shipping_info,omitempty"`
		Items                      []InvoiceItem   `json:"items"`
		InvoiceDate                *Date           `json:"invoice_date,omitempty"`
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
		TotalAmount                *Currency       `json:"total_amount,omitempty"`
		PaymentDetails             []PaymentDetail `json:"payment_details,omitempty"`
		RefundDetails              []RefundDetail  `json:"refund_details,omitempty"`
		Metadata                   *Metadata       `json:"metadata,omitempty"`
	}

	// InvoiceItem maps to invoice_item object
	InvoiceItem struct {
		Name        string    `json:"name"`
		Description string    `json:"description,omitempty"`
		Quantity    float64   `json:"quantity"`
		UnitPrice   *Currency `json:"unit_price"`
		Tax         *Tax      `json:"tax,omitempty"`
		Date        *Date     `json:"date,omitempty"`
		Discount    *Cost     `json:"discount,omitempty"`
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
		DueDate  *Date           `json:"due_date,omitempty"`
	}

	// Cost maps to cost object
	Cost struct {
		Percent int       `json:"percent,omitempty"`
		Amount  *Currency `json:"amount,omitempty"`
	}

	// ShippingCost maps to shipping_cost object
	ShippingCost struct {
		Amount *Currency `json:"amount,omitempty"`
		Tax    *Tax      `json:"tax,omitempty"`
	}

	// Tax maps to tax object
	Tax struct {
		ID      string    `json:"id,omitempty"`
		Name    string    `json:"name"`
		Percent int       `json:"percent"`
		Amount  *Currency `json:"amount,omitempty"`
	}

	// CustomAmount maps to custom_amount object
	CustomAmount struct {
		Label  string    `json:"label,omitempty"`
		Amount *Currency `json:"amount,omitempty"`
	}

	// PaymentDetail maps to payment_detail object
	PaymentDetail struct {
		Type            PaymentDetailType            `json:"type,omitempty"`
		TransactionID   string                       `json:"transaction_id,omitempty"`
		TransactionType PaymentDetailTransactionType `json:"transaction_type,omitempty"`
		Date            *Date                        `json:"date,omitempty"`
		Method          PaymentDetailMethod          `json:"method"`
		Note            string                       `json:"note,omitempty"`
	}

	// RefundDetail maps to refund_detail object
	RefundDetail struct {
		Type RefundDetailType `json:"type"`
		Date *Date            `json:"date,omitempty"`
		Note string           `json:"note,omitempty"`
	}

	// Metadata maps to metadata object
	Metadata struct {
		CreatedDate     *Datetime `json:"created_date,omitempty"`
		CreatedBy       string    `json:"created_by,omitempty"`
		CancelledDate   *Datetime `json:"cancelled_date,omitempty"`
		CancelledBy     string    `json:"cancelled_by,omitempty"`
		LastUpdatedDate *Datetime `json:"last_updated_date,omitempty"`
		LastUpdatedBy   string    `json:"last_updated_by,omitempty"`
		FirstSentDate   *Datetime `json:"first_sent_date,omitempty"`
		LastSentDate    *Datetime `json:"last_sent_date,omitempty"`
		LastSentBy      *Datetime `json:"last_sent_by,omitempty"`
	}

	// Search maps to search object. Invoice search parameters
	InvoiceSearch struct {
		Email                 string        `json:"email,omitempty"`
		RecipientFirstName    string        `json:"recipient_first_name,omitempty"`
		RecipientLastName     string        `json:"recipient_last_name,omitempty"`
		RecipientBusinessName string        `json:"recipient_business_name,omitempty"`
		Number                string        `json:"number,omitempty"`
		Status                InvoiceStatus `json:"status,omitempty"`
		LowerTotalAmount      *Currency     `json:"lower_total_amount,omitempty"`
		UpperTotalAmount      *Currency     `json:"upper_total_amount,omitempty"`
		StartInvoiceDate      *Datetime     `json:"start_invoice_date,omitempty"`
		EndInvoiceDate        *Datetime     `json:"end_invoice_date,omitempty"`
		StartDueDate          *Datetime     `json:"start_due_date,omitempty"`
		EndDueDate            *Datetime     `json:"end_due_date,omitempty"`
		StartPaymentDate      *Datetime     `json:"start_payment_date,omitempty"`
		EndPaymentDate        *Datetime     `json:"end_payment_date,omitempty"`
		StartCreationDate     *Datetime     `json:"start_creation_date,omitempty"`
		EndCreationDate       *Datetime     `json:"end_creation_date,omitempty"`
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

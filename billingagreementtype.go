package paypal

import "time"

// https://developer.paypal.com/webapps/developer/docs/api/#common-billing-agreements-objects

var (
	PaymentCardTypeVisa          PaymentCardType = "VISA"
	PaymentCardTypeAmex          PaymentCardType = "AMEX"
	PaymentCardTypeSolo          PaymentCardType = "SOLO"
	PaymentCardTypeJCB           PaymentCardType = "JCB"
	PaymentCardTypeStar          PaymentCardType = "STAR"
	PaymentCardTypeDelta         PaymentCardType = "DELTA"
	PaymentCardTypeDiscover      PaymentCardType = "DISCOVER"
	PaymentCardTypeSwitch        PaymentCardType = "SWITCH"
	PaymentCardTypeMaestro       PaymentCardType = "MAESTRO"
	PaymentCardTypeCBNationale   PaymentCardType = "CB_NATIONALE"
	PaymentCardTypeConfinoga     PaymentCardType = "CONFINOGA"
	PaymentCardTypeCofidis       PaymentCardType = "COFIDIS"
	PaymentCardTypeElectron      PaymentCardType = "ELECTRON"
	PaymentCardTypeCetelem       PaymentCardType = "CETELEM"
	PaymentCardTypeChinaUnionPay PaymentCardType = "CHINA_UNION_PAY"
	PaymentCardTypeMasterCard    PaymentCardType = "MASTERCARD"

	PaymentCardStatusExpired PaymentCardStatus = "EXPIRED"
	PaymentCardStatusActive  PaymentCardStatus = "ACTIVE"

	CreditTypeBillMeLater            CreditType = "BILL_ME_LATER"
	CreditTypePaypalExtrasMasterCard CreditType = "PAYPAL_EXTRAS_MASTERCARD"
	CreditTypeEbayMasterCard         CreditType = "EBAY_MASTERCARD"
	CreditTypePaypalSmartConnect     CreditType = "PAYPAL_SMART_CONNECT"
)

type (
	PaymentCardType   string
	PaymentCardStatus string
	CreditType        string

	// Agreement maps to agreement object
	Agreement struct {
		ID                          string               `json:"id,omitempty"`
		Name                        string               `json:"name"`
		Description                 string               `json:"description"`
		StartDate                   *DatetimeRFC3339     `json:"start_date"`
		Payer                       *AgreementPayer      `json:"payer"`
		ShippingAddress             *Address             `json:"shipping_address,omitempty"`
		OverrideMerchantPreferences *MerchantPreferences `json:"override_merchant_preferences,omitempty"`
		OverrideChargeModels        []ChargeModels       `json:"override_charge_models,omitempty"`
		Plan                        *PatchPlan           `json:"plan"`
		CreateTime                  *time.Time           `json:"create_time,omitempty"`
		UpdateTime                  *time.Time           `json:"update_time,omitempty"`
		Links                       []Links              `json:"links,omitempty"`
	}

	// AgreementPayer maps to the payer object in Billing Agreements
	AgreementPayer struct {
		PaymentMethod      PaymentMethod                `json:"payment_method"`
		FundingInstruments []AgreementFundingInstrument `json:"funding_instruments,omitempty"`
		FundingOptionID    string                       `json:"funding_option_id,omitempty"`
		PayerInfo          *AgreementPayerInfo          `json:"payer_info,omitempty"`
	}

	// AgreementFundingInstrument maps to the funding_instrument object in Billing Agreements
	AgreementFundingInstrument struct {
		CreditCard       *AgreementCreditCard `json:"credit_card,omitempty"`
		CreditCardToken  *CreditCardToken     `json:"credit_card_token,omitempty"`
		PaymentCard      *PaymentCard         `json:"payment_card,omitempty"`
		PaymentCardToken *PaymentCardToken    `json:"payment_card_token,omitempty"`
		BankAccount      string               `json:"bank_account,omitempty"`
		BankAccountToken *BankToken           `json:"bank_token,omitempty"`
		Credit           *Credit              `json:"credit,omitempty"`
	}

	// AgreementCreditCard maps to the credit_card object in Billing Agreements
	AgreementCreditCard struct {
		ID             string          `json:"id,omitempty"`
		Number         string          `json:"number"`
		Type           CreditCardType  `json:"type"`
		ExpireMonth    string          `json:"expire_month"`
		ExpireYear     string          `json:"expire_year"`
		CVV2           string          `json:"cvv2,omitempty"`
		FirstName      string          `json:"first_name,omitempty"`
		LastName       string          `json:"last_name,omitempty"`
		BillingAddress *Address        `json:"billing_address,omitempty"`
		State          CreditCardState `json:"state,omitempty"`
		ValidUntil     string          `json:"valid_until,omitempty"`
		Links          []Links         `json:"links,omitempty"`
	}

	// PaymentCard maps to payment_card object
	PaymentCard struct {
		ID                 string            `json:"id,omitempty"`
		Number             string            `json:"number"`
		Type               PaymentCardType   `json:"type"`
		ExpireMonth        string            `json:"expire_month"`
		ExpireYear         string            `json:"expire_year"`
		StartMonth         string            `json:"start_month,omitempty"`
		StartYear          string            `json:"start_year,omitempty"`
		CVV2               string            `json:"cvv2,omitempty"`
		FirstName          string            `json:"first_name,omitempty"`
		LastName           string            `json:"last_name,omitempty"`
		BillingAddress     *Address          `json:"billing_address,omitempty"`
		ExternalCustomerID string            `json:"external_customer_id"`
		Status             PaymentCardStatus `json:"status,omitempty"`
		ValidUntil         string            `json:"valid_until,omitempty"`
		Links              []Links           `json:"links,omitempty"`
	}

	// PaymentCardToken maps to payment_card_token object
	// A resource representing a payment card that can be used to fund a payment.
	PaymentCardToken struct {
		PaymentCardID      string          `json:"payment_card_id"`
		ExternalCustomerID string          `json:"external_customer_id"`
		Last4              string          `json:"last4,omitempty"`
		Type               PaymentCardType `json:"type"`
		ExpireMonth        string          `json:"expire_month,omitempty"`
		ExpireYear         string          `json:"expire_year,omitempty"`
	}

	// BankToken maps to bank_token object
	// A resource representing a bank that can be used to fund a payment.
	BankToken struct {
		BankID                 string `json:"bank_id"`
		ExternalCustomerID     string `json:"external_customer_id"`
		MandateReferenceNumber string `json:"mandate_reference_number,omitempty"`
	}

	// Credit maps to credit object
	// A resource representing a credit instrument.
	Credit struct {
		ID    string     `json:"id,omitempty"`
		Type  CreditType `json:"type"`
		Terms string     `json:"terms,omitempty"`
	}

	// AgreementPayerInfo maps to payer_info object in billing agreement
	// A resource representing information about a Payer.
	AgreementPayerInfo struct {
		Email           string                    `json:"email,omitempty"`
		FirstName       string                    `json:"first_name,omitempty"`
		LastName        string                    `json:"last_name,omitempty"`
		PayerID         string                    `json:"payer_id,omitempty"`
		Phone           string                    `json:"phone,omitempty"`
		BillingAddress  *Address                  `json:"billing_address,omitempty"`
		ShippingAddress *AgreementShippingAddress `json:"shipping_address,omitempty"`
	}

	// AgreementShippingAddress maps to shipping_address object in billing agreement
	// Extended Address object used as shipping address in a payment.
	AgreementShippingAddress struct {
		ID             string `json:"id,omitempty"`
		RecipientName  string `json:"recipient_name"`
		DefaultAddress bool   `json:"default_address,omitempty"`
		Line1          string `json:"line1"`
		Line2          string `json:"line2,omitempty"`
		City           string `json:"city"`
		CountryCode    string `json:"country_code"`
		PostalCode     string `json:"postal_code,omitempty"`
		State          string `json:"state,omitempty"`
		Phone          string `json:"phone,omitempty"`
	}

	// OverrideChargeModel maps to overridec_charge_model object
	// A resource representing an override_charge_model to be used during creation
	// of the agreement.
	OverrideChargeModels struct {
		ChargeID string    `json:"charge_id"`
		Amount   *Currency `json:"amount"`
	}

	// AgreementStateDescriptor maps to agreement_state_descriptor object
	// Description of the current state of the agreement.
	AgreementStateDescriptor struct {
		Note   string    `json:"note,omitempty"`
		Amount *Currency `json:"amount"`
	}

	// AgreementTransactons maps to agreement_transactions object
	// A resource representing agreement_transactions that is returned during a
	// transaction search.
	AgreementTransactions struct {
		AgreementTransactionList []AgreementTransaction `json:"agreement_transaction_list"`
	}

	// AgreementTransaction maps to agreement_transaction object
	// A resource representing an agreement_transaction that is returned during
	// a transaction search.
	AgreementTransaction struct {
		TransactionID   string    `json:"transaction_id,omitempty"`
		Status          string    `json:"status,omitempty"`
		TransactionType string    `json:"transaction_type,omitempty"`
		Amount          *Currency `json:"amount"`
		FeeAmount       *Currency `json:"fee_amount"`
		NetAmount       *Currency `json:"net_amount"`
		PayerEmail      string    `json:"payer_email,omitempty"`
		PayerName       string    `json:"payer_name,omitempty"`
		TimeUpdated     string    `json:"time_updated,omitempty"`
		TimeZone        string    `json:"time_zone,omitempty"`
	}
)

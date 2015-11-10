package paypal

import (
	"time"
)

// https://developer.paypal.com/webapps/developer/docs/api/#common-payments-objects

var (
	AuthorizationStatePending           AuthorizationState = "pending"
	AuthorizationStateAuthorized        AuthorizationState = "authorized"
	AuthorizationStateCaptured          AuthorizationState = "captured"
	AuthorizationStatePartiallyCaptured AuthorizationState = "partially_captured"
	AuthorizationStateExpired           AuthorizationState = "expired"
	AuthorizationStateVoided            AuthorizationState = "voided"

	CaptureStatePending           CaptureState = "pending"
	CaptureStateCompleted         CaptureState = "completed"
	CaptureStateRefunded          CaptureState = "refunded"
	CaptureStatePartiallyRefunded CaptureState = "partially_refunded"

	CreditCardStateExpired CreditCardState = "expired"
	CreditCardStateOK      CreditCardState = "ok"

	OrderStatePending           OrderState = "PENDING"
	OrderStateCompleted         OrderState = "COMPLETED"
	OrderStateRefunded          OrderState = "REFUNDED"
	OrderStatePartiallyRefunded            = "PARTIALLY_REFUNDED"

	PendingReasonPayerShippingUnconfirmed PendingReason = "PAYER-SHIPPING-UNCONFIRMED"
	PendingReasonMultiCurrency            PendingReason = "MULTI-CURRENCY"
	PendingReasonRiskReview               PendingReason = "RISK-REVIEW"
	PendingReasonRegulatoryReview         PendingReason = "REGULATORY-REVIEW"
	PendingReasonVerificationRequired     PendingReason = "VERIFICATION-REQUIRED"
	PendingReasonOrder                    PendingReason = "ORDER"
	PendingReasonOther                    PendingReason = "OTHER"

	ReasonCodeChargeback                              ReasonCode = "CHARGEBACK"
	ReasonCodeGuarantee                               ReasonCode = "GUARANTEE"
	ReasonCodeBuyerComplaint                          ReasonCode = "BUYER_COMPLAINT"
	ReasonCodeRefund                                  ReasonCode = "REFUND"
	ReasonCodeUnconfirmedShippingAddress              ReasonCode = "UNCONFIRMED_SHIPPING_ADDRESS"
	ReasonCodeEcheck                                  ReasonCode = "ECHECK"
	ReasonCodeInternationalWithdrawal                 ReasonCode = "INTERNATIONAL_WITHDRAWAL"
	ReasonCodeReceivingPreferenceMandatesManualAction ReasonCode = "RECEIVING_PREFERENCE_MANDATES_MANUAL_ACTION"
	ReasonCodePaymentReview                           ReasonCode = "PAYMENT_REVIEW"
	ReasonCodeRegulatoryReview                        ReasonCode = "REGULATORY_REVIEW"
	ReasonCodeUnilateral                              ReasonCode = "UNILATERAL"
	ReasonCodeVerificationRequired                    ReasonCode = "VERIFICATION_REQUIRED"

	ProtectionEligibilityEligible          ProtectionEligibility = "ELIGIBLE"
	ProtectionEligibilityPartiallyEligible ProtectionEligibility = "PARTIALLY_ELIGIBLE"
	ProtectionEligibilityIneligible        ProtectionEligibility = "INELIGIBLE"

	ProtectionEligibilityTypeEligible                    ProtectionEligibilityType = "ELIGIBLE"
	ProtectionEligibilityTypeItemNotReceivedEligible     ProtectionEligibilityType = "ITEM_NOT_RECEIVED_ELIGIBLE"
	ProtectionEligibilityTypeIneligible                  ProtectionEligibilityType = "INELIGIBLE"
	ProtectionEligibilityTypeUnauthorizedPaymentEligible ProtectionEligibilityType = "UNAUTHORIZED_PAYMENT_ELIGIBLE"

	PaymentMethodCreditCard PaymentMethod = "credit_card"
	PaymentMethodPaypal     PaymentMethod = "paypal"

	PayerStatusVerified   PayerStatus = "VERIFIED"
	PayerStatusUnverified PayerStatus = "UNVERIFIED"

	PaymentStateCreated  PaymentState = "created"
	PaymentStateApproved PaymentState = "approved"
	PaymentStateFailed   PaymentState = "failed"
	PaymentStatePending  PaymentState = "pending"
	PaymentStateCanceled PaymentState = "canceled"
	PaymentStateExpired  PaymentState = "expired"

	AddressTypeResidential AddressType = "residential"
	AddressTypeBusiness    AddressType = "business"
	AddressTypeMailbox     AddressType = "mailbox"

	PaymentIntentSale      PaymentIntent = "sale"
	PaymentIntentAuthorize PaymentIntent = "authorize"
	PaymentIntentOrder     PaymentIntent = "order"

	RefundStatePending   RefundState = "pending"
	RefundStateCompleted RefundState = "completed"
	RefundStateFailed    RefundState = "failed"

	SaleStatePending           SaleState = "pending"
	SaleStateCompleted         SaleState = "completed"
	SaleStateRefunded          SaleState = "refunded"
	SaleStatePartiallyRefunded SaleState = "partially_refunded"

	SalePaymentModeInstantTransfer    SalePaymentMode = "INSTANT_TRANSFER"
	SalePaymentModeManualBankTransfer SalePaymentMode = "MANUAL_BANK_TRANSFER"
	SalePaymentModeDelayedTransfer    SalePaymentMode = "DELAYED_TRANSFER"
	SalePaymentModeEcheck             SalePaymentMode = "ECHECK"
)

type (
	AuthorizationState        string
	CaptureState              string
	CreditCardState           string
	OrderState                string
	PendingReason             string
	ReasonCode                string
	ProtectionEligibility     string
	ProtectionEligibilityType string
	TaxIDType                 string
	PaymentState              string
	AddressType               string
	PaymentMethod             string
	PayerStatus               string
	PaymentIntent             string
	RefundState               string
	SaleState                 string
	SalePaymentMode           string

	// Address maps to address object
	Address struct {
		Line1       string `json:"line1"`
		Line2       string `json:"line2,omitempty"`
		City        string `json:"city"`
		CountryCode string `json:"country_code"`
		PostalCode  string `json:"postal_code,omitempty"`
		State       string `json:"state,omitempty"`
		Phone       string `json:"phone,omitempty"`
	}

	// Amount maps to the amount object
	Amount struct {
		Currency string   `json:"currency"`
		Total    string   `json:"total"`
		Details  *Details `json:"details,omitempty"`
	}

	// Authorization maps to the  authorization object
	Authorization struct {
		Id                        int64              `json:"-"`
		Amount                    *Amount            `json:"amount,omitempty"`
		CreateTime                *time.Time         `json:"create_time,omitempty"`
		UpdateTime                *time.Time         `json:"update_time,omitempty"`
		State                     AuthorizationState `json:"state,omitempty"`
		ParentPayment             string             `json:"parent_payment,omitempty"`
		ID                        string             `json:"id,omitempty"`
		ValidUntil                *time.Time         `json:"valid_until,omitempty"`
		Links                     []Links            `json:"links,omitempty"`
		ClearingTime              string             `json:"clearing_time,omitempty"`
		ProtectionEligibility     string             `json:"protection_eligibility,omitempty"`
		ProtectionEligibilityType string             `json:"protection_eligibility_type,omitempty"`
	}

	// Capture maps to the capture object
	Capture struct {
		Id             int64        `json:"-"`
		Amount         *Amount      `json:"amount,omitempty"`
		IsFinalCapture bool         `json:"is_final_capture"`
		CreateTime     *time.Time   `json:"create_time,omitempty"`
		UpdateTime     *time.Time   `json:"update_time,omitempty"`
		State          CaptureState `json:"state,omitempty"`
		ParentPayment  string       `json:"parent_payment,omitempty"`
		ID             string       `json:"id,omitempty"`
		Links          []Links      `json:"links,omitempty"`
	}

	// Details maps to the details object
	Details struct {
		Id               int64  `json:"-"`
		Shipping         string `json:"shipping,omitempty"`
		Subtotal         string `json:"subtotal"`
		Tax              string `json:"tax,omitempty"`
		Fee              string `json:"fee,omitempty"`
		HandlingFee      string `json:"handling_fee,omitempty"`
		Insurance        string `json:"insurance,omitempty"`
		ShippingDiscount string `json:"shipping_discount,omitempty"`
	}

	// PaymentError maps to the error object for payments
	PaymentError struct {
		Name            string               `json:"name,omitempty"`
		DebugID         string               `json:"debug_id,omitempty"`
		Message         string               `json:"message,omitempty"`
		InformationLink string               `json:"information_link,omitempty"`
		Details         *PaymentErrorDetails `json:"details,omitempty"`
	}

	// PaymentErrorDetails maps to the error_details object for payments
	PaymentErrorDetails struct {
		Field string `json:"field,omitempty"`
		Issue string `json:"issue,omitempty"`
	}

	// CreditCard maps to credit_card object
	CreditCard struct {
		ID             string          `json:"id,omitempty"`
		PayerID        string          `json:"payer_id,omitempty"`
		Number         string          `json:"number"`
		Type           string          `json:"type"`
		ExpireMonth    string          `json:"expire_month"`
		ExpireYear     string          `json:"expire_year"`
		CVV2           string          `json:"cvv2,omitempty"`
		FirstName      string          `json:"first_name,omitempty"`
		LastName       string          `json:"last_name,omitempty"`
		BillingAddress *Address        `json:"billing_address,omitempty"`
		State          CreditCardState `json:"state,omitempty"`
		ValidUntil     string          `json:"valid_until,omitempty"`
	}

	// CreditCardToken maps to credit_card_token object
	CreditCardToken struct {
		CreditCardID string `json:"credit_card_id"`
		PayerID      string `json:"payer_id,omitempty"`
		Last4        string `json:"last4,omitempty"`
		ExpireYear   string `json:"expire_year,omitempty"`
		ExpireMonth  string `json:"expire_month,omitempty"`
	}

	// FundingInstrument maps to funding_instrument object
	FundingInstrument struct {
		CreditCard      *CreditCard      `json:"credit_card,omitempty"`
		CreditCardToken *CreditCardToken `json:"credit_card_token,omitempty"`
	}

	// Item maps to item object
	Item struct {
		Quantity    int    `json:"quantity"`
		Name        string `json:"name"`
		Price       string `json:"price"`
		Currency    string `json:"currency"`
		SKU         string `json:"sku,omitempty"`
		Description string `json:"description,omitempty"`
		Tax         string `json:"tax,omitempty"`
	}

	// ItemList maps to item_list object
	ItemList struct {
		Items           []Item           `json:"items,omitempty"`
		ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
	}

	// Order maps to order object
	Order struct {
		ID                        string        `json:"id,omitempty"`
		PurchaseUnitReferenceID   string        `json:"purchase_unit_reference_id,omitempty"`
		CreateTime                *time.Time    `json:"create_time,omitempty"`
		UpdateTime                *time.Time    `json:"update_time,omitempty"`
		Amount                    []Amount      `json:"amount,omitempty"`
		State                     OrderState    `json:"state,omitempty"`
		PendingReason             PendingReason `json:"pending_reason,omitempty"`
		ReasonCode                ReasonCode    `json:"reason_code,omitempty"`
		ClearingTime              string        `json:"clearing_time,omitempty"`
		ProtectionEligibility     string        `json:"protection_eligibility,omitempty"`
		ProtectionEligibilityType string        `json:"protection_eligiblity_type,omitempty"`
	}

	// Payer maps to payer object
	Payer struct {
		PaymentMethod      PaymentMethod       `json:"payment_method"`
		FundingInstruments []FundingInstrument `json:"funding_instruments,omitempty"`
		PayerInfo          *PayerInfo          `json:"payer_info,omitempty"`
		Status             PayerStatus         `json:"payer_status,omitempty"`
	}

	// PayerInfo maps to payer_info object
	PayerInfo struct {
		Email           string           `json:"email,omitempty"`
		FirstName       string           `json:"first_name,omitempty"`
		LastName        string           `json:"last_name,omitempty"`
		PayerID         string           `json:"payer_id,omitempty"`
		Phone           string           `json:"phone,omitempty"`
		ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
		TaxIDType       TaxIDType        `json:"tax_id_type,omitempty"`
		TaxID           string           `json:"tax_id,omitempty"`
	}

	// Payment maps to payment object
	Payment struct {
		Intent              PaymentIntent `json:"intent"`
		Payer               *Payer        `json:"payer"`
		Transactions        []Transaction `json:"transactions"`
		RedirectURLs        *RedirectURLs `json:"redirect_urls,omitempty"`
		ID                  string        `json:"id,omitempty"`
		CreateTime          *time.Time    `json:"create_time,omitempty"`
		State               PaymentState  `json:"state,omitempty"`
		UpdateTime          *time.Time    `json:"update_time,omitempty"`
		ExperienceProfileID string        `json:"experience_profile_id,omitempty"`
	}

	// PaymentExecution maps to payment_execution object
	PaymentExecution struct {
		PayerID      string        `json:"payer_id,omitempty"`
		Transactions []Transaction `json:"transactions,omitempty"`
	}

	// RedirectURLs maps to redirect_urls object
	RedirectURLs struct {
		ReturnURL string `json:"return_url,omitempty"`
		CancelURL string `json:"cancel_url,omitempty"`
	}

	// Refund maps to refund object
	Refund struct {
		ID            string      `json:"id,omitempty"`
		Amount        *Amount     `json:"amount,omitempty"`
		CreateTime    *time.Time  `json:"create_time,omitempty"`
		State         RefundState `json:"state,omitempty"`
		CaptureID     string      `json:"capture_id,omitempty"`
		ParentPayment string      `json:"parent_payment,omitempty"`
		UpdateTime    *time.Time  `json:"update_time,omitempty"`
	}

	// Resource can be either sale, authorization, capture or refund object
	Resource struct {
		Sale          *Sale          `json:"sale,omitempty"`
		Authorization *Authorization `json:"authorization,omitempty"`
		Capture       *Capture       `json:"capture,omitempty"`
		Refund        *Refund        `json:"refund,omitempty"`
	}

	// Sale maps to sale object
	Sale struct {
		ID                        string                    `json:"id,omitempty"`
		Amount                    *Amount                   `json:"amount,omitempty"`
		Description               string                    `json:"description,omitempty"`
		CreateTime                *time.Time                `json:"create_time,omitempty"`
		State                     SaleState                 `json:"state,omitempty"`
		ParentPayment             string                    `json:"parent_payment,omitempty"`
		UpdateTime                *time.Time                `json:"update_time,omitempty"`
		PaymentMode               SalePaymentMode           `json:"payment_mode,omitempty"`
		PendingReason             PendingReason             `json:"pending_reason,omitempty"`
		ReasonCode                ReasonCode                `json:"reason_code,omitempty"`
		ClearingTime              string                    `json:"clearing_time,omitempty"`
		ProtectionEligibility     ProtectionEligibility     `json:"protection_eligibility,omitempty"`
		ProtectionEligibilityType ProtectionEligibilityType `json:"protection_eligibility_type,omitempty"`
		Links                     []Links                   `json:"links,omitempty"`
	}

	// ShippingAddress maps to shipping_address object
	ShippingAddress struct {
		RecipientName string      `json:"recipient_name,omitempty"`
		Type          AddressType `json:"type,omitempty"`
		Line1         string      `json:"line1"`
		Line2         string      `json:"line2,omitempty"`
		City          string      `json:"city"`
		CountryCode   string      `json:"country_code"`
		PostalCode    string      `json:"postal_code,omitempty"`
		State         string      `json:"state,omitempty"`
		Phone         string      `json:"phone,omitempty"`
	}

	// Transaction maps to transaction object
	Transaction struct {
		Amount           *Amount    `json:"amount"`
		Description      string     `json:"description,omitempty"`
		ItemList         *ItemList  `json:"item_list,omitempty"`
		RelatedResources []Resource `json:"related_resources,omitempty"`
		InvoiceNumber    string     `json:"invoice_number,omitempty"`
		Custom           string     `json:"custom,omitempty"`
		SoftDescriptor   string     `json:"soft_descriptor,omitempty"`
	}
)

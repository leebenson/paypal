package paypal

import "time"

// https://developer.paypal.com/webapps/developer/docs/api/#plan-object

var (
	PlanTypeFixed    PlanType = "FIXED"
	PlanTypeInfinite PlanType = "INFINITE"

	PlanStateCreated  PlanState = "CREATED"
	PlanStateActive   PlanState = "ACTIVE"
	PlanStateInactive PlanState = "INACTIVE"

	PaymentDefinitionTypeTrial   PaymentDefinitionType = "TRIAL"
	PaymentDefinitionTypeRegular PaymentDefinitionType = "REGULAR"

	ChargeModelsTypeShipping ChargeModelsType = "shipping"
	ChargeModelsTypeTax      ChargeModelsType = "tax"

	PatchOperationAdd     PatchOperation = "add"
	PatchOperationRemove  PatchOperation = "remove"
	PatchOperationReplace PatchOperation = "replace"
	PatchOperationMove    PatchOperation = "move"
	PatchOperationCopy    PatchOperation = "copy"
	PatchOperationTest    PatchOperation = "test"
)

type (
	PlanType              string
	PlanState             string
	PaymentDefinitionType string
	TermType              string
	ChargeModelsType      string
	PatchOperation        string

	// Plan maps to plan object
	Plan struct {
		ID                  string               `json:"id"`
		Name                string               `json:"name"`
		Description         string               `json:"description"`
		Type                PlanType             `json:"type"`
		State               PlanState            `json:"state"`
		Payee               *Payee               `json:"payee"`
		CreateTime          *time.Time           `json:"create_time"`
		UpdateTime          *time.Time           `json:"update_time"`
		PaymentDefinitions  []PaymentDefinition  `json:"payment_definitions"`
		Terms               []Terms              `json:"terms"`
		MerchantPreferences *MerchantPreferences `json:"merchant_preferences,omitempty"`
		Links               []Links              `json:"links"`
	}

	// Payee maps to payee object
	Payee struct {
		Email                string `json:"email"`
		MerchantID           string `json:"merchant_id"`
		Phone                *Phone `json:"phone,omitempty"`
		AdditionalProperties string `json:"additional_properties,omitempty"`
	}

	// Phone maps to phone object
	Phone struct {
		CountryCode    string `json:"country_code"`
		NationalNumber string `json:"national_number"`
		Extension      string `json:"extension,omitempty"`
	}

	// PaymentDefinition maps to payment_definition object
	PaymentDefinition struct {
		ID                string                `json:"id"`
		Name              string                `json:"name"`
		Type              PaymentDefinitionType `json:"type"`
		FrequencyInterval string                `json:"frequency_interval"`
		Frequency         string                `json:"frequency"`
		Cycles            string                `json:"cycles"`
		Amount            *Currency             `json:"amount"`
		ChargeModels      []ChargeModels        `json:"charge_models,omitempty"`
	}

	// ChargeModels maps to charge_models object
	ChargeModels struct {
		ID     string           `json:"id"`
		Type   ChargeModelsType `json:"type"`
		Amount *Currency        `json:"amount"`
	}

	// Terms maps to terms object
	Terms struct {
		ID               string    `json:"id"`
		Type             TermType  `json:"type"`
		MaxBillingAmount *Currency `json:"max_billing_amount"`
		Occurrences      string    `json:"occurrences"`
		AmountRange      *Currency `json:"amount_range"`
		BuyerEditable    string    `json:"buyer_editable"`
	}

	// MerchantPreferences maps to merchant_preferences boject
	MerchantPreferences struct {
		ID                      string    `json:"id"`
		SetupFee                *Currency `json:"setup_fee,omitempty"`
		CancelURL               string    `json:"cancel_url"`
		ReturnURL               string    `json:"return_url"`
		NotifyURL               string    `json:"notify_url"`
		MaxFailAttemps          string    `json:"max_fail_attemps,omitempty"` // Default is 0, which is unlimited
		AutoBillAmount          string    `json:"auto_bill_amount,omitempty,omitempty"`
		InitialFailAmountAction string    `json:"initial_fail_amount_action,omitempty"`
		AcceptedPaymentType     string    `json:"accepted_payment_type"`
		CharSet                 string    `json:"char_set"`
	}

	// PatchRequest maps to patch_request object
	PatchRequest struct {
		OP    PatchOperation `json:"op"`
		Path  string         `json:"path"`
		Value string         `json:"value"`
		From  string         `json:"from"`
	}

	// PlanList maps to plan_list object
	PlanList struct {
		Plans []Plan  `json:"plans"`
		Links []Links `json:"links"`
	}
)

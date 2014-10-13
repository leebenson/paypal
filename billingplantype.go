package paypal

import "time"

// https://developer.paypal.com/webapps/developer/docs/api/#plan-object

var (
	PlanTypeFixed    PlanType = "fixed"
	PlanTypeInfinite PlanType = "infinite"

	PlanStateCreated  PlanState = "CREATED"
	PlanStateActive   PlanState = "ACTIVE"
	PlanStateInactive PlanState = "INACTIVE"

	PaymentDefinitionTypeTrial   PaymentDefinitionType = "TRIAL"
	PaymentDefinitionTypeRegular PaymentDefinitionType = "REGULAR"

	ChargeModelsTypeShipping ChargeModelsType = "SHIPPING"
	ChargeModelsTypeTax      ChargeModelsType = "TAX"

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
		ID                  string               `json:"id,omitempty"`
		Name                string               `json:"name"`
		Description         string               `json:"description"`
		Type                PlanType             `json:"type"`
		State               PlanState            `json:"state,omitempty"`
		Payee               *Payee               `json:"payee,omitempty"`
		CreateTime          *time.Time           `json:"create_time,omitempty"`
		UpdateTime          *time.Time           `json:"update_time,omitempty"`
		PaymentDefinitions  []PaymentDefinition  `json:"payment_definitions,omitempty"`
		Terms               []Terms              `json:"terms,omitempty"`
		MerchantPreferences *MerchantPreferences `json:"merchant_preferences,omitempty"`
		Links               []Links              `json:"links,omitempty"`
	}

	// PatchPlan is used in PATCH reqeusts
	PatchPlan struct {
		ID                  string               `json:"id,omitempty"`
		Name                string               `json:"name,omitempty"`
		Description         string               `json:"description,omitempty"`
		Type                PlanType             `json:"type,omitempty"`
		State               PlanState            `json:"state,omitempty"`
		Payee               *Payee               `json:"payee,omitempty"`
		CreateTime          *time.Time           `json:"create_time,omitempty"`
		UpdateTime          *time.Time           `json:"update_time,omitempty"`
		PaymentDefinitions  []PaymentDefinition  `json:"payment_definitions,omitempty"`
		Terms               []Terms              `json:"terms,omitempty"`
		MerchantPreferences *MerchantPreferences `json:"merchant_preferences,omitempty"`
		Links               []Links              `json:"links,omitempty"`
	}

	// Payee maps to payee object
	Payee struct {
		Email                string `json:"email"`
		MerchantID           string `json:"merchant_id"`
		Phone                *Phone `json:"phone,omitempty"`
		AdditionalProperties string `json:"additional_properties,omitempty"`
	}

	// Phone maps to phone object
	// See commontype.go

	// PaymentDefinition maps to payment_definition object
	PaymentDefinition struct {
		ID                string                `json:"id,omitempty"`
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
		ID     string           `json:"id,omitempty"`
		Type   ChargeModelsType `json:"type"`
		Amount *Currency        `json:"amount"`
	}

	// Terms maps to terms object
	Terms struct {
		ID               string    `json:"id,omitempty"`
		Type             TermType  `json:"type"`
		MaxBillingAmount *Currency `json:"max_billing_amount"`
		Occurrences      string    `json:"occurrences"`
		AmountRange      *Currency `json:"amount_range"`
		BuyerEditable    string    `json:"buyer_editable"`
	}

	// MerchantPreferences maps to merchant_preferences boject
	MerchantPreferences struct {
		ID                      string    `json:"id,omitempty"`
		SetupFee                *Currency `json:"setup_fee,omitempty"`
		CancelURL               string    `json:"cancel_url"`
		ReturnURL               string    `json:"return_url"`
		NotifyURL               string    `json:"notify_url,omitempty"`
		MaxFailAttempts         string    `json:"max_fail_attempts,omitempty"` // Default is 0, which is unlimited
		AutoBillAmount          string    `json:"auto_bill_amount,omitempty,omitempty"`
		InitialFailAmountAction string    `json:"initial_fail_amount_action,omitempty"`
		AcceptedPaymentType     string    `json:"accepted_payment_type,omitempty"`
		CharSet                 string    `json:"char_set,omitempty"`
	}

	// PlanList maps to plan_list object
	PlanList struct {
		Plans []Plan  `json:"plans"`
		Links []Links `json:"links"`
	}
)

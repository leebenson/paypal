package paypal

var (
	AddressTypeResidential AddressType = "residential"
	AddressTypeBusiness    AddressType = "business"
	AddressTypeMailbox     AddressType = "mailbox"

	CreditCardTypeVisa       CreditCardType = "visa"
	CreditCardTypeMastercard CreditCardType = "mastercard"
	CreditCardTypeDiscover   CreditCardType = "discover"
	CreditCardTypeAmex       CreditCardType = "amex"

	CreditCardStateExpired CreditCardState = "expired"
	CreditCardStateOK      CreditCardState = "ok"
)

type (
	AddressType     string
	CreditCardType  string
	CreditCardState string

	// Links maps to links object
	Links struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
		// TODO: Support HyperSchema with its multiple types per field
		// TargetSchema HyperSchema `json:"targetSchema"`
		Method  string `json:"method"`
		Enctype string `json:"enctype"`
		// Schema HyperSchema `json:"schema"`
	}

	// Currency maps to currency object
	// Base object for all financial value related fields (balance, payment due, etc.)
	Currency struct {
		Currency string `json:"currency"`
		Value    string `json:"value"`
	}

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

	// CreditCard maps to credit_card object
	CreditCard struct {
		ID             string          `json:"id,omitempty"`
		PayerID        string          `json:"payer_id,omitempty"`
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
	}

	// CreditCardToken maps to credit_card_token object
	CreditCardToken struct {
		CreditCardID string `json:"credit_card_id"`
		PayerID      string `json:"payer_id,omitempty"`
		Last4        string `json:"last4,omitempty"`
		ExpireYear   string `json:"expire_year,omitempty"`
		ExpireMonth  string `json:"expire_month,omitempty"`
	}
)

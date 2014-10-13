package paypal

import (
	"fmt"
	"net/http"
)

// https://developer.paypal.com/webapps/developer/docs/api/#vault

// StoreCreditCard stores credit card details with Paypal. To use the stored card,
// use the returned ID as CreditCardID within a CreditCardToken. Including a PayerID
// is also recommended
func (c *Client) StoreCreditCard(creditCard *CreditCard) (*CreditCard, error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/vault/credit-card", c.APIBase), creditCard)
	if err != nil {
		return nil, err, nil
	}

	v := &CreditCard{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// DeleteStoredCreditCard delete details of a credit card. Note that even though
// a credit card is deleted, some limited information about that credit card is
// still provided when a sale is retrieved.
func (c *Client) DeleteStoredCreditCard(creditCardID string) (error, *http.Response) {
	req, err := NewRequest("DELETE", fmt.Sprintf("%s/vault/credit-card/%s", c.APIBase, creditCardID), nil)
	if err != nil {
		return err, nil
	}

	resp, err := c.SendWithAuth(req, nil)

	return err, resp
}

// GetStoredCreditCard returns details of a stored credit card.
func (c *Client) GetStoredCreditCard(creditCardID string) (*CreditCard, error, *http.Response) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/vault/credit-card/%s", c.APIBase, creditCardID), nil)
	if err != nil {
		return nil, err, nil
	}

	v := &CreditCard{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// UpdateStoredCreditCard modifies a stored credit card
func (c *Client) UpdateStoredCreditCard(creditCardID string, creditCard *PatchCreditCard) (*CreditCard, error, *http.Response) {
	req, err := NewRequest("PATCH", fmt.Sprintf("%s/vault/credit-card/%s", c.APIBase, creditCardID), struct {
		Path  string           `json:"path"`
		Value *PatchCreditCard `json:"value"`
		OP    PatchOperation   `json:"op"`
	}{
		Path:  "/",
		Value: creditCard,
		OP:    PatchOperationReplace,
	})
	if err != nil {
		return nil, err, nil
	}

	v := &CreditCard{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

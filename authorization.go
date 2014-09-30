package paypal

import "fmt"

// https://developer.paypal.com/webapps/developer/docs/api/#authorizations

type (
	AuthWithAmountReq struct {
		Amount *Amount `json:"amount"`
	}
)

// GetAuthorization returns an authorization by ID
func (c *Client) GetAuthorization(authID string) (*Authorization, error) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/payments/authorization/%s", c.APIBase, authID), nil)
	if err != nil {
		return nil, err
	}

	v := &Authorization{}

	err = c.SendWithAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// CaptureAuthorization captures and process an existing authorization.
// To use this method, the original payment must have Intent set to PaymentIntentAuthorize
func (c *Client) CaptureAuthorization(authID string, a *Amount, isFinalCapture bool) (*Capture, error) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/authorization/%s/capture", c.APIBase, authID), struct {
		Amount         *Amount `json:"amount"`
		IsFinalCapture bool    `json:"is_final_capture"`
	}{a, isFinalCapture})
	if err != nil {
		return nil, err
	}

	v := &Capture{}

	err = c.SendWithAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// VoidAuthorization voids a previously authorized payment. A fully
// captured authorization cannot be voided
func (c *Client) VoidAuthorization(authID string) (*Authorization, error) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/authorization/%s/void", c.APIBase, authID), nil)
	if err != nil {
		return nil, err
	}

	v := &Authorization{}

	err = c.SendWithAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// ReauthorizeAuthorization reauthorize a Paypal account payment. Paypal recommends
// that a payment should be reauthorized after the initial 3-day honor period to
// ensure that funds are still available. Only paypal account payments can be re-
// authorized
func (c *Client) ReauthorizeAuthorization(authID string, a *Amount) (*Authorization, error) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/authorization/%s/reauthorize", c.APIBase, authID), struct {
		Amount *Amount `json:"amount"`
	}{a})
	if err != nil {
		return nil, err
	}

	v := &Authorization{}

	err = c.SendWithAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

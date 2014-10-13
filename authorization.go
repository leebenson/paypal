package paypal

import (
	"fmt"
	"net/http"
)

// https://developer.paypal.com/webapps/developer/docs/api/#authorizations

type (
	AuthWithAmountReq struct {
		Amount *Amount `json:"amount"`
	}
)

// GetAuthorization returns an authorization by ID
func (c *Client) GetAuthorization(authID string) (*Authorization, error, *http.Response) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/payments/authorization/%s", c.APIBase, authID), nil)
	if err != nil {
		return nil, err, nil
	}

	v := &Authorization{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// CaptureAuthorization captures and process an existing authorization.
// To use this method, the original payment must have Intent set to PaymentIntentAuthorize
func (c *Client) CaptureAuthorization(authID string, a *Amount, isFinalCapture bool) (*Capture, error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/authorization/%s/capture", c.APIBase, authID), struct {
		Amount         *Amount `json:"amount"`
		IsFinalCapture bool    `json:"is_final_capture"`
	}{a, isFinalCapture})
	if err != nil {
		return nil, err, nil
	}

	v := &Capture{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// VoidAuthorization voids a previously authorized payment. A fully
// captured authorization cannot be voided
func (c *Client) VoidAuthorization(authID string) (*Authorization, error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/authorization/%s/void", c.APIBase, authID), nil)
	if err != nil {
		return nil, err, nil
	}

	v := &Authorization{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// ReauthorizeAuthorization reauthorize a Paypal account payment. Paypal recommends
// that a payment should be reauthorized after the initial 3-day honor period to
// ensure that funds are still available. Only paypal account payments can be re-
// authorized
func (c *Client) ReauthorizeAuthorization(authID string, a *Amount) (*Authorization, error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/authorization/%s/reauthorize", c.APIBase, authID), struct {
		Amount *Amount `json:"amount"`
	}{a})
	if err != nil {
		return nil, err, nil
	}

	v := &Authorization{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

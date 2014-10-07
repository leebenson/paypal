package paypal

import (
	"fmt"
	"net/http"
)

// https://developer.paypal.com/webapps/developer/docs/api/#captures

// GetCapture returns details about a captured payment
func (c *Client) GetCapture(captureID string) (*Capture, error, *http.Response) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/payments/capture/%s", c.APIBase, captureID), nil)
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

// RefundCapture refund a captured payment. For partial refunds, a lower
// Amount object can be passed in.
func (c *Client) RefundCapture(captureID string, a *Amount) (*Refund, error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/capture/%s/refund", c.APIBase, captureID), struct {
		Amount *Amount `json:"amount"`
	}{a})
	if err != nil {
		return nil, err, nil
	}

	v := &Refund{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

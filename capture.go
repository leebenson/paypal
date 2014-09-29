package paypal

import "fmt"

// https://developer.paypal.com/webapps/developer/docs/api/#captures

// GetCapture returns details about a captured payment
func (c *Client) GetCapture(captureID string) (*Capture, error) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/payments/capture/%s", c.APIBase, captureID), nil)
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

// RefundCapture refund a captured payment. For partial refunds, a lower
// Amount object can be passed in.
func (c *Client) RefundCapture(captureID string, a *Amount) (*Refund, error) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/capture/%s/refund", c.APIBase, captureID), struct {
		Amount *Amount `json:"amount"`
	}{a})
	if err != nil {
		return nil, err
	}

	v := &Refund{}

	err = c.SendWithAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

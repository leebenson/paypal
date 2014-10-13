package paypal

import (
	"fmt"
	"net/http"
)

// https://developer.paypal.com/webapps/developer/docs/api/#orders

// GetOrder returns details about an order
func (c *Client) GetOrder(orderID string) (*Order, error, *http.Response) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/payments/orders/%s", c.APIBase, orderID), nil)
	if err != nil {
		return nil, err, nil
	}

	v := &Order{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// AuthorizeOrder authorizes an order
func (c *Client) AuthorizeOrder(orderID string, amount *Amount) (*Authorization, error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/orders/%s/authorize", c.APIBase, orderID), struct {
		Amount *Amount `json:"amount"`
	}{
		Amount: amount,
	})
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

// CaptureOrder captures a payment on an order. To use this call, an original payment
// must specify an "intent" of "order"
func (c *Client) CaptureOrder(orderID string, amount *Amount, isFinal bool) (*Capture, error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/orders/%s/capture", c.APIBase, orderID), struct {
		Amount         *Amount `json:"amount"`
		IsFinalCapture bool    `json:"is_final_capture"`
	}{
		Amount:         amount,
		IsFinalCapture: isFinal,
	})
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

// VoidOrder voids an existing order. An order cannot be voided if payment
// has already been partially or fully captured
func (c *Client) VoidOrder(orderID string) (*Order, error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/orders/%s/do-void", c.APIBase, orderID), nil)
	if err != nil {
		return nil, err, nil
	}

	v := &Order{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// RefundOrder refunds an exsting captured order. This only works after the
// order amount is captured. A refund cannot be made if the order is not captured.
// Alias for RefundCapture
func (c *Client) RefundOrder(captureID string, a *Amount) (*Refund, error, *http.Response) {
	return c.RefundCapture(captureID, a)
}

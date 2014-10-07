package paypal

import (
	"fmt"
	"net/http"
)

// https://developer.paypal.com/webapps/developer/docs/api/#payments

type (
	CreatePaymentResp struct {
		*Payment
		Links []Links `json:"links"`
	}

	ExecutePaymentResp struct {
		Intent       PaymentIntent `json:"intent"`
		Payer        *Payer        `json:"payer"`
		Transactions []Transaction `json:"transactions"`
		Links        []Links       `json:"links"`
	}

	ListPaymentsResp struct {
		Payments []Payment `json:"payments"`
	}
)

// CreatePayment creates a payment in Paypal
func (c *Client) CreatePayment(p Payment) (*CreatePaymentResp, error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/payment", c.APIBase), p)
	if err != nil {
		return nil, err, nil
	}

	v := &CreatePaymentResp{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// ExecutePayment completes an approved Paypal payment that has been approved by the payer
func (c *Client) ExecutePayment(paymentID, payerID string, transactions []Transaction) (*ExecutePaymentResp, error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/payment/%s/execute", c.APIBase, paymentID), struct {
		PayerID      string        `json:"payer_id"`
		Transactions []Transaction `json:"transactions"`
	}{
		payerID,
		transactions,
	})
	if err != nil {
		return nil, err, nil
	}

	v := &ExecutePaymentResp{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// GetPayment fetches a payment in Paypal
func (c *Client) GetPayment(id string) (*Payment, error, *http.Response) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/payments/payment/%s", c.APIBase, id), nil)
	if err != nil {
		return nil, err, nil
	}

	v := &Payment{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// ListPayments retrieve payments resources from Paypal
func (c *Client) ListPayments(filter map[string]string) ([]Payment, error, *http.Response) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/payments/payment/", c.APIBase), nil)
	if err != nil {
		return nil, err, nil
	}

	if filter != nil {
		q := req.URL.Query()

		for k, v := range filter {
			q.Set(k, v)
		}

		req.URL.RawQuery = q.Encode()
	}

	var v ListPaymentsResp

	resp, err := c.SendWithAuth(req, &v)
	if err != nil {
		return nil, err, resp
	}

	return v.Payments, nil, resp
}

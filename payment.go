package paypal

import (
	"fmt"
	"time"
)

type (
	CreatePaymentResp struct {
		Intent       PaymentIntent `json:"intent"`
		Payer        *Payer        `json:"payer"`
		Transactions []Transaction `json:"transactions"`
		RedirectURLs *RedirectURLs `json:"redirect_urls"`
		ID           string        `json:"id"`
		CreateTime   time.Time     `json:"create_time"`
		State        PaymentState  `json:"state"`
		UpdateTime   time.Time     `json:"update_time"`
		Links        []Links       `json:"links"`
	}

	ExecutePaymentResp struct {
		Intent       PaymentIntent `json:"intent"`
		Payer        *Payer        `json:"payer"`
		Transactions []Transaction `json:"transactions"`
		Links        []Links       `json:"links"`
	}

	PaymentResp struct {
		Intent       PaymentIntent `json:"intent"`
		Payer        *Payer        `json:"payer"`
		Transactions []Transaction `json:"transactions"`
		RedirectURLs *RedirectURLs `json:"redirect_urls"`
		ID           string        `json:"id"`
		CreateTime   time.Time     `json:"create_time"`
		State        PaymentState  `json:"state"`
		UpdateTime   time.Time     `json:"update_time"`
	}

	ListPaymentsResp struct {
		Payments []PaymentResp `json:"payments"`
	}
)

// CreatePayment creates a payment in Paypal
func (c *Client) CreatePayment(p Payment) (*CreatePaymentResp, error) {
	req, err := NewRequest("POST", fmt.Sprintf("%s%s", c.APIBase, "/payments/payment"), p)
	if err != nil {
		return nil, err
	}

	v := &CreatePaymentResp{}

	err = c.SendWithAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// ExecutePayment completes an approved Paypal payment that has been approved by the payer
func (c *Client) ExecutePayment(paymentID, payerID string, transactions []Transaction) (*ExecutePaymentResp, error) {
	req, err := NewRequest("POST", fmt.Sprintf("%s%s%s/execute", c.APIBase, "/payments/payment/", paymentID), struct {
		PayerID      string        `json:"payer_id"`
		Transactions []Transaction `json:"transactions"`
	}{
		payerID,
		transactions,
	})
	if err != nil {
		return nil, err
	}

	v := &ExecutePaymentResp{}

	err = c.SendWithAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// GetPayment fetches a payment in Paypal
func (c *Client) GetPayment(id string) (*PaymentResp, error) {
	req, err := NewRequest("GET", fmt.Sprintf("%s%s%s", c.APIBase, "/payments/payment/", id), nil)
	if err != nil {
		return nil, err
	}

	v := &PaymentResp{}

	err = c.SendWithAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// ListPayments retrieve payments resources from Paypal
func (c *Client) ListPayments(filter map[string]string) ([]PaymentResp, error) {
	req, err := NewRequest("GET", fmt.Sprintf("%s%s", c.APIBase, "/payments/payment/"), nil)
	if err != nil {
		return nil, err
	}

	if filter != nil {
		q := req.URL.Query()

		for k, v := range filter {
			q.Set(k, v)
		}

		req.URL.RawQuery = q.Encode()
	}

	var v ListPaymentsResp

	err = c.SendWithAuth(req, &v)
	if err != nil {
		return nil, err
	}

	return v.Payments, nil
}

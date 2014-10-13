package paypal

import (
	"fmt"
	"net/http"
)

type (
	RefundReq struct {
		Amount *Amount `json:"amount"`
	}
)

// GetSales returns a sale by ID
func (c *Client) GetSale(saleID string) (*Sale, error, *http.Response) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/payments/sale/%s", c.APIBase, saleID), nil)
	if err != nil {
		return nil, err, nil
	}

	v := &Sale{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// RefundSale refunds a completed payment and accepts an optional
// Amount struct. If Amount is provided, a partial refund is requested,
// or else a full refund is made instead
func (c *Client) RefundSale(saleID string, a *Amount) (*Refund, error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/sale/%s/refund", c.APIBase, saleID), &RefundReq{Amount: a})
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

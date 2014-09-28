package paypal

import (
	"fmt"
)

type (
	RefundReq struct {
		Amount *Amount `json:"amount"`
	}
)

// GetSales returns a sale by ID
func (c *Client) GetSale(saleID string) (*Sale, error) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/payments/sale/%s", c.APIBase, saleID), nil)
	if err != nil {
		return nil, err
	}

	v := &Sale{}

	err = c.SendWithAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// RefundSale refunds a completed payment and accepts an optional
// Amount struct. If Amount is provided, a partial refund is requested,
// or else a full refund is made instead
func (c *Client) RefundSale(saleID string, a *Amount) (*Refund, error) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/sale/%s/refund", c.APIBase, saleID), &RefundReq{Amount: a})
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

package paypal

import "fmt"

// GetRefund returns a refund by ID
func (c *Client) GetRefund(refundID string) (*Refund, error) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/payments/refund/%s", c.APIBase, refundID), nil)
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

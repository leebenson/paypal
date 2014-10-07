package paypal

import (
	"fmt"
	"net/http"
)

// GetRefund returns a refund by ID
func (c *Client) GetRefund(refundID string) (*Refund, error, *http.Response) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/refund/%s", c.APIBase, refundID), nil)
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

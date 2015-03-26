package paypal

import (
	"fmt"
	"time"
)

// https://developer.paypal.com/docs/api/#vault

type (
	// VaultRequest maps to vault_request object
	VaultRequest struct {
		CreditCard
		MerchantID     string `json:"merchant_id,omitempty"`
		ExternalCardID string `json:"external_card_id,omitempty"`
	}

	// VaultResponse maps to vault_response object
	VaultResponse struct {
		VaultRequest
		CreateTime *time.Time `json: "create_time"`
		UpdateTime *time.Time `json: "update_time"`
		State      string     `json: "state"`
		ValidUntil string     `json: "valid_until"`
		Links      []Links    `json:"links"`
	}
)

// StoreInVault will store credit card details with PayPal.
func (c *Client) StoreInVault(cc VaultRequest) (*VaultResponse, error) {

	req, err := NewRequest("POST", fmt.Sprintf("%s/vault/credit-cards", c.APIBase), cc)

	if err != nil {
		return nil, err
	}
	v := &VaultResponse{}

	err = c.SendWithAuth(req, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

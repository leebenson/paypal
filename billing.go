package paypal

import (
	"fmt"
	"net/http"
)

// https://developer.paypal.com/webapps/developer/docs/api/#billing-plans-and-agreements

type (
	ListBillingPlansResp struct {
		Plans []Plan `json:"plans"`
	}
)

// CreateBillingPlan creates an empty billing plan. By default, a created billing
// plan is in a CREATED state. A user cannot subscribe to the billing plan
// unless it has been set to the ACTIVE state.
func (c *Client) CreateBillingPlan(p *Plan) (*Plan, error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/billing-plans", c.APIBase), p)
	if err != nil {
		return nil, err, nil
	}

	v := &Plan{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// UpdateBillingPlan updates data of an existing billing plan. The state of a plan
// must be PlanStateActive before a billing agreement is created
func (c *Client) UpdateBillingPlan(p *Plan) (error, *http.Response) {
	req, err := NewRequest("PATCH", fmt.Sprintf("%s/payments/billing-plans/%s", c.APIBase, p.ID), struct {
		Path  string         `json:"path"`
		Value *Plan          `json:"value"`
		OP    PatchOperation `json:"op"`
	}{
		Path:  "/",
		Value: p,
		OP:    PatchOperationReplace,
	})
	if err != nil {
		return err, nil
	}

	v := &struct{}{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return err, resp
	}

	return nil, resp
}

// GetBillingPlan returns details about a specific billing plan
func (c *Client) GetBillingPlan(planID string) (*Plan, error, *http.Response) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/payments/billing-plans/%s", c.APIBase, planID), nil)
	if err != nil {
		return nil, err, nil
	}

	v := &Plan{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// ListBillingPlans returns billing plans based on their current state: created
// active or deactivated
func (c *Client) ListBillingPlans(filter map[string]string) ([]Plan, error, *http.Response) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/payments/billing-plans", c.APIBase), nil)
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

	var v ListBillingPlansResp

	resp, err := c.SendWithAuth(req, &v)
	if err != nil {
		return nil, err, resp
	}

	return v.Plans, nil, resp
}

// CreateAgreement creates a billing agreement for the buyer. The EC token generates,
// and the buyer must click an approval URL. Through the approval URL, you obtain
// buyer details and the shipping address. After buyer approval, call the execute
// URL to create the billing agreement in the system.
func (c *Client) CreateAgreement(a *Agreement) (*Agreement, error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/billing-agreements", c.APIBase), a)
	if err != nil {
		return nil, err, nil
	}

	v := &Agreement{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// ExecuteAgreement executes an agreement after the buyer approves it.
func (c *Client) ExecuteAgreement(paymentID string) (*Agreement, error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/billing-agreements/%s/agreement-execute", c.APIBase, paymentID), nil)
	if err != nil {
		return nil, err, nil
	}

	v := &Agreement{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// UpdateAgreement updates an agreement
func (c *Client) UpdateAgreement(a *Agreement) (error, *http.Response) {
	req, err := NewRequest("PATCH", fmt.Sprintf("%s/payments/billing-agreements/%s", c.APIBase, a.ID), struct {
		Path  string         `json:"path"`
		Value *Agreement     `json:"value"`
		OP    PatchOperation `json:"op"`
	}{
		Path:  "/",
		Value: a,
		OP:    PatchOperationReplace,
	})
	if err != nil {
		return err, nil
	}

	v := &struct{}{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return err, resp
	}

	return nil, resp
}

// GetAgreement returns an agreement
func (c *Client) GetAgreement(agreementID string) (*Agreement, error, *http.Response) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/payments/billing-agreements/%s", c.APIBase, agreementID), nil)
	if err != nil {
		return nil, err, nil
	}

	v := &Agreement{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// SuspendAgreement suspends an agreement
func (c *Client) SuspendAgreement(agreementID, note string) (error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/billing-agreements/%s/suspend", c.APIBase, agreementID), struct {
		Note string `json:"note"`
	}{
		Note: note,
	})
	if err != nil {
		return err, nil
	}

	v := &struct{}{}

	resp, err := c.SendWithAuth(req, v)

	return err, resp
}

// ReactivateAgreement reactivate an agreement
func (c *Client) ReactivateAgreement(agreementID, note string) (error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/billing-agreements/%s/re-activate", c.APIBase, agreementID), struct {
		Note string `json:"note"`
	}{
		Note: note,
	})
	if err != nil {
		return err, nil
	}

	v := &struct{}{}

	resp, err := c.SendWithAuth(req, v)

	return err, resp
}

// SearchAgreementTransactions searches for transactions within a billing agreement
func (c *Client) SearchAgreementTransactions(agreementID string, filter map[string]string) (*AgreementTransactions, error, *http.Response) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/payments/billing-agreements/%s/transaction", c.APIBase, agreementID), nil)
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

	v := &AgreementTransactions{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// CancelAgreement cancels an agreement
func (c *Client) CancelAgreement(agreementID, note string) (error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/billing-agreements/%s/cancel", c.APIBase, agreementID), struct {
		Note string `json:"note"`
	}{
		Note: note,
	})
	if err != nil {
		return err, nil
	}

	v := &struct{}{}

	resp, err := c.SendWithAuth(req, v)

	return err, resp
}

// SetAgreementBalance sets the outstanding amount of an agreement
func (c *Client) SetAgreementBalance(agreementID string, currency *Currency) (error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/billing-agreements/%s/set-balance", c.APIBase, agreementID), currency)
	if err != nil {
		return err, nil
	}

	v := &struct{}{}

	resp, err := c.SendWithAuth(req, v)

	return err, resp
}

// BillAgreementBalance bills the outstanding amount of an agreement
func (c *Client) BillAgreementBalance(agreementID string, currency *Currency, note string) (error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/billing-agreements/%s/bill-balance", c.APIBase, agreementID), struct {
		Note   string    `json:"note"`
		Amount *Currency `json:"amount"`
	}{
		Note:   note,
		Amount: currency,
	})
	if err != nil {
		return err, nil
	}

	v := &struct{}{}

	resp, err := c.SendWithAuth(req, v)

	return err, resp
}

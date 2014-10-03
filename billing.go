package paypal

import "fmt"

// https://developer.paypal.com/webapps/developer/docs/api/#billing-plans-and-agreements

type (
	ListBillingPlansResp struct {
		Plans []Plan `json:"plans"`
	}
)

// CreateBillingPlan creates an empty billing plan. By default, a created billing
// plan is in a CREATED state. A user cannot subscribe to the billing plan
// unless it has been set to the ACTIVE state.
func (c *Client) CreateBillingPlan(p *Plan) (*Plan, error) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/billing-plans", c.APIBase), p)
	if err != nil {
		return nil, err
	}

	v := &Plan{}

	err = c.SendWithAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// UpdateBillingPlan updates data of an existing billing plan. The state of a plan
// must be PlanStateActive before a billing agreement is created
func (c *Client) UpdateBillingPlan(p *Plan) error {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/billing-plans", c.APIBase), struct {
		Path  string         `json:"path"`
		Value *Plan          `json:"value"`
		OP    PatchOperation `json:"op"`
	}{
		Path:  "/",
		Value: p,
		OP:    PatchOperationReplace,
	})
	if err != nil {
		return err
	}

	v := &struct{}{}

	err = c.SendWithAuth(req, v)
	if err != nil {
		return err
	}

	return nil
}

// GetBillingPlan returns details about a specific billing plan
func (c *Client) GetBillingPlan(planID string) (*Plan, error) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/payments/billing-plans/%s", c.APIBase, planID), nil)
	if err != nil {
		return nil, err
	}

	v := &Plan{}

	err = c.SendWithAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// ListBillingPlans returns billing plans based on their current state: created
// active or deactivated
func (c *Client) ListBillingPlans(filter map[string]string) ([]Plan, error) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/payments/billing-plans", c.APIBase), nil)
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

	var v ListBillingPlansResp

	err = c.SendWithAuth(req, &v)
	if err != nil {
		return nil, err
	}

	return v.Plans, nil
}

// CreateAgreement creates a billing agreement for the buyer. The EC token generates,
// and the buyer must click an approval URL. Through the approval URL, you obtain
// buyer details and the shipping address. After buyer approval, call the execute
// URL to create the billing agreement in the system.
func (c *Client) CreateAgreement(a *Agreement) (*Agreement, error) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/billing-agreements", c.APIBase), a)
	if err != nil {
		return nil, err
	}

	v := &Agreement{}

	err = c.SendWithAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

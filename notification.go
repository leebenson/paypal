package paypal

import (
	"fmt"
	"net/http"
)

// https://developer.paypal.com/webapps/developer/docs/api/#notifications

// ListWebhookEventTypes returns a list of events types that are available to any
// webhook for subscription
func (c *Client) ListWebhookEventTypes() ([]EventType, error, *http.Response) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/notifications/webhooks-event-types", c.APIBase), nil)
	if err != nil {
		return nil, err, nil
	}

	var v struct {
		EventTypes []EventType `json:"event_types"`
	}

	resp, err := c.SendWithAuth(req, &v)
	if err != nil {
		return nil, err, resp
	}

	return v.EventTypes, nil, resp
}

// CreateWebhook creates a webhook. The maximum number of webhooks allowed to be registered
// is 10.
func (c *Client) CreateWebhook(w *Webhook) (*Webhook, error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/notifications/webhooks", c.APIBase), w)
	if err != nil {
		return nil, err, nil
	}

	v := &Webhook{}

	resp, err := c.SendWithAuth(req, &v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// GetWebhook returns a specific webhook.
func (c *Client) GetWebhook(webhookID string) (*Webhook, error, *http.Response) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/notifications/webhooks/%s", c.APIBase, webhookID), nil)
	if err != nil {
		return nil, err, nil
	}

	v := &Webhook{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// ListEventTypesByWebhook returns a list of events types that are subscribed to a webhook
func (c *Client) ListEventTypesByWebhook(webhookID string) ([]EventType, error, *http.Response) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/notifications/webhooks/%s/event-types", c.APIBase, webhookID), nil)
	if err != nil {
		return nil, err, nil
	}

	var v struct {
		EventTypes []EventType `json:"event_types"`
	}

	resp, err := c.SendWithAuth(req, &v)
	if err != nil {
		return nil, err, resp
	}

	return v.EventTypes, nil, resp
}

// ListWebhooks returns all webhooks
func (c *Client) ListWebhooks() ([]Webhook, error, *http.Response) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/notifications/webhooks-event-types", c.APIBase), nil)
	if err != nil {
		return nil, err, nil
	}

	var v struct {
		Webhooks []Webhook `json:"webhooks"`
	}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v.Webhooks, nil, resp
}

// UpdateWebhook updates a webhook
func (c *Client) UpdateWebhook(w *Webhook) (*Webhook, error, *http.Response) {
	req, err := NewRequest("PATCH", fmt.Sprintf("%s/notifications/webhooks/%s", c.APIBase, w.ID), struct {
		Path  string         `json:"path"`
		Value *Webhook       `json:"value"`
		OP    PatchOperation `json:"op"`
	}{
		Path:  "/",
		Value: w,
		OP:    PatchOperationReplace,
	})
	if err != nil {
		return nil, err, nil
	}

	v := &Webhook{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// DeleteWebhook delets a webhook
func (c *Client) DeleteWebhook(webhookID string) (error, *http.Response) {
	req, err := NewRequest("DELETE", fmt.Sprintf("%s/notifications/webhooks/%s", c.APIBase, webhookID), nil)
	if err != nil {
		return err, nil
	}

	v := &struct{}{}

	resp, err := c.SendWithAuth(req, &v)
	if err != nil {
		return err, resp
	}

	return nil, resp
}

// GetWebhookEvent returns a webhook event
func (c *Client) GetWebhookEvent(eventID string) (*Event, error, *http.Response) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/notifications/webhooks-events/%s", c.APIBase, eventID), nil)
	if err != nil {
		return nil, err, nil
	}

	v := &Event{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// SearchWebhookEvents searches for all webhook events
func (c *Client) SearchWebhookEvents(s *WebhookEventSearch) (*EventList, error, *http.Response) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/notifications/webhooks-events", c.APIBase), s)
	if err != nil {
		return nil, err, nil
	}

	v := &EventList{}

	resp, err := c.SendWithAuth(req, &v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// ResendWebhookEvent resend the event notification
func (c *Client) ResendWebhookEvent(eventID string) (*Event, error, *http.Response) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/notifications/webhooks-events/%s/resend", c.APIBase, eventID), nil)
	if err != nil {
		return nil, err, nil
	}

	v := &Event{}

	resp, err := c.SendWithAuth(req, &v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

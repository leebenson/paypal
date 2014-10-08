package paypal

import "time"

// https://developer.paypal.com/webapps/developer/docs/api/#common-notifications-objects

type (
	// EventList maps to event_list object. List of Webhooks event resources
	EventList struct {
		Events []Event `json:"events"`
		Count  int     `json:"count"`
		Links  []Links `json:"links"`
	}

	// EventType maps to event_type object. Contaisn the information for a Webhooks event-type
	EventType struct {
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
	}

	// Webhook maps to webhook object. Represents Webhook resource
	Webhook struct {
		ID         string      `json:"id,omitempty"`
		URL        string      `json:"url"`
		EventTypes []EventType `json:"event_types"`
		Links      []Links     `json:"links,omitempty"`
	}

	// Event maps to event object. Represents a Webhooks event
	Event struct {
		ID           string      `json:"id,omitempty"`
		CreateTime   *time.Time  `json:"create_time,omitempty"`
		ResourceType string      `json:"resource_type,omitempty"`
		EventType    string      `json:"event_type,omitempty"`
		Summary      string      `json:"summary,omitempty"`
		Resource     interface{} `json:"resource,omitempty"`
		Links        []Links     `json:"links,omitempty"`
	}

	// WebhookEventSearch is search parameters for webhook events
	WebhookEventSearch struct {
		PageSize  int        `json:"page_size,omitempty"`
		StartTime *time.Time `json:"start_time,omitempty"`
		EndTime   *time.Time `json:"end_time,omitempty"`
	}
)

package paypal

type (

	// Links maps to links object
	Links struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
		// TODO: Support HyperSchema with its multiple types per field
		// TargetSchema HyperSchema `json:"targetSchema"`
		Method  string `json:"method"`
		Enctype string `json:"enctype"`
		// Schema HyperSchema `json:"schema"`
	}
)

package model

// Webhook model
type Webhook struct {
	// Name of the resource. Can be freely chosen by the user.
	Name *string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description *string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *DateTime `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]interface{} `json:"customData,omitempty"`
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Webhook url (required)
	Url *string `json:"url,omitempty"`
	// HTTP method used for the webhook
	Method WebhookHttpMethod `json:"method,omitempty"`
	// Whether to skip SSL certification verification or not
	InsecureSsl *bool `json:"insecureSsl,omitempty"`
	// Signature used for the webhook
	Signature *WebhookSignature `json:"signature,omitempty"`
	// The json schema of the data that is send as webhook payload
	Schema *map[string]interface{} `json:"schema,omitempty"`
}

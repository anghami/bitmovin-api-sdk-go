package model

// StreamsConfigResponse model
type StreamsConfigResponse struct {
	// The identifier of the stream config
	Id *string `json:"id,omitempty"`
	// UUID of the associated organization
	OrgId *string `json:"orgId,omitempty"`
	// Player style config
	PlayerStyle *map[string]interface{} `json:"playerStyle,omitempty"`
}
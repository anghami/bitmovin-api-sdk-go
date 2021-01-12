package model

// GceAccountRegionSettings model
type GceAccountRegionSettings struct {
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
	// Id of the network for encoding instances (required)
	Network *string `json:"network,omitempty"`
	// Id of the subnet for encoding instances (required)
	SubnetId *string           `json:"subnetId,omitempty"`
	Region   GoogleCloudRegion `json:"region,omitempty"`
}

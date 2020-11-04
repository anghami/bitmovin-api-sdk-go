package model

import (
	"encoding/json"
)

// InterlaceFilter model
type InterlaceFilter struct {
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
	Id                           *string                      `json:"id,omitempty"`
	Mode                         InterlaceMode                `json:"mode,omitempty"`
	VerticalLowPassFilteringMode VerticalLowPassFilteringMode `json:"verticalLowPassFilteringMode,omitempty"`
}

func (m InterlaceFilter) FilterType() FilterType {
	return FilterType_INTERLACE
}
func (m InterlaceFilter) MarshalJSON() ([]byte, error) {
	type M InterlaceFilter
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "INTERLACE"

	return json.Marshal(x)
}

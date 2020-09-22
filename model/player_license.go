package model

// PlayerLicense model
type PlayerLicense struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Name of the resource (required)
	Name *string `json:"name,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ (required)
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// License Key (required)
	LicenseKey *string `json:"licenseKey,omitempty"`
	// Number of impressions recorded (required)
	Impressions *int32 `json:"impressions,omitempty"`
	// Maximum number of impressions (required)
	MaxImpressions *int32 `json:"maxImpressions,omitempty"`
	// Flag if third party licensing is enabled (required)
	ThirdPartyLicensingEnabled *bool `json:"thirdPartyLicensingEnabled,omitempty"`
	// Whitelisted domains (required)
	Domains []Domain `json:"domains,omitempty"`
	// Analytics License Key
	AnalyticsKey *string `json:"analyticsKey,omitempty"`
}

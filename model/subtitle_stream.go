package model

// SubtitleStream model
type SubtitleStream struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Position of the stream in the media, starting from 0.
	Position *int32 `json:"position,omitempty"`
	// Duration of the stream in seconds
	Duration *float64 `json:"duration,omitempty"`
	// Codec of the stream
	Codec *string `json:"codec,omitempty"`
	// Language of the stream
	Language *string `json:"language,omitempty"`
	// Hearing impaired support
	HearingImpaired *bool `json:"hearingImpaired,omitempty"`
}

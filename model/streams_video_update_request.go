package model

// StreamsVideoUpdateRequest model
type StreamsVideoUpdateRequest struct {
	// The new status of the stream
	Status StreamsVideoStatus `json:"status,omitempty"`
	// The new title of the stream
	Title *string `json:"title,omitempty"`
	// The new description of the stream
	Description *string `json:"description,omitempty"`
	// Id of the stream config to use
	ConfigId *string `json:"configId,omitempty"`
}

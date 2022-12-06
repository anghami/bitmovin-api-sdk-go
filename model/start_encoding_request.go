package model

// StartEncodingRequest model
type StartEncodingRequest struct {
	// Allows to encode only part of the input. Defines start (offset) and duration of the desired section. This is not allowed when the Encoding uses any kind of Input Stream resource.
	Trimming *Trimming `json:"trimming,omitempty"`
	// Scheduling parameters of the encoding.
	Scheduling *Scheduling `json:"scheduling,omitempty"`
	// Special tweaks for your encoding job.
	Tweaks *Tweaks `json:"tweaks,omitempty"`
	// Enable frame dropping/duplication to handle variable frames per seconds of video input streams
	HandleVariableInputFps *bool `json:"handleVariableInputFps,omitempty"`
	// The pass mode of the encoding
	EncodingMode EncodingMode `json:"encodingMode,omitempty"`
	// List of preview DASH manifests to be created
	PreviewDashManifests []ManifestResource `json:"previewDashManifests,omitempty"`
	// List of preview HLS manifests to be created
	PreviewHlsManifests []ManifestResource `json:"previewHlsManifests,omitempty"`
	// List of VoD DASH manifests to be created after encoding finished successfully
	VodDashManifests []ManifestResource `json:"vodDashManifests,omitempty"`
	// List of VoD HLS manifests to be created after encoding finished successfully
	VodHlsManifests []ManifestResource `json:"vodHlsManifests,omitempty"`
	// List of VoD SMOOTH manifests to be created after encoding finished successfully
	VodSmoothManifests []ManifestResource `json:"vodSmoothManifests,omitempty"`
	// Version of the manifest generation engine to be used
	ManifestGenerator ManifestGenerator `json:"manifestGenerator,omitempty"`
	// Per-Title settings
	PerTitle *PerTitle `json:"perTitle,omitempty"`
}

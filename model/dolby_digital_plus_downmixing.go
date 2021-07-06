package model

// Downmixing is used to reproduce the complete audio program when the actual decoder outputs do not match the encoded channel layout of the original audio signal.  The process of downmixing takes the information in the channels that do not have corresponding outputs, and mixes this information into the remaining channels.
type DolbyDigitalPlusDownmixing struct {
	// The level shift applied to the C channel when adding to the L and R outputs as a result of downmixing to one Lo/Ro output.
	LoRoCenterMixLevel DolbyDigitalPlusCenterMixLevel `json:"loRoCenterMixLevel,omitempty"`
	// The level shift applied to the C channel when adding to the L and R outputs as a result of downmixing to one Lt/Rt output.
	LtRtCenterMixLevel DolbyDigitalPlusCenterMixLevel `json:"ltRtCenterMixLevel,omitempty"`
	// The level shift applied to the surround channels when downmixing to one Lo/Ro output.
	LoRoSurroundMixLevel DolbyDigitalPlusSurroundMixLevel `json:"loRoSurroundMixLevel,omitempty"`
	// The level shift applied to the surround channels when downmixing to one Lt/Rt output.
	LtRtSurroundMixLevel DolbyDigitalPlusSurroundMixLevel        `json:"ltRtSurroundMixLevel,omitempty"`
	PreferredMode        DolbyDigitalPlusDownmixingPreferredMode `json:"preferredMode,omitempty"`
}

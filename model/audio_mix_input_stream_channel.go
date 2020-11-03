package model


// AudioMixInputStreamChannel model
type AudioMixInputStreamChannel struct {
    // The id of the input stream that should be used for mixing.
    InputStreamId *string `json:"inputStreamId,omitempty"`
    OutputChannelType AudioMixChannelType `json:"outputChannelType,omitempty"`
    // Number of this output channel. If type is 'CHANNEL_NUMBER', this must be set.
    OutputChannelNumber *int32 `json:"outputChannelNumber,omitempty"`
    // List of source channels to be mixed
    SourceChannels []AudioMixInputStreamSourceChannel `json:"sourceChannels,omitempty"`
}




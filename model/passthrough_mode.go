package model


// PassthroughMode : Configure if the encoding should use the video stream as the passthrough mode or a dedicated caption stream.
type PassthroughMode string

// List of possible PassthroughMode values
const (
    PassthroughMode_VIDEO_STREAM PassthroughMode = "VIDEO_STREAM"
    PassthroughMode_CAPTION_STREAM PassthroughMode = "CAPTION_STREAM"
    PassthroughMode_VIDEO_CAPTION_STREAM PassthroughMode = "VIDEO_CAPTION_STREAM"
)



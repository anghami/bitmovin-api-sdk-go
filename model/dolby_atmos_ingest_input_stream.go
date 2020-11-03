package model

import (
    "encoding/json"
)

// DolbyAtmosIngestInputStream model
type DolbyAtmosIngestInputStream struct {
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
    // Id of input (required)
    InputId *string `json:"inputId,omitempty"`
    // Path to the Dolby Atmos input file (required)
    InputPath *string `json:"inputPath,omitempty"`
    // Input file format of the Dolby Atmos input file.  Set it to DAMF if the given input file is a Dolby Atmos Master File (.atmos). Set it to ADM if the given input file is an Audio Definition Model Broadcast Wave Format file (.wav) (required)
    InputFormat DolbyAtmosInputFormat `json:"inputFormat,omitempty"`
}

func (m DolbyAtmosIngestInputStream) InputStreamType() InputStreamType {
    return InputStreamType_DOLBY_ATMOS
}
func (m DolbyAtmosIngestInputStream) MarshalJSON() ([]byte, error) {
    type M DolbyAtmosIngestInputStream
    x := struct {
        Type string `json:"type"`
        M
    }{M: M(m)}

    x.Type = "DOLBY_ATMOS"

    return json.Marshal(x)
}



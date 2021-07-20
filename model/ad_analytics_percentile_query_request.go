package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
)

// AdAnalyticsPercentileQueryRequest model
type AdAnalyticsPercentileQueryRequest struct {
	// Start of timeframe which is queried in UTC format.
	Start *DateTime `json:"start,omitempty"`
	// End of timeframe which is queried in UTC format.
	End *DateTime `json:"end,omitempty"`
	// Analytics license key (required)
	LicenseKey *string                     `json:"licenseKey,omitempty"`
	Filters    []AdAnalyticsAbstractFilter `json:"filters,omitempty"`
	OrderBy    []AdAnalyticsOrderByEntry   `json:"orderBy,omitempty"`
	Dimension  AdAnalyticsAttribute        `json:"dimension,omitempty"`
	Interval   AnalyticsInterval           `json:"interval,omitempty"`
	GroupBy    []AdAnalyticsAttribute      `json:"groupBy,omitempty"`
	// Maximum number of rows returned (max. 200)
	Limit *int64 `json:"limit,omitempty"`
	// Offset of data
	Offset *int64 `json:"offset,omitempty"`
	// The percentage (0-99) used for percentile queries.
	Percentile *int64 `json:"percentile,omitempty"`
}

// UnmarshalJSON unmarshals model AdAnalyticsPercentileQueryRequest from a JSON structure
func (m *AdAnalyticsPercentileQueryRequest) UnmarshalJSON(raw []byte) error {
	var data struct {
		Start      *DateTime                 `json:"start"`
		End        *DateTime                 `json:"end"`
		LicenseKey *string                   `json:"licenseKey"`
		Filters    json.RawMessage           `json:"filters"`
		OrderBy    []AdAnalyticsOrderByEntry `json:"orderBy"`
		Dimension  AdAnalyticsAttribute      `json:"dimension"`
		Interval   AnalyticsInterval         `json:"interval"`
		GroupBy    []AdAnalyticsAttribute    `json:"groupBy"`
		Limit      *int64                    `json:"limit"`
		Offset     *int64                    `json:"offset"`
		Percentile *int64                    `json:"percentile"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var result AdAnalyticsPercentileQueryRequest

	result.Start = data.Start
	result.End = data.End
	result.LicenseKey = data.LicenseKey
	result.OrderBy = data.OrderBy
	result.Dimension = data.Dimension
	result.Interval = data.Interval
	result.GroupBy = data.GroupBy
	result.Limit = data.Limit
	result.Offset = data.Offset
	result.Percentile = data.Percentile

	allOfFilters, err := UnmarshalAdAnalyticsAbstractFilterSlice(bytes.NewBuffer(data.Filters), bitutils.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	result.Filters = allOfFilters

	*m = result

	return nil
}

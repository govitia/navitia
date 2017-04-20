package types

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// UnmarshalJSON implements json.Unmarshaller for a Disruption
func (d *Disruption) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	// We define some of the value as pointers to the real values, allowing us to bypass copying in cases where we don't need to process the data
	data := &struct {
		// The references
		ID                *ID               `json:"id"`
		Status            *string           `json:"status"`
		InputDisruptionID *ID               `json:"disruption_id"`
		InputImpactID     *ID               `json:"impact_id"`
		Severity          *Severity         `json:"severity"`
		Periods           *[]Period         `json:"application_periods"`
		Messages          *[]Message        `json:"messages"`
		Impacted          *[]ImpactedObject `json:"impacted_objects"`
		Cause             *string           `json:"cause"`
		Category          *string           `json:"category"`

		// Those we will process
		LastUpdated string `json:"updated_at"`
	}{
		ID:                &d.ID,
		Status:            &d.Status,
		InputDisruptionID: &d.InputDisruptionID,
		InputImpactID:     &d.InputImpactID,
		Severity:          &d.Severity,
		Periods:           &d.Periods,
		Messages:          &d.Messages,
		Impacted:          &d.Impacted,
		Cause:             &d.Cause,
		Category:          &d.Category,
	}

	// Let's create the error generator
	gen := unmarshalErrorMaker{"Disruption"}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "Error while unmarshalling Disruption")
	}

	// Now we process the Update time
	d.LastUpdated, err = parseDateTime(data.LastUpdated)
	if err != nil {
		return gen.err(err, "LastUpdated", "updated_at", data.LastUpdated, "parseDateTime failed")
	}

	// Finished !
	return nil
}

// UnmarshalJSON implements json.Unmarshaller for a Period
func (p *Period) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	// We define some of the value as pointers to the real values, allowing us to bypass copying in cases where we don't need to process the data
	data := &struct {
		// Those we will process
		Begin string `json:"begin"`
		End   string `json:"end"`
	}{}

	// Let's create the error generator
	gen := unmarshalErrorMaker{"Period"}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "Error while unmarshalling Period")
	}

	// Now we process the times.
	p.Begin, err = parseDateTime(data.Begin)
	if err != nil {
		return gen.err(err, "Begin", "begin", data.Begin, "parseDateTime failed")
	}
	p.End, err = parseDateTime(data.End)
	if err != nil {
		return gen.err(err, "End", "end", data.End, "parseDateTime failed")
	}

	// Finished !
	return nil
}

// UnmarshalJSON implements json.Unmarshaller for a Severity
func (s *Severity) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	// We define some of the value as pointers to the real values, allowing us to bypass copying in cases where we don't need to process the data
	data := &struct {
		// The references
		Name     *string `json:"name"`
		Priority *int    `json:"priority,omitempty"` // As priority can be null, and 0 is the highest priority.
		Effect   *Effect `json:"effect"`

		// Those we will process
		Color string `json:"color"`
	}{
		Name:     &s.Name,
		Priority: s.Priority,
		Effect:   &s.Effect,
	}

	// Let's create the error generator
	gen := unmarshalErrorMaker{"Severity"}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "Error while unmarshalling Disruption")
	}

	// Process the color
	if str := data.Color; len(str) == 6 {
		clr, err := parseColor(str)
		if err != nil {
			return gen.err(err, "Color", "color", str, "error in parseColor")
		}
		s.Color = clr
	}

	return nil
}

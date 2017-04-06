package types

import (
	"encoding/json"
	"strconv"
	"time"
)

func parseDateTime(datetime string) (time.Time, error) {
	//TODO
	return time.Now(), nil
}

// Journey implements json.Unmarshaller
func (j Journey) UnmarshalJSON(b []byte) error {
	var data map[string]string
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	for key, value := range data {
		switch {
		case key == "duration":
			parsed, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
			j.Duration = time.Duration(parsed) / time.Second
		case key == "transfers":
			parsed, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return err
			}
			j.Transfers = uint(parsed)
		case key == "departure_date_time":
			time, err := parseDateTime(value)
			if err != nil {
				return err
			}
			j.Departure = time
		case key == "requested_date_time":
			time, err := parseDateTime(value)
			if err != nil {
				return err
			}
			j.Requested = time
		case key == "arrival_date_time":
			time, err := parseDateTime(value)
			if err != nil {
				return err
			}
			j.Arrival = time
		case key == "sections":
			var sections = &[]Section{}
			// Parse
			err := json.Unmarshal([]byte(value), sections)
			if err != nil {
				return err
			}
		case key == "type":
			j.Type = JourneyQualification(value)
		}
	}
	return nil
}

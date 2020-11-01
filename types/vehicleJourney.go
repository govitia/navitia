package types

// VehicleJourney gives informations on vehicle transportation schedule and details.
type VehicleJourney struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	Codes           []Code          `json:"codes"`
	Disruptions     []Disruption    `json:"disruptions"`
	Calendars       []Calendar      `json:"calendars"`
	StopTimes       []StopTime      `json:"stop_times"`
	ValidityPattern ValidityPattern `json:"validity_pattern"`
	JourneyPattern  JourneyPattern  `json:"journey_pattern"`
	Headsign        string          `json:"headsign"`
	Trip            Trip            `json:"trip"`
}

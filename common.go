package types

// DataFreshness codes for a specific data freshness requirement: realtime or base_schedule
type DataFreshness string

const (
	// DataFreshnessRealTime means you'll get undisrupted journeys
	DataFreshnessRealTime DataFreshness = "realtime"
	// DataFreshnessBaseSchedule means you can get disrupted journeys in the response.
	DataFreshnessBaseSchedule = "base_schedule"
)

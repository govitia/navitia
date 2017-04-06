package types

type ID interface {
	String() string
	FormatURL() string
}

// DataFreshness codes for a specific data freshness requirement: realtime or base_schedule
type DataFreshness string

const (
	// When using the following parameter, you'll get undisrupted journeys
	DataFreshnessRealTime DataFreshness = "realtime"
	// When using the following parameter you can get disrupted journeys in the response.
	DataFreshnessBaseSchedule = "base_schedule"
)

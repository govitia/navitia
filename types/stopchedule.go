package types

// A StopSchedule is a timetable of a specific route going through a StopPoint
//
// See http://doc.navitia.io/#stop-schedule
type StopSchedule struct {
	AdditionalInfo string

	// Useful information to display
	Display Display

	// The route of the schedule
	Route Route

	// When does a vehicle stops at the stop point
	DateTimes []PTDateTime

	// The Stop Point in question
	StopPoint StopPoint
}

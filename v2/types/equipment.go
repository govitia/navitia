package types

type Equipment string

const (
	EHasWheelchairAccessibility Equipment = "has_wheelchair_accessibility"
	EHasBikeAccepted            Equipment = "has_bike_accepted"
	EHasAirConditioned          Equipment = "has_air_conditioned"
	EHasVisualAnnouncement      Equipment = "has_visual_announcement"
	EHasAudibleAnnouncement     Equipment = "has_audible_announcement"
	EHasAppropriateEscort       Equipment = "has_appropriate_escort"
	EHasAppropriateSignage      Equipment = "has_appropriate_signage"
	EHasSchoolVehicle           Equipment = "has_school_vehicle"
	EHasWheelchairBoarding      Equipment = "has_wheelchair_boarding"
	EHasShelteredStairs         Equipment = "has_sheltered"
	EHasElevator                Equipment = "has_elevator"
	EHasEscalator               Equipment = "has_escalator"
	EHasBikeDepot               Equipment = "has_bike_depot"
)

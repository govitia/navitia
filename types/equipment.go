package types

// An Equipment codes for specific equipment the public transport object has
type Equipment string

// EquipmentWheelchairAccessibility are known equipments
const (
	EquipmentWheelchairAccessibility Equipment = "has_wheelchair_accessibility"
	EquipmentBikeAccepted                      = "has_bike_accepted"
	EquipmentAirConditioned                    = "has_air_conditioned"
	EquipmentVisualAnnouncement                = "has_visual_announcement"
	EquipmentAudibleAnnouncement               = "has_audible_announcement"
	EquipmentAppropriateEscort                 = "has_appropriate_escort"
	EquipmentAppropriateSignage                = "has_appropriate_signage"
	EquipmentSchoolVehicle                     = "has_school_vehicle"
	EquipmentWheelchairBoarding                = "has_wheelchair_boarding"
	EquipmentSheltered                         = "has_sheltered"
	EquipmentElevator                          = "has_elevator"
	EquipmentEscalator                         = "has_escalator"
	EquipmentBikeDepot                         = "has_bike_depot"
)

// knownEquipments lists all the known equipments
var knownEquipments = []Equipment{
	EquipmentWheelchairAccessibility,
	EquipmentBikeAccepted,
	EquipmentAirConditioned,
	EquipmentVisualAnnouncement,
	EquipmentAudibleAnnouncement,
	EquipmentAppropriateEscort,
	EquipmentAppropriateSignage,
	EquipmentSchoolVehicle,
	EquipmentWheelchairBoarding,
	EquipmentSheltered,
	EquipmentElevator,
	EquipmentEscalator,
	EquipmentBikeDepot,
}

// Known reports whether an equipment is known
func (eq Equipment) Known() bool {
	for _, k := range knownEquipments {
		if eq == k {
			return true
		}
	}
	return false
}

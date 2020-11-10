package types

// An Equipment codes for specific equipment the public transport object has
type Equipment string

// EquipmentWheelchairAccessibility are known equipments
const (
	EquipmentWheelchairAccessibility Equipment = "has_wheelchair_accessibility"
	EquipmentBikeAccepted            Equipment = "has_bike_accepted"
	EquipmentAirConditioned          Equipment = "has_air_conditioned"
	EquipmentVisualAnnouncement      Equipment = "has_visual_announcement"
	EquipmentAudibleAnnouncement     Equipment = "has_audible_announcement"
	EquipmentAppropriateEscort       Equipment = "has_appropriate_escort"
	EquipmentAppropriateSignage      Equipment = "has_appropriate_signage"
	EquipmentSchoolVehicle           Equipment = "has_school_vehicle"
	EquipmentWheelchairBoarding      Equipment = "has_wheelchair_boarding"
	EquipmentSheltered               Equipment = "has_sheltered"
	EquipmentElevator                Equipment = "has_elevator"
	EquipmentEscalator               Equipment = "has_escalator"
	EquipmentBikeDepot               Equipment = "has_bike_depot"
)

// knownEquipments lists all the known equipments
var knownEquipments = [...]Equipment{
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

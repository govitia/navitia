package gonavitia

type Equipment string

// Known equipments
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

// Equipments is a map of equipments to their description
var Equipments = map[Equipment]string{
//TODO
}

// Equipment can be stringified with its description
func (eq Equipment) String() string {
	if desc, ok := Equipments[eq]; ok {
		return desc
	}
	return string(eq)
}

package types

type StopAreaEquipment struct {
	// The equipment details associated with the stop area
	EquipmentDetails EquipmentDetails `json:"equipment_details"`
	// The stop area to which the equipment_details is associated
	StopArea *StopArea `json:"stop_area"`
}

// StopAreaEquipments defines a list of objects that maps equipments details for each stop area.
type StopAreaEquipments []*StopAreaEquipment

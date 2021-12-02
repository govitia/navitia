package types

type EquipmentReport struct {
	// The line to which equipments are associated
	Line *Line `json:"line "`
	// A list of objects that describes equipments for each stop area
	StopAreaEquipments StopAreaEquipments `json:"stop_area_equipments"`
}

// EquipmentReports defines av list of objects that maps each line with its associated stop area equipments.
type EquipmentReports []*EquipmentReport

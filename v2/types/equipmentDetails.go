package types

type EquipmentDetail struct {
	ID                  string                `json:"id"`
	Name                string                `json:"name"`
	CurrentAvailability EquipmentAvailability `json:"current_availability"`
	EmbeddedType        string                `json:"embedded_type"`
}

type EquipmentDetails []*EquipmentDetail

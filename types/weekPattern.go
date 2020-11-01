package types

type WeekPattern struct {
	Monday    bool `json:"monday"`
	Tuesday   bool `json:"tuesday"`
	Friday    bool `json:"friday"`
	Wednesday bool `json:"wednesday"`
	Thursday  bool `json:"thursday"`
	Sunday    bool `json:"sunday"`
	Saturday  bool `json:"saturday"`
}

package types

type Departure struct {
	DisplayInformations Display   `json:"display_informations"`
	StopPoint           StopPoint `json:"stop_point"`
	Route               Route     `json:"route"`
	Links               []Link    `json:"links"`
	StopDateTime
}

type StopDateTime struct {
	Links                 []Link `json:"links"`
	ArrivalDateTime       string `json:"arrival_date_time"`
	DepartureDateTime     string `json:"departure_date_time"`
	BaseArrivalDateTime   string `json:"base_arrival_date_time"`
	BaseDepartureDateTime string `json:"base_departure_date_time"`
	DataFreshness         string `json:"data_freshness"`
}

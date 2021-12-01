package types

// PhysicalMode are fastened and normalized (though the list can -rarely- be extended).
// So it's easier for integrators to map it to a pictogram, but prefer commercial_mode for a text output.
//
// The idea is to use physical modes when building a request to Navitia,
// and commercial modes when building an output to the traveler.
//
// Example: If you want to propose modes filter in your application,
// you should use PhysicalMode rather than CommercialMode.
//
// Here is the valid id list:
//
//	physical_mode:Air
//	physical_mode:Boat
//	physical_mode:Bus
//	physical_mode:BusRapidTransit
//	physical_mode:Coach
//	physical_mode:Ferry
//	physical_mode:Funicular
//	physical_mode:LocalTrain
//	physical_mode:LongDistanceTrain
//	physical_mode:Metro
//	physical_mode:RailShuttle
//	physical_mode:RapidTransit
//	physical_mode:Shuttle
//	physical_mode:SuspendedCableCar
//	physical_mode:Taxi
//	physical_mode:Train
//	physical_mode:Tramway
//
// You can use these ids in the forbidden_uris[] parameter from journeys parameters for example.
type PhysicalMode struct {
	ID              string            `json:"id"`               // Identifier of the physical mode
	Name            string            `json:"name"`             // Name of the physical mode
	CommercialModes []*CommercialMode `json:"commercial_modes"` // Commercial modes of this physical mode
}

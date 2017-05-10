package types

import "testing"

// TestIDCheck checks if ID.Check returns an error when given an empty ID
func TestID_Check(t *testing.T) {
	id := ID("")
	if id.Check() == nil {
		t.Errorf("Received no error even though we expect one")
	}
}

func TestID_Type(t *testing.T) {
	// An empty string means we expect an error
	var samples = map[ID]string{
		"stop_area:OIF:SA:10:1188":         "stop_area",
		"stop_point:OIF:SP:10:10":          "stop_point",
		"line:OIF:066066020:AOIF364":       "line",
		"company:OIF:783":                  "company",
		"network:sncf":                     "network",
		"network:OIF:639":                  "network",
		"route:OIF:002002002:BDE":          "route",
		"route:OIF:002002002:CEN":          "route",
		"commercial_mode:VAL":              "commercial_mode",
		"physical_mode:CheckIn":            "physical_mode",
		"poi:n715165598":                   "poi",
		"poi_type:amenity:bicycle_parking": "poi_type",
		"blabla:invalid":                   "",
		"poi_type-invalid":                 "",
	}

	for id, expected := range samples {
		identified, err := id.Type()
		t.Logf("For sample id \"%s\"\n\twe expect\t \"%s\"\n\twe received\t \"%s\"", id, expected, identified)
		switch {
		case err != nil && expected != "":
			t.Errorf("expected no errors, got one: %v", err)
		case err == nil && expected == "":
			t.Error("expected an error, didn't get one")
		case identified != expected:
			t.Error("ID.Type failed to correctly identify")
		}
	}
}

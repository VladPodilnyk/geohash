package geohash

import "testing"

func TestEncodeCoordinatesWithVariousPrecisions(t *testing.T) {
	lat := 37.421542
	lon := -122.085589

	expectedData := map[int]string{
		1:  "9",
		2:  "9q",
		3:  "9q9",
		4:  "9q9h",
		5:  "9q9hv",
		6:  "9q9hvu",
		7:  "9q9hvu7",
		8:  "9q9hvu7u",
		9:  "9q9hvu7ur",
		10: "9q9hvu7ur9",
		11: "9q9hvu7ur94",
		12: "9q9hvu7ur94c",
	}

	for precision, expected := range expectedData {
		result, _ := Encode(lat, lon, precision)
		if result != expected {
			t.Errorf("Expected %s, but got %s", expected, result)
		}
	}
}

func TestEncodeCoordinatesWithInvalidPrecision(t *testing.T) {
	lat := 37.421542
	lon := -122.085589

	_, err := Encode(lat, lon, 0)
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

func TestEncodeCoordinatesWithInvalidLatitude(t *testing.T) {
	lat := -100.0
	lon := -122.085589

	_, err := Encode(lat, lon, 1)
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

func TestEncodeCoordinatesWithInvalidLongitude(t *testing.T) {
	lat := 37.421542
	lon := -200.0

	_, err := Encode(lat, lon, 1)
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

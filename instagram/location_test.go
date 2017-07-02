package instagram

import "testing"

func Test_getFromLocationPage(t *testing.T) {
	lat := 55.7522
	lng := 37.6156
	slug := "moscow-russia"
	name := "Moscow, Russia"

	jsonBody := []byte("{\"location\":{\"id\":\"17326249\",\"name\":\"Moscow, Russia\",\"has_public_page\":true,\"lat\":55.7522,\"lng\":37.6156,\"slug\":\"moscow-russia\"},\"logging_page_id\":\"locationPage_17326249\"}")

	location, err := getFromLocationPage(jsonBody)
	if err != nil {
		t.Error(err)
	}
	if location.Lat != lat {
		t.Errorf("Wrong location info. Expect slug %f, but get %f.", lat, location.Lat)
	}
	if location.Lng != lng {
		t.Errorf("Wrong location info. Expect slug %f, but get %f.", lng, location.Lng)
	}
	if location.Slug != slug {
		t.Errorf("Wrong location info. Expect slug %s, but get %s.", slug, location.Slug)
	}
	if location.Name != name {
		t.Errorf("Wrong location info. Expect slug %s, but get %s.", name, location.Name)
	}
}

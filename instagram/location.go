package instagram

// A Location describes an Instagram location info.
type Location struct {
	ID         string
	Name       string
	PublicPage bool
	Lat        float64
	Lng        float64
	Slug       string
}

func getFromLocationPage(info map[string]interface{}) (Location, bool) {
	json, ok := info["location"].(map[string]interface{})
	if !ok {
		return Location{}, false
	}

	location := Location{}
	location.ID, _ = json["id"].(string)
	location.Name, _ = json["name"].(string)
	location.PublicPage, _ = json["has_public_page"].(bool)
	location.Lat, _ = json["lat"].(float64)
	location.Lng, _ = json["lng"].(float64)
	location.Slug, _ = json["slug"].(string)

	return location, true
}

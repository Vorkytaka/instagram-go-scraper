package instagram

type Location struct {
	Id              string
	Name            string
	Has_public_page bool
	Lat             float64
	Lng             float64
	Slug            string
}

func GetFromLocationPage(info map[string]interface{}) (Location, bool) {
	json, ok := info["location"].(map[string]interface{})
	if !ok {
		return Location{}, false
	}

	location := Location{}
	location.Id, _ = json["id"].(string)
	location.Name, _ = json["name"].(string)
	location.Has_public_page, _ = json["has_public_page"].(bool)
	location.Lat, _ = json["lat"].(float64)
	location.Lng, _ = json["lng"].(float64)
	location.Slug, _ = json["slug"].(string)

	return location, true
}

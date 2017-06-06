//
// location.go
// Copyright 2017 Konstantin Dovnar
//
// Licensed under the Apache License, Version 2.0 (the "License");
// http://www.apache.org/licenses/LICENSE-2.0
//

package instagram

import "encoding/json"

// A Location describes an Instagram location info.
type Location struct {
	ID         string
	Name       string
	PublicPage bool
	Lat        float64
	Lng        float64
	Slug       string
}

func getFromLocationPage(data []byte) (Location, error) {
	var locationJSON struct {
		Location struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			HasPublicPage bool `json:"has_public_page"`
			Lat           float64 `json:"lat"`
			Lng           float64 `json:"lng"`
			Slug          string `json:"slug"`
		} `json:"location"`
	}

	err := json.Unmarshal(data, &locationJSON)
	if err != nil {
		return Location{}, err
	}

	location := Location{}
	location.ID = locationJSON.Location.ID
	location.Name = locationJSON.Location.Name
	location.PublicPage = locationJSON.Location.HasPublicPage
	location.Lat = locationJSON.Location.Lat
	location.Lng = locationJSON.Location.Lng
	location.Slug = locationJSON.Location.Slug

	return location, nil
}

package instagram_scraper

import (
	"testing"
)

func Test_GetAccoutByUsername(t *testing.T) {
	for _, test_case := range []struct {
		username, full_name, id string
	}{
		{"instagram", "Instagram", "25025320"},
		{"solidlsnake", "Konstantin", "248188406"},
	} {
		account := GetAccoutByUsername(test_case.username)
		if account.Username != test_case.username ||
		   account.Full_name != test_case.full_name ||
		   account.Id != test_case.id {
			t.Error("Unexpected account info.")
		}
	}
}

func Test_GetMediaByUrl(t *testing.T) {
	for _, test_case := range []struct {
		url, code, username, media_type string
	}{
		{
			"https://www.instagram.com/p/ceiqEstT6r/",
			"ceiqEstT6r",
			"solidlsnake",
			"image",
		},
		{
			"https://www.instagram.com/p/12376OtT5o/",
			"12376OtT5o",
			"solidlsnake",
			"video",
		},
	} {
		media := GetMedyaByUrl(test_case.url)
		if media.Code != test_case.code ||
		   media.Owner.Username != test_case.username ||
		   media.Media_type != test_case.media_type {
			t.Error("Unexpected media info.")
		}
	}
}

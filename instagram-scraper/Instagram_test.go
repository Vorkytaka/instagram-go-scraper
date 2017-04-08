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
		media := GetMediaByUrl(test_case.url)
		if media.Code != test_case.code ||
		   media.Owner.Username != test_case.username ||
		   media.Media_type != test_case.media_type {
			t.Error("Unexpected media info.")
		}
	}
}

func Test_GetMediaByCode(t *testing.T) {
	for _, test_case := range []struct {
		code, username, media_type string
	}{
		{
			"ceiqEstT6r",
			"solidlsnake",
			"image",
		},
		{
			"12376OtT5o",
			"solidlsnake",
			"video",
		},
	} {
		media := GetMediaByCode(test_case.code)
		if media.Code != test_case.code ||
		   media.Owner.Username != test_case.username ||
		   media.Media_type != test_case.media_type {
			t.Error("Unexpected media info.")
		}
	}
}

func Test_GetUserMedia_quantity(t *testing.T) {
	count := int(GetAccoutByUsername("solidlsnake").Media_count)

	for _, test_case := range []struct {
		username string
		quantity uint16
		expected int
	}{
		{ "instagram", 10, 10 },
		{ "solidlsnake", 999, count },
	} {
		medias := GetAccountMedia(test_case.username, test_case.quantity)
		if len(medias) != test_case.expected {
			t.Error("Wrong numbers of media.")
		}
	}
}

func Test_GetAllUserMedia_quantity(t *testing.T) {
	for _, test_case := range []struct {
		username string
	}{
		{ "eminem" },
		{ "solidlsnake" },
	} {
		expected := int(GetAccoutByUsername(test_case.username).Media_count)
		medias := GetAllAccountMedia(test_case.username)
		if len(medias) != expected {
			t.Error("Wrong numbers of media.")
		}
	}
}
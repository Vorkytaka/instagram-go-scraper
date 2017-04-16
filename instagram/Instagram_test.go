package instagram

import (
	"testing"
)

func Test_GetAccoutByUsername_exist(t *testing.T) {
	for _, test_case := range []struct {
		username, full_name, id string
	}{
		{"instagram", "Instagram", "25025320"},
		{"solidlsnake", "Konstantin", "248188406"},
	} {
		account, err := GetAccountByUsername(test_case.username)
		if err != nil ||
			account.Username != test_case.username ||
			account.FullName != test_case.full_name ||
			account.ID != test_case.id {
			t.Error("Unexpected account info.")
		}
	}
}

func Test_GetAccoutByUsername_notExist(t *testing.T) {
	for _, test_case := range []struct {
		username string
	}{
		{"fhusdhfjashbfjfghyashf"},
		{"fhusadhfyasifjasduiash"},
	} {
		_, err := GetAccountByUsername(test_case.username)
		if err == nil {
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
		media, err := GetMediaByURL(test_case.url)
		if err != nil ||
			media.Code != test_case.code ||
			media.Owner.Username != test_case.username ||
			media.Type != test_case.media_type {
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
		media, err := GetMediaByCode(test_case.code)
		if err != nil ||
			media.Code != test_case.code ||
			media.Owner.Username != test_case.username ||
			media.Type != test_case.media_type {
			t.Error("Unexpected media info.")
		}
	}
}

func Test_GetMediaByCode_notExist(t *testing.T) {
	for _, test_case := range []struct {
		code string
	}{
		{"aaaaaaaaaa"},
		{"abcdefghij"},
	} {
		_, err := GetMediaByCode(test_case.code)
		if err == nil {
			t.Error("Unexpected media info.")
		}
	}
}

func Test_GetUserMedia_quantity(t *testing.T) {
	account, _ := GetAccountByUsername("solidlsnake")
	count := int(account.MediaCount)

	for _, test_case := range []struct {
		username string
		quantity uint16
		expected int
	}{
		{"instagram", 10, 10},
		{"solidlsnake", 999, count},
	} {
		medias, err := GetAccountMedia(test_case.username, test_case.quantity)
		if err != nil || len(medias) != test_case.expected {
			t.Error("Wrong numbers of media.")
		}
	}
}

func Test_GetAllUserMedia_quantity(t *testing.T) {
	for _, test_case := range []struct {
		username string
	}{
		{"eminem"},
		{"solidlsnake"},
	} {
		account, _ := GetAccountByUsername(test_case.username)
		expected := int(account.MediaCount)
		medias, err := GetAllAccountMedia(test_case.username)
		if err != nil || len(medias) != expected {
			t.Error("Wrong numbers of media.")
		}
	}
}

func Test_GetLocationMedia_quantity(t *testing.T) {
	for _, test_case := range []struct {
		location_id string
		quantity    uint16
	}{
		{"17326249", 10},
		{"17326249", 25},
	} {
		medias, err := GetLocationMedia(test_case.location_id, test_case.quantity)
		if err != nil || len(medias) != int(test_case.quantity) {
			t.Error("Wrong numbers of media.")
		}
	}
}

func Test_GetLocationTopMedia_quantity(t *testing.T) {
	for _, test_case := range []struct {
		location_id string
	}{
		{"17326249"},
	} {
		medias, err := GetLocationTopMedia(test_case.location_id)
		if err != nil || len(medias) != 9 {
			t.Error("Wrong numbers of media.")
		}
	}
}

func Test_GetLocationById(t *testing.T) {
	for _, test_case := range []struct {
		location_id, slug string
	}{
		{"17326249", "moscow-russia"},
		{"212988663", "new-york-new-york"},
	} {
		location, err := GetLocationByID(test_case.location_id)
		if err != nil ||
			location.Slug != test_case.slug {
			t.Error("Wrong location info")
		}
	}
}

func Test_GetTagMedia_quantity(t *testing.T) {
	for _, test_case := range []struct {
		tag      string
		quantity uint16
	}{
		{"lol", 10},
		{"haha", 25},
	} {
		medias, err := GetTagMedia(test_case.tag, test_case.quantity)
		if err != nil || len(medias) != int(test_case.quantity) {
			t.Error("Wrong numbers of media.")
		}
	}
}

func Test_GetTagTopMedia_quantity(t *testing.T) {
	for _, test_case := range []struct {
		tag string
	}{
		{"lol"},
	} {
		medias, err := GetTagTopMedia(test_case.tag)
		if err != nil || len(medias) != 9 {
			t.Error("Wrong numbers of media.")
		}
	}
}

func Test_SearchForUsers(t *testing.T) {
	for _, test_case := range []struct {
		username, id string
	}{
		{"solidlsnake", "248188406" },
	} {
		accounts, err := SearchForUsers(test_case.username)
		if err != nil ||
			accounts[0].ID != test_case.id {
			t.Error("Wrong account information.")
		}
	}
}

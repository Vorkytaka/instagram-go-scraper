package instagram

import (
	"testing"
)

func Test_GetAccoutByUsername_exist(t *testing.T) {
	for _, testCase := range []struct {
		username, fullName, id string
	}{
		{"instagram", "Instagram", "25025320"},
		{"solidlsnake", "Konstantin", "248188406"},
	} {
		account, err := GetAccountByUsername(testCase.username)
		if err != nil ||
			account.Username != testCase.username ||
			account.FullName != testCase.fullName ||
			account.ID != testCase.id {
			t.Errorf(
				"Unexpected account info. Expect full name %s, but get %s. Expect id %s, but get %s",
				testCase.fullName,
				testCase.id,
				account.FullName,
				account.ID,
			)
		}
	}
}

func Test_GetAccoutByUsername_notExist(t *testing.T) {
	for _, testCase := range []struct {
		username string
	}{
		{"fhusdhfjashbfjfghyashf"},
		{"fhusadhfyasifjasduiash"},
	} {
		_, err := GetAccountByUsername(testCase.username)
		if err == nil {
			t.Error("Unexpected account info, account must be not exist.")
		}
	}
}

func Test_GetMediaByUrl(t *testing.T) {
	for _, testCase := range []struct {
		url, code, username, mediaType string
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
		media, err := GetMediaByURL(testCase.url)
		if err != nil ||
			media.Code != testCase.code ||
			media.Owner.Username != testCase.username ||
			media.Type != testCase.mediaType {
			t.Error("Unexpected media info.")
		}
	}
}

func Test_GetMediaByCode(t *testing.T) {
	for _, testCase := range []struct {
		code, username, mediaType string
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
		media, err := GetMediaByCode(testCase.code)
		if err != nil ||
			media.Code != testCase.code ||
			media.Owner.Username != testCase.username ||
			media.Type != testCase.mediaType {
			t.Error("Unexpected media info.")
		}
	}
}

func Test_GetMediaByCode_notExist(t *testing.T) {
	for _, testCase := range []struct {
		code string
	}{
		{"aaaaaaaaaa"},
		{"abcdefghij"},
	} {
		_, err := GetMediaByCode(testCase.code)
		if err == nil {
			t.Error("Unexpected media info.")
		}
	}
}

func Test_GetUserMedia_quantity(t *testing.T) {
	account, _ := GetAccountByUsername("solidlsnake")
	count := int(account.MediaCount)

	for _, testCase := range []struct {
		username string
		quantity uint16
		expected int
	}{
		{"instagram", 10, 10},
		{"solidlsnake", 999, count},
	} {
		medias, err := GetAccountMedia(testCase.username, testCase.quantity)
		if err != nil || len(medias) != testCase.expected {
			t.Error("Wrong numbers of media.")
		}
	}
}

func Test_GetAllUserMedia_quantity(t *testing.T) {
	for _, testCase := range []struct {
		username string
	}{
		{"eminem"},
		{"solidlsnake"},
	} {
		account, _ := GetAccountByUsername(testCase.username)
		expected := int(account.MediaCount)
		medias, err := GetAllAccountMedia(testCase.username)
		if err != nil || len(medias) != expected {
			t.Error("Wrong numbers of media.")
		}
	}
}

func Test_GetLocationMedia_quantity(t *testing.T) {
	for _, testCase := range []struct {
		locationID string
		quantity   uint16
	}{
		{"17326249", 10},
		{"17326249", 25},
	} {
		medias, err := GetLocationMedia(testCase.locationID, testCase.quantity)
		if err != nil || len(medias) != int(testCase.quantity) {
			t.Errorf("Wrong numbers of media. Expect %d, get %d.", testCase.quantity, len(medias))
		}
	}
}

func Test_GetLocationTopMedia_quantity(t *testing.T) {
	for _, testCase := range []struct {
		locationID string
	}{
		{"17326249"},
	} {
		medias, err := GetLocationTopMedia(testCase.locationID)
		if err != nil || len(medias) != 9 {
			t.Error("Wrong numbers of media.")
		}
	}
}

func Test_GetLocationById(t *testing.T) {
	for _, testCase := range []struct {
		locationID, slug string
	}{
		{"17326249", "moscow-russia"},
		{"212988663", "new-york-new-york"},
	} {
		location, err := GetLocationByID(testCase.locationID)
		if err != nil ||
			location.Slug != testCase.slug {
			t.Errorf("Wrong location info. Expect slug %s, but get %s.", testCase.slug, location.Slug )
		}
	}
}

func Test_GetTagMedia_quantity(t *testing.T) {
	for _, testCase := range []struct {
		tag      string
		quantity uint16
	}{
		{"lol", 10},
		{"haha", 25},
	} {
		medias, err := GetTagMedia(testCase.tag, testCase.quantity)
		if err != nil || len(medias) != int(testCase.quantity) {
			t.Error("Wrong numbers of media.")
		}
	}
}

func Test_GetTagTopMedia_quantity(t *testing.T) {
	for _, testCase := range []struct {
		tag string
	}{
		{"lol"},
	} {
		medias, err := GetTagTopMedia(testCase.tag)
		if err != nil || len(medias) != 9 {
			t.Error("Wrong numbers of media.")
		}
	}
}

func Test_SearchForUsers(t *testing.T) {
	for _, testCase := range []struct {
		username, id string
	}{
		{"solidlsnake", "248188406" },
	} {
		accounts, err := SearchForUsers(testCase.username)
		if err != nil ||
			accounts[0].ID != testCase.id {
			t.Error("Wrong account information.")
		}
	}
}

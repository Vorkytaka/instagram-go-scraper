package instagram

import (
	"testing"
)

func Test_GetAccountByUsername(t *testing.T) {
	for _, testCase := range []struct {
		username string
	}{
		{"instagram" },
		{"vorkytaka" },
	} {
		_, err := GetAccountByUsername(testCase.username)
		if err != nil {
			t.Error(err)
		}
	}
}

func Test_GetAccountByUsername_notExist(t *testing.T) {
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

func Test_GetMediaByCode(t *testing.T) {
	for _, testCase := range []struct {
		code string
	}{
		{"ceiqEstT6r" },
		{"12376OtT5o" },
		{"BUxZSI3B5yq" },
	} {
		_, err := GetMediaByCode(testCase.code)
		if err != nil {
			t.Error(err)
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
	account, _ := GetAccountByUsername("vorkytaka")
	count := int(account.MediaCount)

	for _, testCase := range []struct {
		username string
		quantity uint16
		expected int
	}{
		{"instagram", 10, 10},
		{"vorkytaka", 999, count},
	} {
		medias, err := GetAccountMedia(testCase.username, testCase.quantity)
		if err != nil {
			t.Error(err)
		}

		if len(medias) != testCase.expected {
			t.Error("Wrong numbers of media.")
		}
	}
}

func Test_GetAllUserMedia_quantity(t *testing.T) {
	for _, testCase := range []struct {
		username string
	}{
		{"drdre"},
		{"vorkytaka"},
	} {
		account, _ := GetAccountByUsername(testCase.username)
		expected := int(account.MediaCount)
		medias, err := GetAllAccountMedia(testCase.username)
		if err != nil {
			t.Error(err)
		}
		if len(medias) != expected {
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
		if err != nil {
			t.Error(err)
		}
		if len(medias) != int(testCase.quantity) {
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
		if err != nil {
			t.Error(err)
		}
		if len(medias) != 9 {
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
		_, err := GetLocationByID(testCase.locationID)
		if err != nil {
			t.Error(err)
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
		if err != nil {
			t.Error(err)
		}
		if len(medias) != int(testCase.quantity) {
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
		if err != nil {
			t.Error(err)
		}
		if len(medias) != 9 {
			t.Error("Wrong numbers of media.")
		}
	}
}

func Test_SearchForUsers(t *testing.T) {
	for _, testCase := range []struct {
		username string
	}{
		{"vorkytaka"},
	} {
		_, err := SearchForUsers(testCase.username)
		if err != nil {
			t.Error(err)
		}
	}
}

func Test_MediaUpdate(t *testing.T) {
	for _, testCase := range []struct {
		code string
	}{
		{"BNhikAbg5Ph" },
		{"BNUAZOPgdsH" },
	} {
		media := Media{}
		media.Code = testCase.code
		err := media.Update()
		if err != nil {
			t.Error(err)
		}
		if len(media.Caption) == 0 ||
			media.Date == 0 || len(media.ID) == 0 || media.LikesCount == 0 ||
			len(media.Type) == 0 || len(media.MediaURL) == 0 {
			t.Error("Media wasn't update")
		}
	}
}

func Test_AccountUpdate(t *testing.T) {
	for _, testCase := range []struct {
		username, fullname string
	}{
		{"vorkytaka", "Konstantin"},
	} {
		account := Account{}
		account.Username = testCase.username
		err := account.Update()
		if err != nil {
			t.Error(err)
		}
		if account.FullName != testCase.fullname {
			t.Error("Account didn't update")
		}
	}
}

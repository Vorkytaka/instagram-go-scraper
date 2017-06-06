package instagram

import (
	"testing"
)

func Test_GetAccoutByUsername(t *testing.T) {
	for _, testCase := range []struct {
		username, fullname, id, biography, profilePicURL, profilePicURLhd string
		verified                                                          bool
	}{
		{
			"instagram",
			"Instagram",
			"25025320",
			"Discovering — and telling — stories from around the world. Curated by Instagram’s community team.",
			"https://scontent-arn2-1.cdninstagram.com/t51.2885-19/s150x150/14719833_310540259320655_1605122788543168512_a.jpg",
			"https://scontent-arn2-1.cdninstagram.com/t51.2885-19/s320x320/14719833_310540259320655_1605122788543168512_a.jpg",
			true,
		},
		{
			"solidlsnake",
			"Konstantin",
			"248188406",
			"",
			"https://scontent-arn2-1.cdninstagram.com/t51.2885-19/s150x150/17125816_320904131658190_1521093063361953792_a.jpg",
			"https://scontent-arn2-1.cdninstagram.com/t51.2885-19/s320x320/17125816_320904131658190_1521093063361953792_a.jpg",
			false,
		},
	} {
		account, err := GetAccountByUsername(testCase.username)
		if err != nil {
			t.Error(err)
		}
		if account.Username != testCase.username {
			t.Errorf("Account username is incorrect.\nExpect %s, get %s.", account.Username, testCase.username)
		}
		if account.FullName != testCase.fullname {
			t.Errorf("Account fullname is incorrect.\nExpect %s, get %s.", account.FullName, testCase.fullname)
		}
		if account.ID != testCase.id {
			t.Errorf("Account id is incorrect.\nExpect %s, get %s.", account.ID, testCase.id)
		}
		if account.Biography != testCase.biography {
			t.Errorf("Account biography is incorrect.\nExpect %s, get %s.", account.Biography, testCase.biography)
		}
		if account.Verified != testCase.verified {
			t.Errorf("Account verified field is incorrect.\nExpect %t, get %t.", account.ProfilePicURLhd, testCase.profilePicURLhd)
		}
		if account.MediaCount == 0 {
			t.Error("Account has empty media count.")
		}
		if account.Followers == 0 {
			t.Error("Account has empty followers count.")
		}
		if account.Follows == 0 {
			t.Error("Account has empty following count.")
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

func Test_GetMediaByCode(t *testing.T) {
	for _, testCase := range []struct {
		caption, code, id, mediaType, mediaURL                    string
		ownerID, ownerProfilePicURL, ownerUsername, ownerFullName string
	}{
		{
			"Дружище, есть че?",
			"ceiqEstT6r",
			"512999832411258539",
			"image",
			"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/e15/11325321_112286479105192_62945882_n.jpg",
			"248188406",
			"https://scontent-arn2-1.cdninstagram.com/t51.2885-19/s150x150/17125816_320904131658190_1521093063361953792_a.jpg",
			"solidlsnake",
			"Konstantin",
		},
		{
			"Зов джунглей",
			"12376OtT5o",
			"970208779275943528",
			"video",
			"https://scontent-arn2-1.cdninstagram.com/t50.2886-16/11175992_1616299585273258_350800542_n.mp4",
			"248188406",
			"https://scontent-arn2-1.cdninstagram.com/t51.2885-19/s150x150/17125816_320904131658190_1521093063361953792_a.jpg",
			"solidlsnake",
			"Konstantin",
		},
	} {
		media, err := GetMediaByCode(testCase.code)
		if err != nil {
			t.Error(err)
		}
		if media.Caption != testCase.caption {
			t.Errorf("Media caption is incorrect.\nExpect %s, get %s.", media.Caption, testCase.caption)
		}
		if media.Code != testCase.code || media.MediaList[0].Code != testCase.code {
			t.Errorf("Media code is incorrect.\nExpect %s, get %s.", media.Code, testCase.code)
		}
		if media.ID != testCase.id {
			t.Errorf("Media id is incorrect.\nExpect %s, get %s.", media.ID, testCase.id)
		}
		if media.Type != testCase.mediaType || media.MediaList[0].Type != testCase.mediaType {
			t.Errorf("Media type is incorrect.\nExpect %s, get %s.", media.Type, testCase.mediaType)
		}
		if media.Owner.ID != testCase.ownerID {
			t.Errorf("Media's owner ID is incorrect.\nExpect %s, get %s.", media.Owner.ID, testCase.ownerID)
		}
		if media.Owner.Username != testCase.ownerUsername {
			t.Errorf("Media's owner username is incorrect.\nExpect %s, get %s.", media.Owner.Username, testCase.ownerUsername)
		}
		if media.Owner.FullName != testCase.ownerFullName {
			t.Errorf("Media's owner fullname is incorrect.\nExpect %s, get %s.", media.Owner.FullName, testCase.ownerFullName)
		}
		if media.LikesCount == 0 {
			t.Error("Media has empty likes count.")
		}
	}
}

func Test_GetMediaByCode_slider(t *testing.T) {
	for _, testCase := range []struct {
		code  string
		count int
	}{
		{"BUxZSI3B5yq", 2 },
	} {
		media, err := GetMediaByCode(testCase.code)
		if err != nil {
			t.Error(err)
		}
		if media.Code != testCase.code {
			t.Errorf("Media code is incorrect.\nExpect %s, get %s.", media.Code, testCase.code)
		}
		if len(media.MediaList) != testCase.count {
			t.Errorf("Media list got incorrect size.\nExpect %d, get %d.", testCase.count, len(media.MediaList))
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
		{"solidlsnake"},
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

func Test_getFromAccountMediaList(t *testing.T) {
	media, err := GetAllAccountMedia("solidlsnake")
	lastPos := len(media) - 1
	if err != nil {
		t.Error(err)
	}
	if media[lastPos].Caption != "Всем добра и расслабона." {
		t.Errorf("Media caption is incorrect.")
	}
	if media[lastPos].Code != "YgBPkNtTz_" {
		t.Errorf("Media code is incorrect.")
	}
	if media[lastPos].ID != "441358231205657855_248188406" {
		t.Errorf("Media id is incorrect.")
	}
	if media[lastPos].Type != "image" {
		t.Errorf("Media type is incorrect.")
	}
	if media[lastPos].Owner.Username != "solidlsnake" {
		t.Errorf("Media's owner username is incorrect.")
	}
	if media[lastPos].Owner.FullName != "Konstantin" {
		t.Errorf("Media's owner fullname is incorrect.")
	}
	if media[lastPos].LikesCount == 0 {
		t.Error("Media has empty likes count.")
	}
	if media[lastPos].CommentsCount == 0 {
		t.Error("Media has empty comments count.")
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
		location, err := GetLocationByID(testCase.locationID)
		if err != nil {
			t.Error(err)
		}
		if location.Slug != testCase.slug {
			t.Errorf("Wrong location info. Expect slug %s, but get %s.", testCase.slug, location.Slug)
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
		username, id, fullname string
	}{
		{
			"solidlsnake",
			"248188406",
			"Konstantin",
		},
	} {
		accounts, err := SearchForUsers(testCase.username)
		if err != nil {
			t.Error(err)
		}
		if accounts[0].ID != testCase.id {
			t.Error("Incorrect account ID.")
		}
		if accounts[0].FullName != testCase.fullname {
			t.Error("Incorrect account fullname.")
		}
		if accounts[0].Private {
			t.Error("Incorrect account private field.")
		}
		if accounts[0].Verified {
			t.Error("Incorrect account verified field.")
		}
		if accounts[0].Followers == 0 {
			t.Error("Incorrect account followers count.")
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
		{"solidlsnake", "Konstantin"},
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

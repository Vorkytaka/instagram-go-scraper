package instagram_scraper

import (
	"testing"
)

func Test_GetAccoutByUsername(t *testing.T) {
	account := GetAccoutByUsername("instagram")
	if account.Username != "instagram" || account.Full_name != "Instagram" {
		t.Error("Unexpected account info.")
	}
}

func Test_GetMediaByUrl(t *testing.T) {
	media := GetMedyaByUrl("https://www.instagram.com/p/ceiqEstT6r/")
	if media.Code != "ceiqEstT6r" || media.Owner.Username != "solidlsnake" {
		t.Error("Unexpected media info.")
	}
}

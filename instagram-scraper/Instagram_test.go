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

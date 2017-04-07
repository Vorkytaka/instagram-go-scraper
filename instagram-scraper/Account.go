package instagram_scraper

type Account struct {
	Biography          string
	Connected_fb_page  bool
	External_url       string
	Followers          float64
	Follows            float64
	Full_name          string
	Id                 uint64
	Is_private         bool
	Is_verified        bool
	Media_count        float64
	Profile_pic_url    string
	Profile_pic_url_hd string
	Username           string
}

func GetFromAccountPage(info map[string]interface{}) (account Account) {
	user := info["user"].(map[string]interface{})

	account.Biography, _ = user["biography"].(string)
	account.Connected_fb_page, _ = user["connected_fb_page"].(bool)
	account.External_url, _ = user["external_url"].(string)
	account.Full_name, _ = user["full_name"].(string)
	account.Id, _ = user["id"].(uint64)
	account.Is_private, _ = user["is_private"].(bool)
	account.Is_verified, _ = user["is_verified"].(bool)
	account.Profile_pic_url, _ = user["profile_pic_url"].(string)
	account.Profile_pic_url_hd, _ = user["profile_pic_url_hd"].(string)
	account.Username, _ = user["username"].(string)

	followed_by := user["followed_by"].(map[string]interface{})
	account.Followers, _ = followed_by["count"].(float64)

	follows := user["follows"].(map[string]interface{})
	account.Follows, _ = follows["count"].(float64)

	media := user["media"].(map[string]interface{})
	account.Media_count, _ = media["count"].(float64)

	return account
}

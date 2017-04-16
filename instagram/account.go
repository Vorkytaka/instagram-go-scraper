package instagram

// An Account describes an Instagram account info.
type Account struct {
	Biography         string
	ConnectedFacebook bool
	ExternalURL       string
	Followers         uint32
	Follows           uint32
	FullName          string
	ID                string
	Private           bool
	verified          bool
	MediaCount        uint32
	ProfilePicURL     string
	ProfilePicURLhd   string
	Username          string
}

func getFromAccountPage(info map[string]interface{}) (Account, bool) {
	user, ok := info["user"].(map[string]interface{})
	if !ok {
		return Account{}, false
	}

	account := Account{}
	account.Biography, _ = user["biography"].(string)
	account.ConnectedFacebook, _ = user["connected_fb_page"].(bool)
	account.ExternalURL, _ = user["external_url"].(string)
	account.FullName, _ = user["full_name"].(string)
	account.ID, _ = user["id"].(string)
	account.Private, _ = user["is_private"].(bool)
	account.verified, _ = user["is_verified"].(bool)
	account.ProfilePicURL, _ = user["profile_pic_url"].(string)
	account.ProfilePicURLhd, _ = user["profile_pic_url_hd"].(string)
	account.Username, _ = user["username"].(string)

	var fnum float64

	followedBy := user["followed_by"].(map[string]interface{})
	fnum, _ = followedBy["count"].(float64)
	account.Followers = uint32(fnum)

	follows := user["follows"].(map[string]interface{})
	fnum, _ = follows["count"].(float64)
	account.Follows = uint32(fnum)

	media := user["media"].(map[string]interface{})
	fnum, _ = media["count"].(float64)
	account.MediaCount = uint32(fnum)

	return account, true
}

func getFromSearchPage(info map[string]interface{}) (Account, bool) {
	user, ok := info["user"].(map[string]interface{})
	if !ok {
		return Account{}, false
	}

	account := Account{}
	account.ID, _ = user["pk"].(string)
	account.Username, _ = user["username"].(string)
	account.FullName, _ = user["full_name"].(string)
	account.Private, _ = user["is_private"].(bool)
	account.verified, _ = user["is_verified"].(bool)
	account.ProfilePicURL, _ = user["profile_pic_url"].(string)

	followers, ok := user["follower_count"].(float64)
	if ok {
		account.Followers = uint32(followers)
	}

	return account, true
}

package instagram_scraper

type Media struct {
	Caption        string
	Code           string
	Comments_count uint32
	Date           uint64
	Id             string
	Is_ad          bool
	Likes_bount    uint32
	Media_url      string
	Media_type     string
	Owner          Account
}

func GetFromMediaPage(info map[string]interface{}) (media Media) {
	media_info := info["media"].(map[string]interface{})

	media.Caption, _ = media_info["caption"].(string)
	media.Code, _ = media_info["code"].(string)
	media.Id = media_info["id"].(string)
	media.Is_ad = media_info["is_ad"].(bool)

	var fnum float64

	comments, _ := media_info["comments"].(map[string]interface{})
	fnum, _ = comments["count"].(float64)
	media.Comments_count = uint32(fnum)

	fnum, _ = media_info["date"].(float64)
	media.Date = uint64(fnum)

	likes, _ := media_info["likes"].(map[string]interface{})
	fnum = likes["count"].(float64)
	media.Likes_bount = uint32(fnum)

	if media_info["is_video"].(bool) {
		media.Media_type = "video"
		media.Media_url = media_info["video_url"].(string)
	} else {
		media.Media_type = "image"
		media.Media_url = media_info["display_src"].(string)
	}

	owner, _ := media_info["owner"].(map[string]interface{})
	media.Owner.Id, _ = owner["id"].(string)
	media.Owner.Profile_pic_url, _ = owner["profile_pic_url"].(string)
	media.Owner.Username, _ = owner["username"].(string)
	media.Owner.Full_name, _ = owner["full_name"].(string)
	media.Owner.Is_private, _ = owner["is_private"].(bool)

	return
}

func GetFromAccountMediaList(info map[string]interface{}) (media Media) {
	media.Id, _ = info["id"].(string)
	return
}

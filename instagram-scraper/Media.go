package instagram_scraper

type Media struct {
	Caption        string
	Code           string
	Comments_count float64
	Date           uint64
	Display_src    string
	Id             string
	Is_ad          bool
	Likes_bount    float64
	Owner          Account
	Media_type     string
}

func GetFromMediaPage(info map[string]interface{}) (media Media) {
	media_info := info["media"].(map[string]interface{})

	media.Caption, _ = media_info["caption"].(string)
	media.Code, _ = media_info["code"].(string)
	media.Display_src = media_info["display_src"].(string)
	media.Id = media_info["id"].(string)
	media.Is_ad = media_info["is_ad"].(bool)

	comments, _ := media_info["comments"].(map[string]interface{})
	media.Comments_count, _ = comments["count"].(float64)

	fdate, _ := media_info["date"].(float64)
	media.Date = uint64(fdate)

	likes, _ := media_info["likes"].(map[string]interface{})
	media.Likes_bount = likes["count"].(float64)

	if media_info["is_video"].(bool) {
		media.Media_type = "video"
	} else {
		media.Media_type = "image"
	}

	owner, _ := media_info["owner"].(map[string]interface{})
	media.Owner.Id, _ = owner["id"].(uint64)
	media.Owner.Profile_pic_url, _ = owner["profile_pic_url"].(string)
	media.Owner.Username, _ = owner["username"].(string)
	media.Owner.Full_name, _ = owner["full_name"].(string)
	media.Owner.Is_private, _ = owner["is_private"].(bool)

	return
}

package instagram_scraper

import (
	"strconv"
)

const TYPE_IMAGE = "image"
const TYPE_VIDEO = "video"

type Media struct {
	Caption        string
	Code           string
	Comments_count uint32
	Date           uint64
	Id             string
	Is_ad          bool
	Likes_count    uint32
	Media_type     string
	Media_url      string
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
	media.Likes_count = uint32(fnum)

	if media_info["is_video"].(bool) {
		media.Media_type = TYPE_VIDEO
		media.Media_url = media_info["video_url"].(string)
	} else {
		media.Media_type = TYPE_IMAGE
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
	media.Code, _ = info["code"].(string)
	media.Id, _ = info["id"].(string)
	media.Media_type, _ = info["type"].(string)

	sdate := info["created_time"].(string)
	media.Date, _ = strconv.ParseUint(sdate, 10, 64)

	caption, ok := info["caption"].(map[string]interface{})
	if ok {
		media.Caption, _ = caption["text"].(string)
	}

	user, ok := info["user"].(map[string]interface{})
	if ok {
		media.Owner.Username, _ = user["username"].(string)
		media.Owner.Full_name, _ = user["full_name"].(string)
		media.Owner.Id, _ = user["id"].(string)
		media.Owner.Profile_pic_url, _ = user["profile_picture"].(string)
	}

	likes, ok := info["likes"].(map[string]interface{})
	if ok {
		fnum, _ := likes["count"].(float64)
		media.Likes_count = uint32(fnum)
	}

	comments, ok := info["comments"].(map[string]interface{})
	if ok {
		fnum, _ := comments["count"].(float64)
		media.Comments_count = uint32(fnum)
	}

	if media.Media_type == TYPE_VIDEO {
		videos, ok := info["videos"].(map[string]interface{})
		if ok {
			standard_resolution, ok := videos["standard_resolution"].(map[string]interface{})
			if ok {
				media.Media_url, _ = standard_resolution["url"].(string)
			}
		}
	} else {
		images, ok := info["images"].(map[string]interface{})
		if ok {
			standard_resolution, ok := images["standard_resolution"].(map[string]interface{})
			if ok {
				media.Media_url, _ = standard_resolution["url"].(string)
			}
		}
	}

	return
}

func GetFromLocationMediaList(info map[string]interface{}) (media Media) {
	media.Id, _ = info["id"].(string)
	media.Code, _ = info["code"].(string)
	media.Media_url, _ = info["thumbnail_src"].(string)
	media.Caption, _ = info["caption"].(string)

	fnum, _ := info["date"].(float64)
	media.Date = uint64(fnum)

	likes, ok := info["likes"].(map[string]interface{})
	if ok {
		fnum, _ := likes["count"].(float64)
		media.Likes_count = uint32(fnum)
	}

	comments, ok := info["comments"].(map[string]interface{})
	if ok {
		fnum, _ := comments["count"].(float64)
		media.Comments_count = uint32(fnum)
	}

	owner, _ := info["comments"].(map[string]interface{})
	media.Owner.Id, _ = owner["id"].(string)

	if info["is_video"].(bool) {
		media.Media_type = TYPE_VIDEO
	} else {
		media.Media_type = TYPE_IMAGE
	}

	return
}

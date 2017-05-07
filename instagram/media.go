//
// media.go
// Copyright 2017 Konstantin Dovnar
//
// Licensed under the Apache License, Version 2.0 (the "License");
// http://www.apache.org/licenses/LICENSE-2.0
//

package instagram

import (
	"strconv"
)

// TypeImage is a string that define image type for media.
const TypeImage = "image"

// TypeVideo is a string that define video type for media.
const TypeVideo = "video"

// A Media describes an Instagram media info.
type Media struct {
	Caption       string
	Code          string
	CommentsCount uint32
	Date          uint64
	ID            string
	AD            bool
	LikesCount    uint32
	Type          string
	MediaURL      string
	Owner         Account
}

func getFromMediaPage(info map[string]interface{}) (Media, bool) {
	mediaInfo, ok := info["graphql"].(map[string]interface{})
	if !ok {
		return Media{}, false
	}

	mediaInfo, ok = mediaInfo["shortcode_media"].(map[string]interface{})
	if !ok {
		return Media{}, false
	}

	media := Media{}
	media.Code, _ = mediaInfo["shortcode"].(string)
	media.ID = mediaInfo["id"].(string)
	media.AD = mediaInfo["is_ad"].(bool)
	media.Caption, _ = mediaInfo["caption"].(string)

	var fnum float64

	comments, _ := mediaInfo["edge_media_to_comment"].(map[string]interface{})
	fnum, _ = comments["count"].(float64)
	media.CommentsCount = uint32(fnum)

	fnum, _ = mediaInfo["taken_at_timestamp"].(float64)
	media.Date = uint64(fnum)

	likes, _ := mediaInfo["edge_media_preview_like"].(map[string]interface{})
	fnum = likes["count"].(float64)
	media.LikesCount = uint32(fnum)

	caption, _ := mediaInfo["edge_media_to_caption"].(map[string]interface{})
	caption2, _ := caption["edges"].([]interface{})
	caption, _ = caption2[0].(map[string]interface{})
	caption, _ = caption["node"].(map[string]interface{})
	media.Caption = caption["text"].(string)

	if mediaInfo["is_video"].(bool) {
		media.Type = TypeVideo
		media.MediaURL = mediaInfo["video_url"].(string)
	} else {
		media.Type = TypeImage
		media.MediaURL = mediaInfo["display_url"].(string)
	}

	owner, _ := mediaInfo["owner"].(map[string]interface{})
	media.Owner.ID, _ = owner["id"].(string)
	media.Owner.ProfilePicURL, _ = owner["profile_pic_url"].(string)
	media.Owner.Username, _ = owner["username"].(string)
	media.Owner.FullName, _ = owner["full_name"].(string)
	media.Owner.Private, _ = owner["is_private"].(bool)

	return media, true
}

func getFromAccountMediaList(info interface{}) (Media, bool) {
	body, ok := info.(map[string]interface{})
	if !ok {
		return Media{}, false
	}

	media := Media{}
	media.Code, _ = body["code"].(string)
	media.ID, _ = body["id"].(string)
	media.Type, _ = body["type"].(string)

	sdate := body["created_time"].(string)
	media.Date, _ = strconv.ParseUint(sdate, 10, 64)

	caption, ok := body["caption"].(map[string]interface{})
	if ok {
		media.Caption, _ = caption["text"].(string)
	}

	user, ok := body["user"].(map[string]interface{})
	if ok {
		media.Owner.Username, _ = user["username"].(string)
		media.Owner.FullName, _ = user["full_name"].(string)
		media.Owner.ID, _ = user["id"].(string)
		media.Owner.ProfilePicURL, _ = user["profile_picture"].(string)
	}

	likes, ok := body["likes"].(map[string]interface{})
	if ok {
		fnum, _ := likes["count"].(float64)
		media.LikesCount = uint32(fnum)
	}

	comments, ok := body["comments"].(map[string]interface{})
	if ok {
		fnum, _ := comments["count"].(float64)
		media.CommentsCount = uint32(fnum)
	}

	if media.Type == TypeVideo {
		videos, ok := body["videos"].(map[string]interface{})
		if ok {
			standardResolution, ok := videos["standard_resolution"].(map[string]interface{})
			if ok {
				media.MediaURL, _ = standardResolution["url"].(string)
			}
		}
	} else {
		images, ok := body["images"].(map[string]interface{})
		if ok {
			standardResolution, ok := images["standard_resolution"].(map[string]interface{})
			if ok {
				media.MediaURL, _ = standardResolution["url"].(string)
			}
		}
	}

	return media, true
}

func getFromSearchMediaList(info interface{}) (Media, bool) {
	body, ok := info.(map[string]interface{})
	if !ok {
		return Media{}, false
	}

	media := Media{}
	media.ID, _ = body["id"].(string)
	media.Code, _ = body["code"].(string)
	media.MediaURL, _ = body["thumbnail_src"].(string)
	media.Caption, _ = body["caption"].(string)

	fnum, _ := body["date"].(float64)
	media.Date = uint64(fnum)

	likes, ok := body["likes"].(map[string]interface{})
	if ok {
		fnum, _ := likes["count"].(float64)
		media.LikesCount = uint32(fnum)
	}

	comments, ok := body["comments"].(map[string]interface{})
	if ok {
		fnum, _ := comments["count"].(float64)
		media.CommentsCount = uint32(fnum)
	}

	owner, _ := body["owner"].(map[string]interface{})
	media.Owner.ID, _ = owner["id"].(string)

	if body["is_video"].(bool) {
		media.Type = TypeVideo
	} else {
		media.Type = TypeImage
	}

	return media, true
}

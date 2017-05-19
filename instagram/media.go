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
	"encoding/json"
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

// Update try to update media data
func (m *Media) Update() {
	media, err := GetMediaByCode(m.Code)
	if err == nil {
		*m = media
	}
}

func getFromMediaPage(data []byte) (Media, bool) {
	var mediaJson struct {
		Graphql struct {
			ShortcodeMedia struct {
				Typename   string `json:"__typename"`
				ID         string `json:"id"`
				Shortcode  string `json:"shortcode"`
				DisplayURL string `json:"display_url"`
				VideoURL   string `json:"video_url"`
				IsVideo    bool `json:"is_video"`
				EdgeMediaToCaption struct {
					Edges []struct {
						Node struct {
							Text string `json:"text"`
						} `json:"node"`
					} `json:"edges"`
				} `json:"edge_media_to_caption"`
				EdgeMediaToComment struct {
					Count int `json:"count"`
				} `json:"edge_media_to_comment"`
				TakenAtTimestamp int `json:"taken_at_timestamp"`
				EdgeMediaPreviewLike struct {
					Count int `json:"count"`
				} `json:"edge_media_preview_like"`
				Owner struct {
					ID            string `json:"id"`
					ProfilePicURL string `json:"profile_pic_url"`
					Username      string `json:"username"`
					FullName      string `json:"full_name"`
					IsPrivate     bool `json:"is_private"`
				} `json:"owner"`
				IsAd bool `json:"is_ad"`
			} `json:"shortcode_media"`
		} `json:"graphql"`
	}

	err := json.Unmarshal(data, &mediaJson)
	if err != nil {
		return Media{}, false
	}

	media := Media{}
	media.Code = mediaJson.Graphql.ShortcodeMedia.Shortcode
	media.ID = mediaJson.Graphql.ShortcodeMedia.ID
	media.AD = mediaJson.Graphql.ShortcodeMedia.IsAd
	media.Date = uint64(mediaJson.Graphql.ShortcodeMedia.TakenAtTimestamp)
	media.CommentsCount = uint32(mediaJson.Graphql.ShortcodeMedia.EdgeMediaToComment.Count)
	media.LikesCount = uint32(mediaJson.Graphql.ShortcodeMedia.EdgeMediaPreviewLike.Count)
	media.Caption = mediaJson.Graphql.ShortcodeMedia.EdgeMediaToCaption.Edges[0].Node.Text

	if mediaJson.Graphql.ShortcodeMedia.IsVideo {
		media.Type = TypeVideo
		media.MediaURL = mediaJson.Graphql.ShortcodeMedia.VideoURL
	} else {
		media.Type = TypeImage
		media.MediaURL = mediaJson.Graphql.ShortcodeMedia.DisplayURL
	}

	media.Owner.ID = mediaJson.Graphql.ShortcodeMedia.Owner.ID
	media.Owner.ProfilePicURL = mediaJson.Graphql.ShortcodeMedia.Owner.ProfilePicURL
	media.Owner.Username = mediaJson.Graphql.ShortcodeMedia.Owner.Username
	media.Owner.FullName = mediaJson.Graphql.ShortcodeMedia.Owner.FullName
	media.Owner.Private = mediaJson.Graphql.ShortcodeMedia.Owner.IsPrivate

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

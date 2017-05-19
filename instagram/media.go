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

func getFromAccountMediaList(data []byte) (Media, bool) {
	var mediaJson struct {
		ID   string `json:"id"`
		Code string `json:"code"`
		User struct {
			ID             string `json:"id"`
			FullName       string `json:"full_name"`
			ProfilePicture string `json:"profile_picture"`
			Username       string `json:"username"`
		} `json:"user"`
		Images struct {
			StandardResolution struct {
				Width  int `json:"width"`
				Height int `json:"height"`
				URL    string `json:"url"`
			} `json:"standard_resolution"`
		} `json:"images"`
		CreatedTime string `json:"created_time"`
		Caption struct {
			Text string `json:"text"`
		} `json:"caption"`
		Likes struct {
			Count float64 `json:"count"`
		} `json:"likes"`
		Comments struct {
			Count int `json:"count"`
		} `json:"comments"`
		Type string `json:"type"`
		Videos struct {
			StandardResolution struct {
				Width  int `json:"width"`
				Height int `json:"height"`
				URL    string `json:"url"`
			} `json:"standard_resolution"`
		} `json:"videos"`
	}

	err := json.Unmarshal(data, &mediaJson)
	if err != nil {
		return Media{}, false
	}

	media := Media{}
	media.Code = mediaJson.Code
	media.ID = mediaJson.ID
	media.Type = mediaJson.Type
	media.Caption = mediaJson.Caption.Text
	media.LikesCount = uint32(mediaJson.Likes.Count)
	media.CommentsCount = uint32(mediaJson.Comments.Count)

	date, err := strconv.ParseUint(mediaJson.CreatedTime, 10, 64)
	if err == nil {
		media.Date = date
	}

	if media.Type == TypeVideo {
		media.MediaURL = mediaJson.Videos.StandardResolution.URL
	} else {
		media.MediaURL = mediaJson.Images.StandardResolution.URL
	}

	media.Owner.Username = mediaJson.User.Username
	media.Owner.FullName = mediaJson.User.FullName
	media.Owner.ID = mediaJson.User.ID
	media.Owner.ProfilePicURL = mediaJson.User.ProfilePicture

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

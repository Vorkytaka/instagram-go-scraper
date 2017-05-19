//
// account.go
// Copyright 2017 Konstantin Dovnar
//
// Licensed under the Apache License, Version 2.0 (the "License");
// http://www.apache.org/licenses/LICENSE-2.0
//

package instagram

import "encoding/json"

// An Account describes an Instagram account info.
type Account struct {
	Biography       string
	ExternalURL     string
	Followers       uint32
	Follows         uint32
	FullName        string
	ID              string
	Private         bool
	verified        bool
	MediaCount      uint32
	ProfilePicURL   string
	ProfilePicURLhd string
	Username        string
}

func getFromAccountPage(data []byte) (Account, bool) {
	var accountJson struct {
		User struct {
			Biography              string `json:"biography"`
			ExternalURL            string `json:"external_url"`
			FollowedBy struct {
				Count int `json:"count"`
			} `json:"followed_by"`
			Follows struct {
				Count int `json:"count"`
			} `json:"follows"`
			FullName           string `json:"full_name"`
			ID                 string `json:"id"`
			IsPrivate          bool `json:"is_private"`
			IsVerified         bool `json:"is_verified"`
			ProfilePicURL      string `json:"profile_pic_url"`
			ProfilePicURLHd    string `json:"profile_pic_url_hd"`
			Username           string `json:"username"`
			Media struct {
				Count int `json:"count"`
			} `json:"media"`
		} `json:"user"`
	}

	err := json.Unmarshal(data, &accountJson)
	if err != nil {
		return Account{}, false
	}

	account := Account{}
	account.Biography = accountJson.User.Biography
	account.ExternalURL = accountJson.User.ExternalURL
	account.FullName = accountJson.User.FullName
	account.ID = accountJson.User.ID
	account.Private = accountJson.User.IsPrivate
	account.verified = accountJson.User.IsVerified
	account.ProfilePicURL = accountJson.User.ProfilePicURL
	account.ProfilePicURLhd = accountJson.User.ProfilePicURLHd
	account.Username = accountJson.User.Username
	account.Followers = uint32(accountJson.User.FollowedBy.Count)
	account.Follows = uint32(accountJson.User.Follows.Count)
	account.MediaCount = uint32(accountJson.User.Media.Count)

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

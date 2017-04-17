//
// instagram.go
// Copyright 2017 Konstantin Dovnar
//
// Licensed under the Apache License, Version 2.0 (the "License");
// http://www.apache.org/licenses/LICENSE-2.0
//

// Package instagram helps you with requesting to Instagram without a key.
package instagram

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetAccountByUsername try to find account by username.
func GetAccountByUsername(username string) (Account, error) {
	url := fmt.Sprintf(accountInfoURL, username)
	info, err := getJSONFromURL(url)
	if err != nil {
		return Account{}, err
	}
	account, ok := getFromAccountPage(info)
	if !ok {
		return account, errors.New("Can't parse account")
	}
	return account, nil
}

// GetMediaByURL try to find media by url.
// URL should be like https://www.instagram.com/p/12376OtT5o/
func GetMediaByURL(url string) (Media, error) {
	code := strings.Split(url, "/")[4]
	return GetMediaByCode(code)
}

// GetMediaByCode try to find media by code.
// Code can be find in URL to media, after p/.
// If URL to media is https://www.instagram.com/p/12376OtT5o/,
// then code of the media is 12376OtT5o.
func GetMediaByCode(code string) (Media, error) {
	url := fmt.Sprintf(mediaInfoURL, code)
	info, err := getJSONFromURL(url)
	if err != nil {
		return Media{}, err
	}
	media, ok := getFromMediaPage(info)
	if !ok {
		return media, errors.New("Can't parse media")
	}
	return media, nil
}

// GetAccountMedia try to get slice of user's media.
// Limit set how much media you need.
func GetAccountMedia(username string, limit uint16) ([]Media, error) {
	var count uint16
	maxID := ""
	available := true
	medias := []Media{}
	for available && count < limit {
		url := fmt.Sprintf(accountMediaURL, username, maxID)
		jsonBody, err := getJSONFromURL(url)
		if err != nil {
			return nil, err
		}
		available, _ = jsonBody["more_available"].(bool)

		items, _ := jsonBody["items"].([]interface{})
		for _, item := range items {
			if count >= limit {
				return medias, nil
			}
			count++
			media, ok := getFromAccountMediaList(item)
			if ok {
				medias = append(medias, media)
				maxID = media.ID
			}
		}
	}
	return medias, nil
}

// GetAllAccountMedia try to get slice of all user's media.
// It's function the same as GetAccountMedia,
// except limit = count of user's media.
func GetAllAccountMedia(username string) ([]Media, error) {
	account, err := GetAccountByUsername(username)
	if err != nil {
		return nil, err
	}
	count := uint16(account.MediaCount)
	medias, err := GetAccountMedia(username, count)
	if err != nil {
		return nil, err
	}
	return medias, nil
}

// GetLocationMedia try to get slice of last location's media.
// The id is a facebook location id.
// The limit set how much media you need.
func GetLocationMedia(id string, limit uint16) ([]Media, error) {
	var count uint16
	maxID := ""
	hasNext := true
	medias := []Media{}
	for hasNext && count < limit {
		url := fmt.Sprintf(locationURL, id, maxID)
		jsonBody, err := getJSONFromURL(url)
		if err != nil {
			return nil, err
		}
		jsonBody, _ = jsonBody["location"].(map[string]interface{})
		jsonBody, _ = jsonBody["media"].(map[string]interface{})

		nodes, _ := jsonBody["nodes"].([]interface{})
		for _, node := range nodes {
			if count >= limit {
				return medias, nil
			}
			count++
			media, ok := getFromSearchMediaList(node)
			if ok {
				medias = append(medias, media)
			}
		}

		jsonBody, _ = jsonBody["page_info"].(map[string]interface{})
		hasNext, _ = jsonBody["has_next_page"].(bool)
		maxID, _ = jsonBody["end_cursor"].(string)
	}
	return medias, nil
}

// GetLocationTopMedia try to get array of top location's media.
// The id is a facebook location id.
// Length of returned array is 9.
func GetLocationTopMedia(id string) ([9]Media, error) {
	url := fmt.Sprintf(locationURL, id, "")
	jsonBody, err := getJSONFromURL(url)
	if err != nil {
		return [9]Media{}, err
	}
	jsonBody, _ = jsonBody["location"].(map[string]interface{})
	jsonBody, _ = jsonBody["top_posts"].(map[string]interface{})

	medias := [9]Media{}
	nodes, _ := jsonBody["nodes"].([]interface{})
	for i, node := range nodes {
		media, ok := getFromSearchMediaList(node)
		if ok {
			medias[i] = media
		}
	}
	return medias, nil
}

// GetLocationByID try to find location info by id.
// The id is a facebook location id.
func GetLocationByID(id string) (Location, error) {
	url := fmt.Sprintf(locationURL, id, "")
	jsonBody, err := getJSONFromURL(url)
	if err != nil {
		return Location{}, err
	}

	location, ok := getFromLocationPage(jsonBody)
	if !ok {
		return Location{}, errors.New("Can't parse location")
	}
	return location, nil
}

// GetTagMedia try to get slice of last tag's media.
// The limit set how much media you need.
func GetTagMedia(tag string, quantity uint16) ([]Media, error) {
	var count uint16
	maxID := ""
	hasNext := true
	medias := []Media{}
	for hasNext && count < quantity {
		url := fmt.Sprintf(tagURL, tag, maxID)
		jsonBody, err := getJSONFromURL(url)
		if err != nil {
			return nil, err
		}
		jsonBody, _ = jsonBody["tag"].(map[string]interface{})
		jsonBody, _ = jsonBody["media"].(map[string]interface{})

		nodes, _ := jsonBody["nodes"].([]interface{})
		for _, node := range nodes {
			if count >= quantity {
				return medias, nil
			}
			count++
			media, ok := getFromSearchMediaList(node)
			if ok {
				medias = append(medias, media)
			}
		}

		jsonBody, _ = jsonBody["page_info"].(map[string]interface{})
		hasNext, _ = jsonBody["has_next_page"].(bool)
		maxID, _ = jsonBody["end_cursor"].(string)
	}
	return medias, nil
}

// GetTagTopMedia try to get array of top tag's media.
// Length of returned array is 9.
func GetTagTopMedia(tag string) ([9]Media, error) {
	url := fmt.Sprintf(tagURL, tag, "")
	jsonBody, err := getJSONFromURL(url)
	if err != nil {
		return [9]Media{}, err
	}
	jsonBody, _ = jsonBody["tag"].(map[string]interface{})
	jsonBody, _ = jsonBody["top_posts"].(map[string]interface{})

	medias := [9]Media{}
	nodes, _ := jsonBody["nodes"].([]interface{})
	for i, node := range nodes {
		media, ok := getFromSearchMediaList(node)
		if ok {
			medias[i] = media
		}
	}
	return medias, nil
}

// SearchForUsers try to find users by given username.
// Return slice of Account with length of 0 or more.
func SearchForUsers(username string) ([]Account, error) {
	url := fmt.Sprintf(searchURL, username)
	jsonBody, err := getJSONFromURL(url)
	if err != nil {
		return nil, err
	}
	accounts := []Account{}
	users, _ := jsonBody["users"].([]interface{})
	for _, user := range users {
		account, ok := getFromSearchPage(user.(map[string]interface{}))
		if ok {
			accounts = append(accounts, account)
		}
	}
	return accounts, nil
}

func getJSONFromURL(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode == 404 {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var jsonBody map[string]interface{}
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return nil, err
	}

	return jsonBody, nil
}

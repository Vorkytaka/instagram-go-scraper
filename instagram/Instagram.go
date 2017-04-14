// A package that helps you with requesting to Instagram without a key.
package instagram

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetAccoutByUsername(username string) (Account, error) {
	url := fmt.Sprintf(ACCOUNT_JSON_INFO, username)
	info, err := getJsonFromUrl(url)
	if err != nil {
		return Account{}, err
	}
	account, ok := GetFromAccountPage(info)
	if !ok {
		return account, errors.New("Can't parse account")
	}
	return account, nil
}

func GetMediaByUrl(url string) (Media, error) {
	code := strings.Split(url, "/")[4]
	return GetMediaByCode(code)
}

func GetMediaByCode(code string) (Media, error) {
	url := fmt.Sprintf(MEDIA_JSON_INFO, code)
	info, err := getJsonFromUrl(url)
	if err != nil {
		return Media{}, err
	}
	media, ok := GetFromMediaPage(info)
	if !ok {
		return media, errors.New("Can't parse media")
	}
	return media, nil
}

func GetAccountMedia(username string, quantity uint16) ([]Media, error) {
	var count uint16 = 0
	max_id := ""
	available := true
	medias := []Media{}
	for available && count < quantity {
		url := fmt.Sprintf(ACCOUNT_MEDIA_JSON, username, max_id)
		json_body, err := getJsonFromUrl(url)
		if err != nil {
			return nil, err
		}
		available, _ = json_body["more_available"].(bool)

		items, _ := json_body["items"].([]interface{})
		for _, item := range items {
			if count >= quantity {
				return medias, nil
			}
			count++
			media, ok := GetFromAccountMediaList(item)
			if ok {
				medias = append(medias, media)
				max_id = media.Id
			}
		}
	}
	return medias, nil
}

func GetAllAccountMedia(username string) ([]Media, error) {
	account, err := GetAccoutByUsername(username)
	if err != nil {
		return nil, err
	}
	count := uint16(account.Media_count)
	medias, err := GetAccountMedia(username, count)
	if err != nil {
		return nil, err
	}
	return medias, nil
}

func GetLocationMedia(location_id string, quantity uint16) ([]Media, error) {
	var count uint16 = 0
	max_id := ""
	has_next := true
	medias := []Media{}
	for has_next && count < quantity {
		url := fmt.Sprintf(LOCATION_JSON, location_id, max_id)
		json_body, err := getJsonFromUrl(url)
		if err != nil {
			return nil, err
		}
		sub_json, _ := json_body["location"].(map[string]interface{})
		sub_json, _ = sub_json["media"].(map[string]interface{})

		nodes, _ := sub_json["nodes"].([]interface{})
		for _, node := range nodes {
			if count >= quantity {
				return medias, nil
			}
			count++
			media, ok := GetFromLocationMediaList(node)
			if ok {
				medias = append(medias, media)
			}
		}

		sub_json, _ = sub_json["page_info"].(map[string]interface{})
		has_next, _ = sub_json["has_next_page"].(bool)
		max_id, _ = sub_json["end_cursor"].(string)
	}
	return medias, nil
}

func GetLocationTopMedia(location_id string) ([9]Media, error) {
	url := fmt.Sprintf(LOCATION_JSON, location_id, "")
	json_body, err := getJsonFromUrl(url)
	if err != nil {
		return [9]Media{}, err
	}
	sub_json, _ := json_body["location"].(map[string]interface{})
	sub_json, _ = sub_json["top_posts"].(map[string]interface{})

	medias := [9]Media{}
	nodes, _ := sub_json["nodes"].([]interface{})
	for i, node := range nodes {
		media, ok := GetFromLocationMediaList(node)
		if ok {
			medias[i] = media
		}
	}

	sub_json, _ = sub_json["page_info"].(map[string]interface{})
	return medias, nil
}

func GetLocationById(location_id string) (Location, error) {
	url := fmt.Sprintf(LOCATION_JSON, location_id, "")
	json_body, err := getJsonFromUrl(url)
	if err != nil {
		return Location{}, err
	}

	location, ok := GetFromLocationPage(json_body)
	if !ok {
		return Location{}, errors.New("Can't parse location")
	}
	return location, nil
}

func GetTagMedia(tag string, quantity uint16) ([]Media, error) {
	var count uint16 = 0
	max_id := ""
	has_next := true
	medias := []Media{}
	for has_next && count < quantity {
		url := fmt.Sprintf(TAG_JSON, tag, max_id)
		json_body, err := getJsonFromUrl(url)
		if err != nil {
			return nil, err
		}
		sub_json, _ := json_body["tag"].(map[string]interface{})
		sub_json, _ = sub_json["media"].(map[string]interface{})

		nodes, _ := sub_json["nodes"].([]interface{})
		for _, node := range nodes {
			if count >= quantity {
				return medias, nil
			}
			count++
			media, ok := GetFromLocationMediaList(node)
			if ok {
				medias = append(medias, media)
			}
		}

		sub_json, _ = sub_json["page_info"].(map[string]interface{})
		has_next, _ = sub_json["has_next_page"].(bool)
		max_id, _ = sub_json["end_cursor"].(string)
	}
	return medias, nil
}

func GetTagTopMedia(tag string) ([9]Media, error) {
	url := fmt.Sprintf(TAG_JSON, tag, "")
	json_body, err := getJsonFromUrl(url)
	if err != nil {
		return [9]Media{}, err
	}
	sub_json, _ := json_body["tag"].(map[string]interface{})
	sub_json, _ = sub_json["top_posts"].(map[string]interface{})

	medias := [9]Media{}
	nodes, _ := sub_json["nodes"].([]interface{})
	for i, node := range nodes {
		media, ok := GetFromLocationMediaList(node)
		if ok {
			medias[i] = media
		}
	}

	sub_json, _ = sub_json["page_info"].(map[string]interface{})
	return medias, nil
}

func getJsonFromUrl(url string) (json_body map[string]interface{}, err error) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode == 404 {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &json_body)
	if err != nil {
		return nil, err
	}

	return
}

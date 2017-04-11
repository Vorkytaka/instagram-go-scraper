package instagram_scraper

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"log"
	"strings"
	"errors"
)

func GetAccoutByUsername(username string) (Account, error) {
	url := fmt.Sprintf(ACCOUNT_JSON_INFO, username)
	info, err := _GetJsonFromUrl(url)
	if err != nil {
		return Account{}, err
	}
	account, ok := GetFromAccountPage(info)
	if !ok {
		return account, errors.New("Can't parse account")
	}
	return account, nil
}

func GetMediaByUrl(url string) (media Media) {
	code := strings.Split(url, "/")[4]
	media = GetMediaByCode(code)
	return
}

func GetMediaByCode(code string) (media Media) {
	url := fmt.Sprintf(MEDIA_JSON_INFO, code)
	info, err := _GetJsonFromUrl(url)
	if err != nil {
		log.Fatal(err)
	}
	media, ok := GetFromMediaPage(info)
	if !ok {

	}
	return
}

func GetAccountMedia(username string, quantity uint16) (medias []Media) {
	var count uint16 = 0
	max_id := ""
	available := true
	for available && count < quantity {
		url := fmt.Sprintf(ACCOUNT_MEDIA_JSON, username, max_id)
		json_body, err := _GetJsonFromUrl(url)
		if err != nil {
			log.Fatal(err)
		}
		available, _ = json_body["more_available"].(bool)

		items, _ := json_body["items"].([]interface{})
		for _, item := range items {
			if count >= quantity {
				return medias
			}
			count++
			media, ok := GetFromAccountMediaList(item)
			if ok {
				medias = append(medias, media)
				max_id = media.Id
			}
		}
	}
	return
}

func GetLocationMedia(location_id string, quantity uint16) (medias []Media) {
	var count uint16 = 0
	max_id := ""
	has_next := true
	for has_next && count < quantity {
		url := fmt.Sprintf(LOCATION_MEDIA_JSON, location_id, max_id)
		json_body, err := _GetJsonFromUrl(url)
		if err != nil {
			log.Fatal(err)
		}
		sub_json, _ := json_body["location"].(map[string]interface{})
		sub_json, _ = sub_json["media"].(map[string]interface{})

		nodes, _ := sub_json["nodes"].([]interface{})
		for _, node := range nodes {
			if count >= quantity {
				return medias
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
	return
}

func GetLocationTopMedia(location_id string) (medias []Media) {
	var count uint16 = 0
	url := fmt.Sprintf(LOCATION_MEDIA_JSON, location_id, "")
	json_body, err := _GetJsonFromUrl(url)
	if err != nil {
		log.Fatal(err)
	}
	sub_json, _ := json_body["location"].(map[string]interface{})
	sub_json, _ = sub_json["top_posts"].(map[string]interface{})

	nodes, _ := sub_json["nodes"].([]interface{})
	for _, node := range nodes {
		count++
		media, ok := GetFromLocationMediaList(node)
		if ok {
			medias = append(medias, media)
		}
	}

	sub_json, _ = sub_json["page_info"].(map[string]interface{})
	return
}

func GetAllAccountMedia(username string) (medias []Media) {
	account, err := GetAccoutByUsername(username)
	if err != nil {
		log.Fatal(err)
	}
	count := uint16(account.Media_count)
	medias = GetAccountMedia(username, count)
	return medias
}

func _GetJsonFromUrl(url string) (json_body map[string]interface{}, err error) {
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

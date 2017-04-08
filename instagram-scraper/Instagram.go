package instagram_scraper

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"log"
	"strings"
)

func GetAccoutByUsername(username string) (account Account) {
	url := fmt.Sprintf(ACCOUNT_JSON_INFO, username)
	info, err := _GetJsonFromUrl(url)
	if err != nil {
		log.Fatal(err)
	}
	account = GetFromAccountPage(info)
	return
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
	media = GetFromMediaPage(info)
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
			info, _ := item.(map[string]interface{})
			media := GetFromAccountMediaList(info)
			medias = append(medias, media)
			max_id = media.Id
		}
	}
	return
}

func GetAllAccountMedia(username string) (medias []Media) {
	count := uint16(GetAccoutByUsername(username).Media_count)
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

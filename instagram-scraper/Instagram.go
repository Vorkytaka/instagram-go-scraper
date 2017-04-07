package instagram_scraper

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"log"
)

func GetAccoutByUsername(username string) (account Account) {
	url := fmt.Sprintf(ACCOUNT_JSON_INFO, username)
	info := _GetJsonFromUrl(url)
	account = GetFromAccountPage(info)
	return account
}

func GetMedyaByUrl(url string) (media Media) {
	url += "?__a=1"
	info := _GetJsonFromUrl(url)
	media = GetFromMediaPage(info)
	return
}

func _GetJsonFromUrl(url string) (json_body map[string]interface{}) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == 404 {
		log.Fatal("Page Not Found, Code 404")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &json_body)
	if err != nil {
		log.Fatal(err)
	}

	return
}

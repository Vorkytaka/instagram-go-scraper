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

	var info map[string]interface{}
	err = json.Unmarshal(body, &info)
	if err != nil {
		log.Fatal(err)
	}

	account = GetFromAccountPage(info)
	return account
}

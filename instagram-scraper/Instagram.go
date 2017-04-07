package instagram_scraper

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func GetAccoutByUsername(username string) (account Account) {
	url := fmt.Sprintf(ACCOUNT_JSON_INFO, username)
	resp, err := http.Get(url)
	if err != nil {
		
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}

	var info map[string]interface{}
	err = json.Unmarshal(body, &info)
	if err != nil {

	}

	account = GetFromAccountPage(info)
	return account
}

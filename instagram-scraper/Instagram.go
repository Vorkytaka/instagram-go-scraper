package instagram_scraper

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func GetAccoutByUsername(username string) {
	url := fmt.Sprintf(ACCOUNT_JSON_INFO, username)
	resp, err := http.Get(url)

	if err != nil {

	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}

	fmt.Print(body)
}
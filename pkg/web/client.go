package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TODO set user agent properly
var user_agent = "kitigo 0.0.1"

func Get(url string, out interface{}) interface{} {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("User-Agent", user_agent)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	if res.StatusCode != 200 {
		fmt.Println(res.Status)
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&out)

	return out
}

package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// TODO set user agent properly
var user_agent = "kiti 0.0.1"

func Get(url string, out interface{}) error {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("User-Agent", user_agent)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return errors.New("something went wrong://")
	}
	if res.StatusCode != 200 {
		fmt.Println(res.Status)
		return errors.New(res.Status)
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&out)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

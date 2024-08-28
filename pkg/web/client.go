package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// TODO set user agent properly
var user_agent = "kiti 0.0.1"

func get(url string) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("User-Agent", user_agent)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("something went wrong://")
	}
	if res.StatusCode != 200 {
		fmt.Println(res.Status)
		return nil, errors.New(res.Status)
	}

	return res, nil
}

func GetData(url string, out interface{}) error {

	res, err := get(url)
	if err != nil {
		return err
	}

	err = json.NewDecoder(res.Body).Decode(&out)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	return nil
}

func DownloadImage(url string) (string, error) {
	res, err := get(url)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	file, err := os.CreateTemp("", "kiti-")
	if err != nil {
		return "", err
	}

	file.Write(body)

	return file.Name(), nil
}

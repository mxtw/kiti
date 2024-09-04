package web

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

type webClient struct {
	Authorization string
	UserAgent     string
	Client        *http.Client
}

func NewClient() webClient {
	return webClient{
		Authorization: "",
		UserAgent:     "kiti 0.0.1",
		Client:        &http.Client{},
	}
}

func (wc webClient) get(url string) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}

	if wc.UserAgent != "" {
		req.Header.Set("User-Agent", wc.UserAgent)
	}
	if wc.Authorization != "" {
		req.Header.Set("Authorization", wc.Authorization)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, errors.New("something went wrong://")
	}
	if res.StatusCode != 200 {
		log.Println(res.Status)
		return nil, errors.New(res.Status)
	}

	return res, nil
}

func (wc webClient) GetData(url string, out interface{}) error {

	res, err := wc.get(url)
	if err != nil {
		return err
	}

	err = json.NewDecoder(res.Body).Decode(&out)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()

	return nil
}

func (wc webClient) DownloadImage(url string) (string, error) {
	res, err := wc.get(url)
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

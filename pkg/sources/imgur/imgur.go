package imgur

import (
	"errors"
	"fmt"

	"github.com/mxtw/kiti/pkg/util"
	"github.com/mxtw/kiti/pkg/web"
)

type imgurData struct {
	Data data `json:"data"`
}

type data struct {
	Items []item `json:"items"`
}

type item struct {
	Images []image `json:"images"`
}

type image struct {
	Link string `json:"link"`
}

func GetRandomUrlFromTag(tag string, clientId string) (string, error) {

	url := fmt.Sprintf("https://api.imgur.com/3/gallery/t/%s", tag)

	wc := web.NewClient()
	wc.Authorization = fmt.Sprintf("Client-ID %s", clientId)

	var response imgurData
	err := wc.GetData(url, &response)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("problem while fetching data from source")
	}

	pictures := make([]string, 0)

	for _, item := range response.Data.Items {
		for _, image := range item.Images {
			if util.IsFileValid(image.Link) {
				pictures = append(pictures, image.Link)
			}
		}
	}

	return util.RandomImage(pictures), nil
}

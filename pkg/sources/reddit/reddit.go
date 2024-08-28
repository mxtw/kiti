package reddit

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/mxtw/kiti/pkg/util"
	"github.com/mxtw/kiti/pkg/web"
)

type redditData struct {
	Data data `json:"data"`
}

type data struct {
	Children []child `json:"children"`
}

type child struct {
	Data childData `json:"data"`
}

type childData struct {
	Url string `json:"url"`
}

func GetRandomUrlFromSubreddit(subreddit string) (string, error) {

	url := fmt.Sprintf("https://reddit.com/r/%s/new.json", subreddit)

	var reddit redditData
	err := web.GetData(url, &reddit)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("problem while fetching data from source")
	}

	pictures := make([]string, 0)

	for _, child := range reddit.Data.Children {
		if util.IsFileValid(child.Data.Url) {
			pictures = append(pictures, child.Data.Url)
		}
	}

	return pictures[rand.Intn(len(pictures))], nil
}

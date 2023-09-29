package reddit

import (
	"fmt"
	"github.com/mxtw/kitigo/pkg/web"
	"math/rand"
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

func GetRandomUrl(subreddit string) string {

	url := fmt.Sprintf("https://reddit.com/r/%s/new.json", subreddit)

	var reddit redditData
	web.Get(url, &reddit)

	len := len(reddit.Data.Children)
	return reddit.Data.Children[rand.Intn(len)].Data.Url
}

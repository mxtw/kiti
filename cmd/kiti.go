package main

import (
	"fmt"

	"github.com/elias-gill/wallpaper"
	"github.com/mxtw/kiti/pkg/sources/reddit"
)

var version = "0.0.1"

func main() {
	// TODO get this from cli arg and/or config file
	// TODO handle different sources
	url, err := reddit.GetRandomUrlFromSubreddit("blurrypicturesofcats")
	if err != nil {
		fmt.Println(err)
	}

	err = wallpaper.SetFromURL(url)
	if err != nil {
		fmt.Println(err)
	}
}

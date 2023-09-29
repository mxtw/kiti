package main

import (
	"fmt"

	"github.com/elias-gill/wallpaper"
	"github.com/mxtw/kitigo/pkg/sources/reddit"
)

var version = "0.0.1"

func main() {
	// TODO get this from cli arg and/or config file
	url := reddit.GetRandomUrl("blurrypicturesofcats")

	err := wallpaper.SetFromURL(url)
	if err != nil {
		fmt.Println("oopsie")
	}
}

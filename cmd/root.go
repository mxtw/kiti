package cmd

import (
	"log"

	"github.com/mxtw/kiti/pkg/sources/imgur"
	"github.com/mxtw/kiti/pkg/sources/reddit"
	"github.com/mxtw/kiti/pkg/wallpaper"
	"github.com/spf13/cobra"
)

var (
	searchString  string
	source        string
	imgurClientId string

	rootCmd = &cobra.Command{
		Use:   "kiti",
		Short: "Kiti sets your wallpaper to a random image",
		Long:  "Kiti is a cross-platform tool to set your wallpaper to a random image from various sources written in Go.",
		Run: func(cmd *cobra.Command, args []string) {

			var url string
			var err error

			switch source {
			case "reddit":
				url, err = reddit.GetRandomUrlFromSubreddit(searchString)
				if err != nil {
					log.Fatal(err)
				}
			case "imgur":

				if imgurClientId == "" {
					log.Fatal("source imgur requires the `--clientid` flag to be set")
				}

				url, err = imgur.GetRandomUrlFromTag(searchString, imgurClientId)
				if err != nil {
					log.Fatal(err)
				}
			default:
				log.Fatal("invalid source")
			}

			err = wallpaper.SetFromUrl(url)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	// this is probably going to change
	rootCmd.PersistentFlags().StringVarP(
		&searchString,
		"search",
		"s",
		"blurrypicturesofcats",
		"Subreddit (reddit)/Tag (imgur) to pull image from")
	// TODO use viper for config file

	rootCmd.PersistentFlags().StringVarP(
		&source,
		"source",
		"S",
		"reddit",
		"Source to pull images from",
	)

	// TODO this could definitely need some improvement
	rootCmd.PersistentFlags().StringVarP(
		&imgurClientId,
		"clientid",
		"i",
		"",
		"imgur client id (see https://apidocs.imgur.com/#authorization-and-oauth)",
	)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

package cmd

import (
	"fmt"
	"os"

	"github.com/mxtw/kiti/pkg/sources/reddit"
	"github.com/mxtw/kiti/pkg/wallpaper"
	"github.com/spf13/cobra"
)

var (
	subReddit string

	rootCmd = &cobra.Command{
		Use:   "kiti",
		Short: "Kiti sets your wallpaper to a random image",
		Long:  "Kiti is a cross-platform tool to set your wallpaper to a random image from various sources written in Go.",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO handle different/multiple sources
			url, err := reddit.GetRandomUrlFromSubreddit(subReddit)
			if err != nil {
				fmt.Println(err)
			}

			err = wallpaper.SetFromUrl(url)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
)

func init() {
	// this is probably going to change
	rootCmd.PersistentFlags().StringVarP(
		&subReddit,
		"sub",
		"s",
		"blurrypicturesofcats",
		"Subreddit to pull image from")
	// TODO use viper for config file
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

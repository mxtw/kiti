package util

import (
	"log"
	"math/rand"
	"regexp"
)

func IsFileValid(input string) bool {
	// primitive check if it has valid extension
	// TODO more robust check in the future maybe?
	re, err := regexp.Compile(".*.(jpeg|jpg|png)$")
	if err != nil {
		log.Println(err)
	}
	return re.Match([]byte(input))
}

func RandomImage(pictures []string) string {
	return pictures[rand.Intn(len(pictures))]
}

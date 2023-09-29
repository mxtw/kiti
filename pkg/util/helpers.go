package util

import (
	"fmt"
	"regexp"
)

func IsFileValid(input string) bool {
	// primitive check if it has valid extension
	// TODO more robust check in the future maybe?
	re, err := regexp.Compile(".*.(jpeg|jpg|png)$")
	if err != nil {
		fmt.Println(err)
	}
	return re.Match([]byte(input))
}

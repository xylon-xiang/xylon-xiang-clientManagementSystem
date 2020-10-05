package util

import (
	"regexp"
)

// just suit for Url like this: http://localhost:4567/sudent/:student
func GetRealUrl(url string, studentId string) string {

	re := regexp.MustCompile(`/:.*[/]?`)

	urlStr := re.ReplaceAll([]byte(url),
		[]byte("/"+studentId))

	return string(urlStr)
}

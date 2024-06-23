package pkg

import "regexp"

var enumRegexp = regexp.MustCompile(`\(([^)]+)\)`)

func HasEnum(str string) bool {
	matches := enumRegexp.FindStringSubmatch(str)
	return len(matches) > 0
}

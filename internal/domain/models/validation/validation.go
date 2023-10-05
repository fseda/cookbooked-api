package modelvalidation

import "regexp"

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func IsEmailLike(input string) bool {
	return emailRegex.MatchString(input)
}

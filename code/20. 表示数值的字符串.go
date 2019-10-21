package code

import "regexp"

var reg = regexp.MustCompile(`^[\s]*[-+]?(([\d]+([.][\d]*)?)|([.][\d]+))([eE][-+]?[\d]+)?[\s]*$`)

func IsNumeric(str string) bool {
	return reg.MatchString(str)
}

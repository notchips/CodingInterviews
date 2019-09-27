var reg = regexp.MustCompile(`^[\s]*[-+]?(([\d]+([.][\d]*)?)|([.][\d]+))([eE][-+]?[\d]+)?[\s]*$`)

func isNumeric(str string) bool {
	return reg.MatchString(str)
}
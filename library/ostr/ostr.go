package ostr

import (
	"strings"
)

//SplitString2StringSlice convert a string to a string slice
func SplitString2StringSlice(str string) ([]string, error) {
	var (
		res []string
		err error
	)
	strArr := strings.Split(str, "\n")
	res = make([]string, 0, len(strArr))
	for _, sc := range strArr {
		sc = strings.TrimSpace(sc)
		if len(sc) == 0 {
			continue
		}
		res = append(res, sc)
	}
	return res, err
}

package ostr

import (
	"strconv"
	"strings"
)

//JudgePreNumber determine whether the text is purely a number
func JudgePreNumber(str string) bool {
	var rule = []string{" ", ",", "-", "."}
	for _, r := range rule {
		str = strings.ReplaceAll(str, r, "")
	}
	if _, err := strconv.Atoi(str); err == nil {
		return true
	}
	return false
}

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

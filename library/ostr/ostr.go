/*
ostr： strings related operation
*/
package ostr

import (
	"bytes"
	"strconv"
	"strings"
	"sync"
)

var (
	buffPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer([]byte{})
		},
	}
)

// JudgePreNumber determine whether the text is purely a number
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

// SplitString2StringSlice convert a string to a string slice
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

// SplitIntSlice2String convert a int64 slice to a string,like int1,int2,int3
func SplitIntSlice2String(intSlice []int64) string {
	if len(intSlice) == 0 {
		return ""
	}
	if len(intSlice) == 1 {
		return strconv.FormatInt(intSlice[0], 10)
	}
	buf := buffPool.Get().(*bytes.Buffer)
	for _, i := range intSlice {
		buf.WriteString(strconv.FormatInt(i, 10))
		buf.WriteString(",")
	}
	result := buf.String()
	buf.Reset()
	buffPool.Put(buf)
	return result
}

// SplitString2IntSlice split string into int64 slice
func SplitString2IntSlice(str string) ([]int64, error) {
	if len(str) == 0 {
		return nil, nil
	}
	splitStr := strings.Split(str, ",")
	result := make([]int64, 0, len(splitStr))
	for _, s := range splitStr {
		parseInt, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, parseInt)
	}
	return result, nil
}

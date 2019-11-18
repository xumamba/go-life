package ostr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testStr = `
    hello
     
	你 好
	세계
`
	expectStrSlice = []string{"hello", "你 好", "세계"}
)

func Test_SplitString2StringSlice(t *testing.T) {
	stringSlice, err := SplitString2StringSlice(testStr)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, len(expectStrSlice), len(stringSlice))
	assert.Equal(t, expectStrSlice, stringSlice)
}

func Test_JudgePreNumber(t *testing.T) {
	testStr := "1 2,3.4-5"
	judgeResult := JudgePreNumber(testStr)
	assert.Equal(t, true, judgeResult)
}

func Test_SplitIntSlice2String(t *testing.T) {
	testSlice := []int64{12, 16, 1, 93}
	str := SplitIntSlice2String(testSlice)
	assert.Equal(t, "12,16,1,93,", str)
	str2 := contrastPerformance(testSlice)
	assert.Equal(t, "12,16,1,93,", str2)
}

func Test_SplitString2IntSlice(t *testing.T) {
	testStr := "12,16,1,93"
	result, err := SplitString2IntSlice(testStr)
	assert.Equal(t, nil, err)
	assert.Equal(t, []int64{12, 16, 1, 93}, result)
}

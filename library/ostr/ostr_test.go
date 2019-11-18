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

func BenchmarkSplitString2StringSlice(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			SplitString2StringSlice(testStr)
		}
	})
}

//BenchmarkSplitString2StringSlice-4   	 4607433	       248 ns/op	     192 B/op	       2 allocs/op

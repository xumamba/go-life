package ostr

import (
	"strconv"
	"testing"
)

func BenchmarkSplitIntSlice2String(b *testing.B) {
	testSlice := []int64{12, 16, 1, 93}
	b.Run("useBuffer", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			SplitIntSlice2String(testSlice)
		}
	})
	b.Run("useConnection", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			contrastPerformance(testSlice)
		}
	})
}

//BenchmarkSplitIntSlice2String/useBuffer-4         	 8581812	       139 ns/op	      16 B/op	       1 allocs/op
//BenchmarkSplitIntSlice2String/useConnection-4     	 5775472	       204 ns/op	      32 B/op	       4 allocs/op

func BenchmarkSplitString2StringSlice(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = SplitString2StringSlice(testStr)
		}
	})
}

//BenchmarkSplitString2StringSlice-4   	 4607433	       248 ns/op	     192 B/op	       2 allocs/op

func contrastPerformance(intSlice []int64) string {
	if len(intSlice) == 0 {
		return ""
	}
	if len(intSlice) == 1 {
		return strconv.FormatInt(intSlice[0], 10)
	}
	strResult := ""
	for _, i := range intSlice {
		strResult = strResult + strconv.FormatInt(i, 10) + ","
	}
	return strResult
}

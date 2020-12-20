package main

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

var a = []int32("6shb")
var b = []int32("7bs")

func Test(t *testing.T){
	split := strings.Split("a b", " ")
	fmt.Println(split)
}
func sortByBits(arr []int) []int {
	ans := []int{}
	mapBits := map[int][]int{}
	countBits := func(v int) {
		cnt := 0
		for i := uint(0); i < 32; i++ {
			if 1 == 1&(v>>i) {
				cnt++
			}
		}
		mapBits[cnt] = append(mapBits[cnt], v)
	}
	for _, v := range arr {
		countBits(v)
	}
	keys := []int{}
	for i := range mapBits {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	for _, v := range keys {
		sort.Ints(mapBits[v])
		for _, v1 := range mapBits[v] {
			ans = append(ans, v1)
		}
	}
	return ans
}
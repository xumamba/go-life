package main

import (
	"fmt"
	"testing"
)

func TestIdea(t *testing.T) {
	s := []int{1, 2, 3, 4}
	m := make(map[int]int)
	for index, value := range s {
		fmt.Println(&index, &value) // 0xc00001a290 0xc00001a298 循环过程中变量使用同一个地址
		m[index] = value
	}
	fmt.Println(m)
}

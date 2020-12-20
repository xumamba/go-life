package designandimplementation

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	a := [200]interface{}{}
	fmt.Println(&a[0], &a[1])  // 0xc0000a6000 0xc0000a6010  16字节

	aa := [2]int{5, 6}
	bb := [2]int{5, 6}
	fmt.Println(aa==bb)  // true
	fmt.Println(&aa==&bb)  // false
}

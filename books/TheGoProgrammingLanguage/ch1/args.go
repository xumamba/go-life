package ch1

import (
	"fmt"
	"os"
	"strings"
)

func OsArgs() {
	var s,sep string
	for i:=0; i<len(os.Args);i++{
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	fmt.Println(strings.Join(os.Args[1:], " "))
}
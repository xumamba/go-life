package main

import (
	"reflect"
	"testing"
)

type S1 struct {
	name string
	age int
	hobby []string
}

type S2 struct {
	name string
	age int
}

func TestCompareStruct(t *testing.T) {
	s1 := S1{hobby: []string{"read"}}
	s2 := S2{name: "aa", age: 12}
	s3 := S2{age: 12, name: "aa"}
	// mismatched types S1 and S2
	println(s2 == s3) // true
	s4 := S1{hobby: []string{"read"}}
	println(reflect.DeepEqual(s1, s4)) // true
}

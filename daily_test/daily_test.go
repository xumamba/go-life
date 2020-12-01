package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type S1 struct {
	name  string
	age   int
	hobby []string
}

type S2 struct {
	name string
	age  int
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

func TestSelect(t *testing.T) {
	var ch = make(chan int)
	go func() {
		ch <- 1
	}()
	time.Sleep(1 *time.Second)
	select {
	case x := <- ch:
		fmt.Println("receive: ", x)
	case y := <-time.After(3 * time.Second):
		fmt.Println("y is: ", y)
	default:
		fmt.Println("nothing")
	}
	fmt.Println("exit.")
}

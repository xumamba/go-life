package ch2

import (
	"log"
	"os"
)

var global *int

func escapeOne() *int  {
	var x int
	x = 1
	return &x
}

func escapeTwo()  {
	y := new(int)
	*y = 1
	global = y
}

type Name string

var xx Name

func (n Name) String() string {
	return string(n)
}


var cwd string

func init() {
	cwd, err := os.Getwd() // NOTE: wrong!
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory = %s", cwd)
}
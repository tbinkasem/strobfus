// +build ignore

package main

import "fmt"

//go:generate strobfus -filename $GOFILE

// a little comment
var hello string

var yolo = "poeut"

var arr = []string{
	"a",
	"b",
}

var (
	str1 = "coucou"
	arr1 = []string{
		"foo",
		"bar",
	}
	arr2 = []string{"h", "g"}
	arr3 = []string{"unique entry"}
)

func init() {
	fmt.Println("This my init")
}

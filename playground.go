package main

import "fmt"

type hotdog int

var h hotdog = 10

func main() {
	x := 5
	fmt.Printf("%v\n", h)
	z := int(h) / x
	fmt.Printf("%v", z)
}

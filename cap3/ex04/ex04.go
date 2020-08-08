package main

import "fmt"

type jollaz int

var x jollaz = 10
var y int

func main() {
	fmt.Printf("x = %v, do tipo %T\n", x, x)
	x = 42
	fmt.Printf("x = %v\n", x)

	y = int(x)
	fmt.Printf("y = %v do tipo %T", y, y)
}

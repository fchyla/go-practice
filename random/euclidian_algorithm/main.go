package main

import "fmt"

func findints(x int, y int) int {
	for y != 0 {
		t := y
		y = x % y
		x = t
	}
	return x
}

func main() {
	x := 4898903857390434234
	y := 2323890292123178977

	r := findints(x, y)

	fmt.Printf("%d", r)
}
